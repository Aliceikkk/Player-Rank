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

	// 连接源数据库
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
	rows, err := sourceDB.Query(fmt.Sprintf("SELECT uid, bin_data FROM %s", tableName))
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	updated := 0
	for rows.Next() {
		if err := s.processRow(rows, targetDB); err != nil {
			//logger.Error(fmt.Sprintf("Failed to process row for table: %s", tableName), err)
			continue
		}
		updated++
	}
	return updated, rows.Err()
}

func (s *SyncService) processRow(rows *sql.Rows, targetDB *sql.DB) error {
	var uid uint32
	var binData []byte
	if err := rows.Scan(&uid, &binData); err != nil {
		return err
	}

	unzlibedData, err := UnzlibData(binData)
	if err != nil {
		//logger.Error(fmt.Sprintf("Failed to unzlib data for uid %d", uid), err)
		return err
	}

	var playerDataBin pb.PlayerDataBin
	if err := proto.Unmarshal(unzlibedData, &playerDataBin); err != nil {
		//logger.Error(fmt.Sprintf("Failed to unmarshal PlayerDataBin for uid %d", uid), err)
		return err
	}

	basicBin := playerDataBin.GetBasicBin()
	nickname := fmt.Sprintf("Player_%d", uid)
	if basicBin != nil {
		nickname = string(basicBin.GetNickname())
		//logger.Info(fmt.Sprintf("Found BasicBin for uid %d: %+v", uid, map[string]interface{}{
		//	"nickname": nickname,
		//}))
	}

	watcherBin := playerDataBin.GetWatcherBin()
	if watcherBin == nil {
		//logger.Info(fmt.Sprintf("No WatcherBin found for uid %d", uid))
		return nil
	}

	recordValue := watcherBin.GetRecordValue()
	if recordValue == nil {
		//logger.Info(fmt.Sprintf("No RecordValue found for uid %d", uid))
		return nil
	}

	maxCriticalDamage := recordValue.GetMaxCriticalDamage()
	maxFlyMapDistance := recordValue.GetMaxFlyMapDistance()

	//logger.Info(fmt.Sprintf("Found record values for uid %d: %+v", uid, map[string]interface{}{
	//	"nickname":            nickname,
	//	"max_critical_damage": fmt.Sprintf("%.2f", maxCriticalDamage),
	//	"max_fly_distance":    fmt.Sprintf("%.2f", maxFlyMapDistance),
	//}))

	if maxCriticalDamage == 0 && maxFlyMapDistance == 0 {
		//logger.Info(fmt.Sprintf("Skipping record with no relevant data for uid %d", uid))
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
		//logger.Error(fmt.Sprintf("Failed to update stats for uid %d", uid), err)
		return err
	}

	//logger.Info(fmt.Sprintf("Successfully updated player stats for uid %d: %s", uid, nickname))
	return nil
}

// 添加新的方法来获取排行榜数据
func (s *SyncService) GetLeaderboards(db *sql.DB, page, pageSize int) (map[string]interface{}, error) {
	offset := (page - 1) * pageSize

	// 获取最强一击排行榜
	damageRows, err := db.Query(`
		SELECT uid, nickname, max_critical_damage 
		FROM t_player_stats 
		ORDER BY max_critical_damage DESC 
		LIMIT ? OFFSET ?
	`, pageSize, offset)
	if err != nil {
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
		return nil, err
	}
	defer flightRows.Close()

	// 获取总记录数
	var totalDamage, totalFlight int
	err = db.QueryRow("SELECT COUNT(*) FROM t_player_stats WHERE max_critical_damage > 0").Scan(&totalDamage)
	if err != nil {
		return nil, err
	}
	err = db.QueryRow("SELECT COUNT(*) FROM t_player_stats WHERE max_fly_map_distance > 0").Scan(&totalFlight)
	if err != nil {
		return nil, err
	}

	// 获取最后更新时间
	lastUpdate, err := s.GetLastUpdateTime(db)
	if err != nil {
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
		if err := damageRows.Scan(&uid, &nickname, &damage); err != nil {
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
	var lastUpdate int64
	err := db.QueryRow("SELECT MAX(last_update_time) FROM t_player_stats").Scan(&lastUpdate)
	if err != nil {
		return 0, err
	}
	return lastUpdate, nil
}
