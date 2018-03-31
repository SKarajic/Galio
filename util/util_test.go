package util

import (
	"testing"
	"encoding/json"
)

func TestGetJSON(t *testing.T) {
	data := make(map[string]interface{})
	content := GetJSON("https://httpbin.org/get", nil)

	err := json.Unmarshal([]byte(content), &data)

	if err != nil {
		t.Error(err)
	}

	if data["url"].(string) != "https://httpbin.org/get" {
		t.Errorf("Wrong URL received. Should be https://httpbin.org/get, got %s", data["url"].(string))
	}
}
