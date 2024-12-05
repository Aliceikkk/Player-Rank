package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SourceMySQL struct { //玩家存档所在的数据库
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"source_mysql"`
	TargetMySQL struct { //可以和玩家存档所在数据库一个位置 建议还是自建一个新的数据库
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"target_mysql"`
	SyncIntervalMinutes int    `json:"sync_interval_minutes"`
	ApiPort             int    `json:"api_port"`
	HttpPort            int    `json:"http_port"`
	ApiSecret           string `json:"api_secret"`
	Announcement        string `json:"announcement"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return &config, err
}
