package util

import (
	"encoding/json"
	"github.com/maybaby/gscheduler/pkg/logging"
)

func ToMap(jsonString string) map[string]interface{} {
	if jsonString == "" {
		logging.Info("Empty String Found, Return \"\"")
		return nil
	}
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		logging.Error(err)
		return nil
	}
	return jsonMap
}
