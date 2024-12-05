package model

type PlayerStats struct {
	UID              uint32  `json:"uid"`
	Nickname         string  `json:"nickname"`
	MaxCriticalDmg   float32 `json:"max_critical_damage"`
	MaxFlyMapDist    float32 `json:"max_fly_map_distance"`
	LastUpdateTime   int64   `json:"last_update_time"`
} 