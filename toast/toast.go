package toast

import (
	"encoding/json"
)

const (
	ERROR   = "ERROR"
	SUCCESS = "SUCCESS"
)

type ToastType struct {
	Level   string `json:"level"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func Jsonify(level, title, message string) string {
	eventMap := map[string]ToastType{}
	eventMap["showToast"] = ToastType{
		Level:   level,
		Title:   title,
		Message: message,
	}

	jsonData, _ := json.Marshal(eventMap)
	return string(jsonData)
}
