package main

import (
	"data/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"data/config"
	"data/service"
)

func main() {
	logger.InitLogger()
	defer logger.CloseLogger()

	// 添加错误恢复
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("程序发生严重错误: %v", r))
			os.Exit(1)
		}
	}()

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("加载配置失败", err)
		return
	}

	// 创建自动更新数据服务
	syncService := service.NewSyncService(cfg)
	logger.Info("玩家Rank开启成功")

	// 创建管理员服务
	adminService := service.NewAdminService(cfg)
	adminService.StartConfigMonitor() // 启动配置监控

	httpService := service.NewHttpService(syncService, adminService)

	// 创建错误通道
	errChan := make(chan error, 2)

	// 启动数据同步服务
	go func() {
		ticker := time.NewTicker(time.Duration(cfg.SyncIntervalMinutes) * time.Minute)
		defer ticker.Stop()

		// 首次运行
		logger.Info("开始更新全服数据")
		if err := syncService.SyncData(); err != nil {
			logger.Error("首次更新全服数据失败", err)
			errChan <- fmt.Errorf("数据同步失败: %v", err)
			return
		}
		logger.Info("首次数据更新完成")

		for {
			select {
			case <-ticker.C:
				logger.Info("开始定时更新全服数据")
				if err := syncService.SyncData(); err != nil {
					logger.Error("定时更新全服数据失败", err)
					errChan <- fmt.Errorf("数据同步失败: %v", err)
					return
				}
				logger.Info("定时数据更新完成")
			}
		}
	}()

	// 启动API服务
	go func() {
		if err := httpService.StartServer(cfg.ApiPort); err != nil {
			logger.Error("API服务失败", err)
			errChan <- fmt.Errorf("API服务失败: %v", err)
		}
	}()

	// 处理系统信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 等待错误或系统信号
	select {
	case err := <-errChan:
		logger.Error(fmt.Sprintf("服务发生错误: %v", err))
		os.Exit(1)
	case sig := <-sigChan:
		logger.Info(fmt.Sprintf("接收到系统信号: %v, 正在关闭服务...", sig))
		time.Sleep(time.Second)
		os.Exit(0)
	}
}
