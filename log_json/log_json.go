package log_json

import (
	"net/http"
	"time"
)

type LogJSON struct {
	Time    time.Time
	Request *http.Request
}

// NewLogJSON creates a LogJSON object from the given HTTP request.
func NewLogJSON(request *http.Request) LogJSON {
	return LogJSON{Time: time.Now(), Request: request}
}

// Target returns the target of an request in the form of hostname:portnumber
// as a string.
func (l LogJSON) Target() string {
	return l.Request.Host + l.Request.URL.String()
}

// String returns the JSON representation of this LogJSON object as a string.
func (l LogJSON) String() string {
	return "{\"target\":\"" + l.Target() + "\",\"referer\":\"" +
		l.Request.Referer() + "\",\"time\":\"" + l.Time.Format(time.RFC3339) +
		"\"}"
}
