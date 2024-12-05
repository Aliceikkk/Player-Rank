package main

import (
	"data/logger"
	"embed"
	"net/http"
	"strconv"
	"time"

	"data/config"
	"data/service"
)

//go:embed rank/*
var PageFS embed.FS

func main() {
	logger.InitLogger()
	defer logger.CloseLogger()

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config", err)
		return
	}

	// 创建自动更新数据服务
	syncService := service.NewSyncService(cfg)
	logger.Info("玩家Rank开启成功")

	httpService := service.NewHttpService(syncService)

	// 启动同步服务的goroutine
	go func() {
		ticker := time.NewTicker(time.Duration(cfg.SyncIntervalMinutes) * time.Minute)
		defer ticker.Stop()

		for {
			logger.Info("开始更新全服数据")
			if err := syncService.SyncData(); err != nil {
				logger.Error("更新全服数据失败: ", err)
			}
			logger.Info("数据更新完成")

			<-ticker.C
		}
	}()

	// 启动网页服务器
	http.Handle("/", http.FileServer(http.FS(PageFS)))
	go func() {
		port := cfg.HttpPort
		if port == 0 {
			port = 7070
		}
		address := ":" + strconv.Itoa(port)
		logger.Info("HTTP服务运行在端口%s", address)
		if err := http.ListenAndServe(address, nil); err != nil {
			logger.Error("启动HTTP服务失败: %v", err)
		}
	}()

	// 启动Api服务器
	if err := httpService.StartServer(cfg.ApiPort); err != nil {
		logger.Error("Api server failed", err)
	}
}
