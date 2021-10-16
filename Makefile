GOCMD = go
GOBIN = web
GOBUILD = $(GOCMD) build

build: deps swag
	$(GOCMD) mod tidy
	$(GOBUILD) -o $(GOBIN) main.go

build-linux:
	$(GOCMD) mod tidy
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(GOBIN) main.go

deps:
	export GO11MODULE=on
	export GOPROXY="https://goproxy.io,direct"

swag: 
	mkdir -p docs
	go get -u github.com/swaggo/swag/cmd/swag
	swag init

clean:
	rm -rf docs/ web
