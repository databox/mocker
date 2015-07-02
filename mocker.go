package main
import (
	"fmt"
	"net/http"
	"io"
	"time"
	"errors"
)

func onPanic(err error) {
	if err != nil {
		panic(err)
	}
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

func PushServer(w http.ResponseWriter, req *http.Request) {
	prePushHandler(&w, req)

	io.WriteString(w, "Zdravo svet!\n")
}

const AuthPushToken = "secrettoken"

func main() {
	fmt.Println("Mocker Server \\m/ on", time.Now(), "w/ PushToken =", AuthPushToken);

	http.HandleFunc("/", PushServer)

	onPanic(http.ListenAndServe(":1774", nil))
}
