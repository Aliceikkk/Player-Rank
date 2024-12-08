package model

import (
    "crypto/sha256"
    "encoding/hex"
    "time"
)

type Admin struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"-"` // 不输出密码
    CreatedAt time.Time `json:"created_at"`
}

func HashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
} 