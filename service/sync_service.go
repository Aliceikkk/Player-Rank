package service

import (
	"database/sql"
	"fmt"
	"time"

	"data/config"
	"data/logger"
	"data/pb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/proto"
)

type SyncService struct {
	config *config.Config
}

func NewSyncService(config *config.Config) *SyncService {
	return &SyncService{
		config: config,
	}
}

func (s *SyncService) SyncData() error {
	startTime := time.Now()
	totalUpdated := 0

	// 连接玩家数据库
	sourceDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		s.config.SourceMySQL.User,
		s.config.SourceMySQL.Password,
		s.config.SourceMySQL.Host,
		s.config.SourceMySQL.Port,
		s.config.SourceMySQL.Database)

	sourceDB, err := sql.Open("mysql", sourceDSN)
	if err != nil {
		return fmt.Errorf("failed to connect to source database: %v", err)
	}
	defer sourceDB.Close()

	// 连接目标数据库
	targetDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		s.config.TargetMySQL.User,
		s.config.TargetMySQL.Password,
		s.config.TargetMySQL.Host,
		s.config.TargetMySQL.Port,
		s.config.TargetMySQL.Database)

	targetDB, err := sql.Open("mysql", targetDSN)
	if err != nil {
		return fmt.Errorf("failed to connect to target database: %v", err)
	}
	defer targetDB.Close()

	if err := s.createTargetTable(targetDB); err != nil {
		return err
	}

	updated, err := s.processAllTables(sourceDB, targetDB)
	if err != nil {
		return err
	}
	totalUpdated += updated

	duration := time.Since(startTime)
	logger.Info(fmt.Sprintf("更新全服数据完成 - 更新 %d 条数据, 花费了 %v 时间", totalUpdated, duration))
	return nil
}

func (s *SyncService) createTargetTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS t_player_stats (
			uid INT UNSIGNED PRIMARY KEY,
			nickname VARCHAR(32) NOT NULL,
			max_critical_damage FLOAT NOT NULL,
			max_fly_map_distance FLOAT NOT NULL,
			last_update_time BIGINT NOT NULL,
			INDEX idx_update_time (last_update_time)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
	`)
	return err
}

func (s *SyncService) processAllTables(sourceDB, targetDB *sql.DB) (int, error) {
	totalUpdated := 0
	for i := 0; i < 10; i++ {
		tableName := fmt.Sprintf("t_player_data_%d", i)
		//logger.Info(fmt.Sprintf("Processing table: %s", tableName))

		updated, err := s.processTable(tableName, sourceDB, targetDB)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to process table: %s", tableName), err)
			continue
		}
		totalUpdated += updated
	}
	return totalUpdated, nil
}

func (s *SyncService) processTable(tableName string, sourceDB, targetDB *sql.DB) (int, error) {
	// 解析配置中的时间
	configTime, err := time.Parse("2006-01-02 15:04:05", s.config.LastSaveTime)
	if err != nil {
		return 0, fmt.Errorf("解析配置时间失败: %v", err)
	}

	// 添加调试日志
	logger.Info(fmt.Sprintf("配置的时间阈值: %v", configTime))

	// 先检查表中是否有符合条件的数据
	var count int
	err = sourceDB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE last_save_time > ?", tableName), configTime).Scan(&count)
	if err != nil {
		logger.Error(fmt.Sprintf("检查表 %s 数据失败: %v", tableName, err))
	} else {
		logger.Info(fmt.Sprintf("表 %s 中有 %d 条数据需要更新", tableName, count))
	}

	// 修改查询语句，加入last_save_time字段
	rows, err := sourceDB.Query(fmt.Sprintf("SELECT uid, bin_data, last_save_time FROM %s", tableName))
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	updated := 0
	for rows.Next() {
		var uid uint32
		var binData []byte
		var lastSaveTimeStr sql.NullString
		if err := rows.Scan(&uid, &binData, &lastSaveTimeStr); err != nil {
			logger.Error(fmt.Sprintf("扫描行数据失败: %v", err))
			continue
		}

		// 如果 last_save_time 为空或无效，跳过这条数据
		if !lastSaveTimeStr.Valid {
			//logger.Info(fmt.Sprintf("跳过 last_save_time 为空的数据 uid: %d", uid))
			continue
		}

		// 解析时间字符串
		lastSaveTime, err := time.Parse("2006-01-02 15:04:05", lastSaveTimeStr.String)
		if err != nil {
			logger.Error(fmt.Sprintf("解析时间失败 uid: %d, time: %s, error: %v", uid, lastSaveTimeStr.String, err))
			continue
		}

		// 添加调试日志
		//logger.Info(fmt.Sprintf("处理数据 uid: %d, last_save_time: %v", uid, lastSaveTime))

		// 比较时间，只处理配置时间之后的数据
		if !lastSaveTime.After(configTime) {
			//logger.Info(fmt.Sprintf("跳过旧数据 uid: %d, last_save_time: %v", uid, lastSaveTime))
			continue
		}

		if err := s.processRow(rows, targetDB); err != nil {
			//logger.Error(fmt.Sprintf("处理数据失败 uid: %d, error: %v", uid, err))
			continue
		}
		//logger.Info(fmt.Sprintf("成功更新数据 uid: %d", uid))
		updated++
	}
	logger.Info(fmt.Sprintf("表 %s 处理完成，更新了 %d 条数据", tableName, updated))
	return updated, rows.Err()
}

func (s *SyncService) processRow(rows *sql.Rows, targetDB *sql.DB) error {
	var uid uint32
	var binData []byte
	var lastSaveTimeStr string
	if err := rows.Scan(&uid, &binData, &lastSaveTimeStr); err != nil {
		return err
	}

	unzlibedData, err := UnzlibData(binData)
	if err != nil {
		return err
	}

	var playerDataBin pb.PlayerDataBin
	if err := proto.Unmarshal(unzlibedData, &playerDataBin); err != nil {
		return err
	}

	basicBin := playerDataBin.GetBasicBin()
	nickname := fmt.Sprintf("Player_%d", uid)
	if basicBin != nil {
		nickname = string(basicBin.GetNickname())
	}

	watcherBin := playerDataBin.GetWatcherBin()
	if watcherBin == nil {
		return nil
	}

	recordValue := watcherBin.GetRecordValue()
	if recordValue == nil {
		return nil
	}

	maxCriticalDamage := recordValue.GetMaxCriticalDamage()
	maxFlyMapDistance := recordValue.GetMaxFlyMapDistance()

	if maxCriticalDamage == 0 && maxFlyMapDistance == 0 {
		return nil
	}

	return s.updateStats(targetDB, uid, nickname, recordValue)
}

func (s *SyncService) updateStats(db *sql.DB, uid uint32, nickname string, recordValue *pb.RecordValueBin) error {
	_, err := db.Exec(`
		INSERT INTO t_player_stats 
		(uid, nickname, max_critical_damage, max_fly_map_distance, last_update_time)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		nickname = VALUES(nickname),
		max_critical_damage = VALUES(max_critical_damage),
		max_fly_map_distance = VALUES(max_fly_map_distance),
		last_update_time = VALUES(last_update_time)
	`,
		uid,
		nickname,
		recordValue.GetMaxCriticalDamage(),
		recordValue.GetMaxFlyMapDistance(),
		time.Now().Unix(),
	)

	if err != nil {
		return err
	}

	return nil
}

// 添加新的方法来获取排行榜数据
func (s *SyncService) GetLeaderboards(db *sql.DB, page, pageSize int) (map[string]interface{}, error) {
	offset := (page - 1) * pageSize

	// 获取最强一击排行榜
	damageRows, err := db.Query(`
		SELECT 
			nickname, 
			uid,
			max_critical_damage
		FROM t_player_stats 
		ORDER BY max_critical_damage DESC 
		LIMIT ? OFFSET ?
	`, pageSize, offset)
	if err != nil {
		logger.Error("获取最强一击排行榜数据失败", err)
		return nil, err
	}
	defer damageRows.Close()

	// 获取最远飞行距离排行榜
	flightRows, err := db.Query(`
		SELECT uid, nickname, max_fly_map_distance 
		FROM t_player_stats 
		ORDER BY max_fly_map_distance DESC 
		LIMIT ? OFFSET ?
	`, pageSize, offset)
	if err != nil {
		logger.Error("获取最远飞行距离排行榜数据失败", err)
		return nil, err
	}
	defer flightRows.Close()

	// 获取总记录数
	var totalDamage, totalFlight int
	err = db.QueryRow("SELECT COUNT(*) FROM t_player_stats WHERE max_critical_damage > 0").Scan(&totalDamage)
	if err != nil {
		logger.Error("获取最强一击总记录数失败", err)
		return nil, err
	}
	err = db.QueryRow("SELECT COUNT(*) FROM t_player_stats WHERE max_fly_map_distance > 0").Scan(&totalFlight)
	if err != nil {
		logger.Error("获取最远飞行距离总记录数失败", err)
		return nil, err
	}

	// 获取最后更新时间
	lastUpdate, err := s.GetLastUpdateTime(db)
	if err != nil {
		logger.Error("获取最后更新时间失败", err)
		return nil, err
	}

	result := map[string]interface{}{
		"damage": map[string]interface{}{
			"data":       []map[string]interface{}{},
			"total":      totalDamage,
			"totalPages": (totalDamage + pageSize - 1) / pageSize,
			"page":       page,
		},
		"flight": map[string]interface{}{
			"data":       []map[string]interface{}{},
			"total":      totalFlight,
			"totalPages": (totalFlight + pageSize - 1) / pageSize,
			"page":       page,
		},
		"timestamp": lastUpdate,
	}

	// 处理数据...（和之前一样，但数据放在data字段中）
	damageData := result["damage"].(map[string]interface{})["data"].([]map[string]interface{})
	for damageRows.Next() {
		var uid uint32
		var nickname string
		var damage float64
		if err := damageRows.Scan(&nickname, &uid, &damage); err != nil {
			logger.Error(fmt.Sprintf("扫描最强一击数据失败: %v", err))
			continue
		}
		damageData = append(damageData, map[string]interface{}{
			"uid":                 uid,
			"nickname":            nickname,
			"max_critical_damage": damage,
		})
	}
	result["damage"].(map[string]interface{})["data"] = damageData

	flightData := result["flight"].(map[string]interface{})["data"].([]map[string]interface{})
	for flightRows.Next() {
		var uid uint32
		var nickname string
		var distance float64
		if err := flightRows.Scan(&uid, &nickname, &distance); err != nil {
			logger.Error(fmt.Sprintf("扫描飞行距离数据失败: %v", err))
			continue
		}
		flightData = append(flightData, map[string]interface{}{
			"uid":              uid,
			"nickname":         nickname,
			"max_fly_distance": distance,
		})
	}
	result["flight"].(map[string]interface{})["data"] = flightData

	return result, nil
}

func (s *SyncService) GetLastUpdateTime(db *sql.DB) (int64, error) {
	var lastUpdate sql.NullInt64
	err := db.QueryRow("SELECT MAX(last_update_time) FROM t_player_stats").Scan(&lastUpdate)
	if err != nil {
		logger.Error("获取最后更新时间失败", err)
		return 0, err
	}

	// 如果没有数据或者值为NULL，返回当前时间戳
	if !lastUpdate.Valid {
		return time.Now().Unix(), nil
	}

	return lastUpdate.Int64, nil
}
