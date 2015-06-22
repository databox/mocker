all: test

deps:
	# go get -u github.com/gorilla/mux
	# go get -u github.com/gorilla/pat
	go get -u github.com/julienschmidt/httprouter

test: deps
	go test ./... -v -p=1 -gcflags "-N -l"

build:
	go build ./...
