#
# Makefile
# ykharuk, 2019-12-19 10:45
#

IMG ?= golang:1.13
VERSION ?= 0.0.1
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

# enable go modules, disabled CGO
#GOENV ?= GO111MODULE=on CGO_ENABLED=0
#export GO111MODULE=on
#export CGO_ENABLED=0
export GOPRIVATE="github.com/strato190/go-ktest-app"

# we build in a docker image, unless we are set to BUILD=local
GO ?= docker run --rm -v $(PWD):/app -w /app $(IMG) env $(GOENV)
ifeq ($(BUILD),local)
GO =
endif

GIT_COMMIT=$(shell git log -1 --pretty=format:"%H")

init:
	go mod init github.com/strato190/go-ktest-app

build: clean
	@go build -o go-ktest-app

build-docker: deps
	@docker build --build-arg VERSION=$(VERSION) --build-arg GIT_COMMIT=$(GIT_COMMIT) -t go-ktest-app:$(VERSION) -f Dockerfile .

deps:
	@go mod tidy
	@go mod verify

vendor:
	@go mod vendor

lint:
	golint ./...

fmt:
	gofmt -w $(GOFMT_FILES)

complex:
	gocyclo -avg .

clean: 
	rm -rf go-ktest-app
	go clean -i .