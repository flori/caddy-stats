package log_json

import (
	"net/http"
	"time"
)

type LogJSON struct {
	Time    time.Time
	Request *http.Request
}

func NewLogJSON(request *http.Request) LogJSON {
	return LogJSON{Time: time.Now(), Request: request}
}

func (l LogJSON) Target() string {
	return l.Request.Host + l.Request.URL.String()
}

func (l LogJSON) String() string {
	return "{\"target\":\"" + l.Target() + "\",\"referer\":\"" +
		l.Request.Referer() + "\",\"time\":\"" + l.Time.Format(time.RFC3339) +
		"\"}"
}
