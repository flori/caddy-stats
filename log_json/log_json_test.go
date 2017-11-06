package log_json

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestLogJSON(t *testing.T) {
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Panicf("Could not create HTTP request: %v", err)
	}
	l := NewLogJSON(req)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(l.String()), &data)
	if err != nil {
		t.Panicf("Could not parse JSON output: %v", err)
	}
	if data["target"] != "/foo" {
		t.Errorf("wrong JSON output %s", l.String())
	}
}
