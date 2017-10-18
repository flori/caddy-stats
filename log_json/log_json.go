package log_json

import (
	"encoding/json"
	"log"
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

func (l LogJSON) Map() map[string]string {
	return map[string]string{
		"target":  l.Target(),
		"referer": l.Request.Referer(),
		"time":    l.Time.Format(time.RFC3339),
	}
}

// String returns the JSON representation of this LogJSON object as a string.
func (l LogJSON) String() string {
	t, err := json.Marshal(l.Map())
	if err != nil {
		log.Fatal(err)
	}
	return string(t)
}
