package main
import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)

type KPI struct {
	Key   string
	Value float32
	Date  string
}

type KPIWrap struct {
	Data []map[string]interface{} `json:"data"`
}

type APIError struct {
	Status string `json:"status"`
}

const (
	PushToken = "push_token"
	ServerBind = ":1447"
)

func main() {
	fmt.Println("Mocker Server \\m/ on", time.Now(),
		"w/ PushToken =", PushToken,
		"w/ ServerBind =", ServerBind,
	);

	http.HandleFunc("/", Push)
	http.HandleFunc("/lastpushes", LastPushes)

	if err := http.ListenAndServe(ServerBind, securityWrap(http.DefaultServeMux)); err != nil {
		panic(err)
	}
}

func securityWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		shouldFail := false

		if r.Header.Get("Content-Type") != "application/json" {
			json.NewEncoder(w).Encode(&APIError{Status: "Error: Content type should be application/json"})
			shouldFail = true
		}

		//TODO: Something should be done in API_node about this.
		if !shouldFail && r.Method != "POST" {
			json.NewEncoder(w).Encode(&APIError{Status: "Only POST goes with us,..."})
			shouldFail = true
		}

		if pushToken, _, ok := r.BasicAuth(); !shouldFail && ok && pushToken != PushToken {
			json.NewEncoder(w).Encode(&APIError{Status: "Error: No authorization sent"})
			shouldFail = true
		}

		if !shouldFail {
			h.ServeHTTP(w, r)
		}
	})
}

func Push(w http.ResponseWriter, req *http.Request) {
	if err := json.NewEncoder(w).Encode(&APIError{Status:"ok"}); err != nil {
		panic(err)
	}
}

func LastPushes(w http.ResponseWriter, req *http.Request) {
	p := KPI{
		Key:"temp.ly",
		Value: 100.3,
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
