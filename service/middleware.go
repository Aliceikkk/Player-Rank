package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SignatureMiddleware struct {
	apiSecret string
}

func NewSignatureMiddleware(apiSecret string) *SignatureMiddleware {
	return &SignatureMiddleware{
		apiSecret: apiSecret,
	}
}

func (m *SignatureMiddleware) Verify(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取签名和时间戳
		signature := r.URL.Query().Get("signature")
		timestamp := r.URL.Query().Get("timestamp")

		if signature == "" || timestamp == "" {
			http.Error(w, "Missing signature or timestamp", http.StatusBadRequest)
			return
		}

		// 验证时间戳是否在有效期内（例如5分钟）
		ts, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil || time.Now().Unix()-ts > 300 {
			http.Error(w, "Invalid or expired timestamp", http.StatusBadRequest)
			return
		}

		// 收集所有参数
		params := make(map[string]string)
		for key, values := range r.URL.Query() {
			if key != "signature" {
				params[key] = values[0]
			}
		}

		// 验证签名
		if !m.verifySignature(params, signature) {
			http.Error(w, "Invalid signature", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func (m *SignatureMiddleware) verifySignature(params map[string]string, signature string) bool {
	// 按字母顺序排序参数
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 构建签名字符串
	var signParts []string
	for _, k := range keys {
		signParts = append(signParts, fmt.Sprintf("%s=%s", k, params[k]))
	}
	signStr := strings.Join(signParts, "&") + "&key=" + m.apiSecret
	//logger.Info(fmt.Sprintf("Backend sign string: %s", signStr))

	hash := sha256.New()
	hash.Write([]byte(signStr))
	expectedSignature := hex.EncodeToString(hash.Sum(nil))
	//logger.Info(fmt.Sprintf("Expected signature: %s", expectedSignature))
	//logger.Info(fmt.Sprintf("Received signature: %s", signature))

	return signature == expectedSignature
}
