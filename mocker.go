package main
import (
	"fmt"
	"net/http"
	_ "github.com/julienschmidt/httprouter"
	_ "github.com/gorilla/handlers"
	"log"
	"github.com/elazarl/goproxy"
	_ "regexp"
	_ "bufio"
	_ "net"
)

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Mocker Server \\m/")

	// router := httprouter.New()
	// router.GET("/", Index)
	// log.Fatal(http.ListenAndServe(":7722", router))

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true


	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(func (req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		fmt.Println("req", req)
		return req, nil
	})

	// proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*:80$"))).HijackConnect(func(req *http.Request, client net.Conn, ctx *goproxy.ProxyCtx) {
	// 	fmt.Println("Got req", req)
	// 	log.Println("Got request", req)
    //
	// 	defer func() {
	// 		if e := recover(); e != nil {
	// 			ctx.Logf("error connecting to remote: %v", e)
	// 			client.Write([]byte("HTTP/1.1 500 Cannot reach destination\r\n\r\n"))
	// 		}
	// 		client.Close()
	// 	}()
    //
	// 	clientBuf := bufio.NewReadWriter(bufio.NewReader(client), bufio.NewWriter(client))
    //
	// 	remote, err := net.Dial("tcp", req.URL.Host)
	// 	orPanic(err)
    //
	// 	remoteBuf := bufio.NewReadWriter(bufio.NewReader(remote), bufio.NewWriter(remote))
    //
	// 	for {
	// 		req, err := http.ReadRequest(clientBuf.Reader)
	// 		orPanic(err)
	// 		orPanic(req.Write(remoteBuf))
	// 		orPanic(remoteBuf.Flush())
	// 		resp, err := http.ReadResponse(remoteBuf.Reader, req)
	// 		orPanic(err)
	// 		orPanic(resp.Write(clientBuf.Writer))
	// 		orPanic(clientBuf.Flush())
	// 	}
	// })

	log.Fatal(http.ListenAndServe(":8080", proxy))

	fmt.Println("Done.")
}
