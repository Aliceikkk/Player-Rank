package service

import (
	"data/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type HttpService struct {
	syncService *SyncService
}

func NewHttpService(syncService *SyncService) *HttpService {
	return &HttpService{
		syncService: syncService,
	}
}

func (h *HttpService) StartServer(port int) error {
	signMiddleware := NewSignatureMiddleware(h.syncService.config.ApiSecret)

	// API路由
	http.HandleFunc("/api/leaderboards", signMiddleware.Verify(h.handleLeaderboards))
	http.HandleFunc("/api/announcement", signMiddleware.Verify(h.handleAnnouncement))
	http.HandleFunc("/api/check-update", signMiddleware.Verify(h.handleCheckUpdate))

	addr := fmt.Sprintf(":%d", port)
	logger.Info(fmt.Sprintf("API服务运行在端口: %d", port))
	return http.ListenAndServe(addr, nil)
}

func (h *HttpService) handleLeaderboards(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 获取页码参数，默认为1
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	pageSize := 8

	// 连接数据库
	targetDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		h.syncService.config.TargetMySQL.User,
		h.syncService.config.TargetMySQL.Password,
		h.syncService.config.TargetMySQL.Host,
		h.syncService.config.TargetMySQL.Port,
		h.syncService.config.TargetMySQL.Database)

	db, err := sql.Open("mysql", targetDSN)
	if err != nil {
		logger.Error("Failed to connect to database", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 获取分页的排行榜数据
	data, err := h.syncService.GetLeaderboards(db, page, pageSize)
	if err != nil {
		logger.Error("Failed to get leaderboards", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (h *HttpService) handleAnnouncement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"announcement": h.syncService.config.Announcement,
	}

	//logger.Info(fmt.Sprintf("Sending announcement: %s", h.syncService.config.Announcement))

	json.NewEncoder(w).Encode(response)
}

func (h *HttpService) handleCheckUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// 连接数据库
	targetDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		h.syncService.config.TargetMySQL.User,
		h.syncService.config.TargetMySQL.Password,
		h.syncService.config.TargetMySQL.Host,
		h.syncService.config.TargetMySQL.Port,
		h.syncService.config.TargetMySQL.Database)

	db, err := sql.Open("mysql", targetDSN)
	if err != nil {
		logger.Error("Failed to connect to database", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 获取上次更新时间
	lastUpdate, err := h.syncService.GetLastUpdateTime(db)
	if err != nil {
		logger.Error("Failed to get last update time", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 获取客户端传来的时间戳
	clientTimestamp := r.URL.Query().Get("timestamp")
	clientTime, err := strconv.ParseInt(clientTimestamp, 10, 64)
	if err != nil {
		http.Error(w, "Invalid timestamp", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"hasUpdate": lastUpdate > clientTime,
		"timestamp": lastUpdate,
	}
	json.NewEncoder(w).Encode(response)
}
