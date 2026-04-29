package utils

import "time"

func FmtDateTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.UTC().Format(time.RFC3339Nano)
}

func GenerateJWT(userID, name, email string) string {
	// 这里应该使用一个真正的 JWT 库来生成 token，这里只是一个示例
	return "mocked-jwt-token-for-" + userID
}
