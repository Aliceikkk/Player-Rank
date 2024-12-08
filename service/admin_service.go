package service

import (
	"data/config"
	"data/logger"
	"data/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type AdminService struct {
	config     *config.Config
	configLock sync.RWMutex
}

func NewAdminService(config *config.Config) *AdminService {
	return &AdminService{
		config: config,
	}
}

// 检查是否是首次登录
func (s *AdminService) IsFirstLogin(db *sql.DB) (bool, error) {
	// 先检查表是否存在
	_, err := db.Query("SELECT 1 FROM t_admin LIMIT 1")
	if err != nil {
		// 如果表不存在，创建表
		_, err = db.Exec(`
            CREATE TABLE IF NOT EXISTS t_admin (
                id INT AUTO_INCREMENT PRIMARY KEY,
                username VARCHAR(32) NOT NULL UNIQUE,
                password VARCHAR(64) NOT NULL,
                created_at DATETIME DEFAULT CURRENT_TIMESTAMP
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
        `)
		if err != nil {
			logger.Error("创建管理员表失败:", err)
			return false, err
		}
		return true, nil
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM t_admin").Scan(&count)
	if err != nil {
		logger.Error("查询管理员数量失败:", err)
		return false, err
	}
	return count == 0, nil
}

// 创建管理员账号
func (s *AdminService) CreateAdmin(db *sql.DB, username, password string) error {
	hashedPassword := model.HashPassword(password)
	// 检查用户名是否已存在
	var exists int
	err := db.QueryRow("SELECT COUNT(*) FROM t_admin WHERE username = ?", username).Scan(&exists)
	if err != nil {
		logger.Error("检查用户名是否存在失败:", err)
		return err
	}
	if exists > 0 {
		return fmt.Errorf("用户名已存在")
	}

	_, err = db.Exec(`
        INSERT INTO t_admin (username, password, created_at)
        VALUES (?, ?, ?)
    `, username, hashedPassword, time.Now())
	if err != nil {
		logger.Error("插入管理员记录失败:", err)
		return err
	}
	return err
}

// 验证管理员登录
func (s *AdminService) VerifyAdmin(db *sql.DB, username, password string) (bool, error) {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM t_admin WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("用户名不存在:", username)
			return false, nil
		}
		logger.Error("查询管理员密码失败:", err)
		return false, err
	}

	hashedPassword := model.HashPassword(password)
	logger.Info(fmt.Sprintf("验证管理员登录: username=%s, valid=%v", username, storedPassword == hashedPassword))
	return storedPassword == hashedPassword, nil
}

// 获取配置
func (s *AdminService) GetConfig() *config.Config {
	s.configLock.RLock()
	defer s.configLock.RUnlock()
	return s.config
}

// 保存配置
func (s *AdminService) SaveConfig(newConfig *config.Config) error {
	s.configLock.Lock()
	defer s.configLock.Unlock()

	// 验证时间格式
	if _, err := time.Parse("2006-01-02 15:04:05", newConfig.LastSaveTime); err != nil {
		return fmt.Errorf("invalid time format: %v", err)
	}

	// 只更新允许修改的字段
	s.config.Announcement = newConfig.Announcement
	s.config.LastSaveTime = newConfig.LastSaveTime
	s.config.SyncIntervalMinutes = newConfig.SyncIntervalMinutes

	// 将配置写入文件
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(s.config); err != nil {
		return err
	}

	return nil
}

// 清空排行榜
func (s *AdminService) ClearLeaderboard(db *sql.DB) error {
	_, err := db.Exec("TRUNCATE TABLE t_player_stats")
	return err
}

// 监控配置文件变化
func (s *AdminService) StartConfigMonitor() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			newConfig, err := config.LoadConfig()
			if err != nil {
				logger.Error("监控配置文件失败:", err)
				continue
			}

			s.configLock.Lock()
			s.config = newConfig
			s.configLock.Unlock()

			logger.Info("配置文件已重新加载")
		}
	}()
}
