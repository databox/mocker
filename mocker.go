package main
import (
	"fmt"
	"net/http"
	"io"
	"time"
	"errors"
	"encoding/json"
)

func onPanic(err error) {
	if err != nil {
		panic(err)
	}
}

type KPI struct {
	Key   string
	Value float32
	Date  string
}

type KPIWrap struct {
	Data []map[string]interface{} `json:"data"`
}

const AuthPushToken = "secrettoken"

func main() {
	fmt.Println("Mocker Server \\m/ on", time.Now(), "w/ PushToken =", AuthPushToken);

	http.HandleFunc("/", Push)
	http.HandleFunc("/lastpushes", LastPushes)

	onPanic(http.ListenAndServe(":1774", nil))
}


func prePushHandler(w *http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		onPanic(errors.New("Wrong method! Only POST goes with us!"))
		//TODO: Improve!
	}

	if req.Header.Get("Content-Type") != "application/json" {
		onPanic(errors.New("Wrong \"Content-Type\""))
		//TODO: Improve!
	}

	if pushToken, _, ok := req.BasicAuth(); ok {
		if pushToken != AuthPushToken {
			onPanic(errors.New(fmt.Sprintf("Wrong token \"%s\"", pushToken)))
			//TODO: Handle wrong token!
		}
	}
}

func Push(w http.ResponseWriter, req *http.Request) {
	prePushHandler(&w, req)
	io.WriteString(w, "Zdravo svet!\n")
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
