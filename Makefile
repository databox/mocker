all: test

deps:
	# go get -u github.com/gorilla/mux
	# go get -u github.com/gorilla/pat
	go get -u github.com/gorilla/handlers
	go get -u github.com/julienschmidt/httprouter
	go get -u github.com/elazarl/goproxy

test: deps
	go test ./*_test.go -v

build:
	go build ./...
