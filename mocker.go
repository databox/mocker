package main
import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

func main() {
	fmt.Println("Mocker Server \\m/")

	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":7722", router))

	fmt.Println("Done.")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Ellow in Index!")
}
