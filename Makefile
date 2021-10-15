GOCMD = go
GOBIN = web
GOBUILD = $(GOCMD) build

build: deps
	$(GOCMD) mod tidy
	$(GOBUILD) -o $(GOBIN) main.go

build-linux:
	$(GOCMD) mod tidy
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(GOBIN) main.go

deps:
	export GO11MODULE=on
	export GOPROXY="https://goproxy.io,direct"

