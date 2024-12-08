package service

import (
	"data/config"
	"data/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type HttpService struct {
	syncService  *SyncService
	adminService *AdminService
}

func NewHttpService(syncService *SyncService, adminService *AdminService) *HttpService {
	return &HttpService{
		syncService:  syncService,
		adminService: adminService,
	}
}

func (h *HttpService) getDB() (*sql.DB, error) {
	targetDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		h.syncService.config.TargetMySQL.User,
		h.syncService.config.TargetMySQL.Password,
		h.syncService.config.TargetMySQL.Host,
		h.syncService.config.TargetMySQL.Port,
		h.syncService.config.TargetMySQL.Database)

	return sql.Open("mysql", targetDSN)
}

func (h *HttpService) StartServer(port int) error {
	signMiddleware := NewSignatureMiddleware(h.syncService.config.ApiSecret)

	// API路由
	http.HandleFunc("/api/leaderboards", signMiddleware.Verify(h.handleLeaderboards))
	http.HandleFunc("/api/announcement", signMiddleware.Verify(h.handleAnnouncement))
	http.HandleFunc("/api/check-update", signMiddleware.Verify(h.handleCheckUpdate))
	http.HandleFunc("/api/admin/login", h.handleAdminLogin)
	http.HandleFunc("/api/admin/config", h.handleConfig)
	http.HandleFunc("/api/admin/clear", h.handleClearLeaderboard)

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

func (h *HttpService) handleAdminLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("解析登录请求失败:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Invalid request format",
		})
		return
	}

	db, err := h.getDB()
	if err != nil {
		logger.Error("数据库连接失败:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Database connection error",
		})
		return
	}
	defer db.Close()

	// 检查是否是首次登录
	isFirst, err := h.adminService.IsFirstLogin(db)
	if err != nil {
		logger.Error("检查首次登录失败:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Failed to check first login",
		})
		return
	}

	if isFirst {
		// 首次登录，创建管理员账号
		if err := h.adminService.CreateAdmin(db, req.Username, req.Password); err != nil {
			logger.Error("创建管理��账号失败:", err)
			http.Error(w, "Failed to create admin", http.StatusInternalServerError)
			return
		}
		response := map[string]interface{}{
			"success": true,
			"message": "Admin account created",
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Error("编码响应失败:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		return
	}

	// 验证登录
	valid, err := h.adminService.VerifyAdmin(db, req.Username, req.Password)
	if err != nil {
		logger.Error("验证管理员失败:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": valid,
	}
	if !valid {
		response["message"] = "Invalid credentials"
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Error("编码响应失败:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}

func (h *HttpService) handleConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		// 只返回可以修改的配置字段
		config := h.adminService.GetConfig()
		response := map[string]interface{}{
			"announcement":          config.Announcement,
			"last_save_time":       config.LastSaveTime,
			"sync_interval_minutes": config.SyncIntervalMinutes,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Method == "POST" {
		var newConfig config.Config
		if err := json.NewDecoder(r.Body).Decode(&newConfig); err != nil {
			logger.Error("解析配置失败:", err)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "Invalid request format",
			})
			return
		}

		if err := h.adminService.SaveConfig(&newConfig); err != nil {
			logger.Error("保存配置失败:", err)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "Failed to save config",
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
		})
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func (h *HttpService) handleClearLeaderboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	db, err := h.getDB()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := h.adminService.ClearLeaderboard(db); err != nil {
		http.Error(w, "Failed to clear leaderboard", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}
