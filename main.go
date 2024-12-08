package main

import (
	"data/logger"
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"data/config"
	"data/service"
)

//go:embed rank/*
var PageFS embed.FS

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

	// 初始化验证服务
	verifyService := service.NewVerifyService()
	
	if !verifyService.CheckActivation() {
		machineID := verifyService.GenerateMachineID()
		fmt.Printf("请输入激活码 (机器ID: %s): ", machineID)
		
		var activationCode string
		fmt.Scanln(&activationCode)
		
		if !verifyService.VerifyActivationCode(activationCode, machineID) {
			logger.Error("激活码无效")
			return
		}
		
		if err := verifyService.SaveActivationCode(activationCode); err != nil {
			logger.Error("保存激活码失败:", err)
			return
		}
	}

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
	errChan := make(chan error, 3)

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

	// 启动HTTP服务
	go func() {
		port := cfg.HttpPort
		if port == 0 {
			port = 7070
		}
		address := ":" + strconv.Itoa(port)
		logger.Info(fmt.Sprintf("HTTP服务运行在端口%s", address))
		
		if err := http.ListenAndServe(address, http.FileServer(http.FS(PageFS))); err != nil {
			logger.Error(fmt.Sprintf("HTTP服务失败: %v", err))
			errChan <- fmt.Errorf("HTTP服务失败: %v", err)
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
		// 这里可以添加优雅关闭的逻辑
		time.Sleep(time.Second) // 给一些时间让服务完成当前操作
		os.Exit(0)
	}
}
