.DEFAULT_GOAL := all
SHELL := /bin/bash

all: install

CGO_ENABLED ?= 0
GO111MODULE ?= on
GOFLAGS ?= "-mod=mod"
OUTBIN ?= goodwe

NAME := $(shell basename $(PWD))

VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_TIME ?= $(shell /bin/date +%FT%T%z)
PKG ?= $(shell go list -m | grep $(NAME))
LD_FLAGS ?= "-s -w"
GO_OPTS ?= -trimpath
OPTIMIZE ?= false

version:
	@echo $(NAME) $(VERSION) $(BUILD_TIME) $(PKG) $(LD_FLAGS)


.PHONY: all-build
all-build: build-linux-amd64 build-linux-arm64 build-linux-386

.PHONY: download
download:
	@echo Download go.mod dependencies
	go mod download


.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) go build \
    	-ldflags $(LD_FLAGS) $(GO_OPTS) \
    	-o bin/$(OUTBIN)_$(GOOS)_$(GOARCH) \
    	./cmd/$(OUTBIN)/main.go
	@if [ "$(OPTIMIZE)" = "true" ]; then \
		upx --best bin/$(OUTBIN)_$(GOOS)_$(GOARCH) ; \
	fi

.PHONY: install
install:
	CGO_ENABLED=$(CGO_ENABLED) go install \
    	-ldflags $(LD_FLAGS)  \
    	./...

build-%:
	@$(MAKE) build                        \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

build-linux-arm64:
	@$(MAKE) build-linux_arm64 --no-print-directory

build-linux-amd64:
	@$(MAKE) build-linux_amd64 --no-print-directory

build-linux-386:
	@$(MAKE) build-linux_386 --no-print-directory

.PHONY: prepush
prepush: vet fmt tidy test

.PHONY: test
test:
	CGO_ENABLED=0 go build ./...
	CGO_ENABLED=1 go test -race -cover -v -mod=readonly ./... && echo -e "\033[32mSUCCESS\033[0m" || (echo -e "\033[31mFAILED\033[0m" && exit 1)

.PHONY: bench
bench:
	go test -test.timeout=30m -benchmem -run ^$$ -benchtime=20s -bench . ./... && echo -e "\033[32mSUCCESS\033[0m" || (echo -e "\033[31mFAILED\033[0m" && exit 1)

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: go_list
go_list:
	go list -u -m all

.PHONY: go_update_all
go_update_all:
	go get -t -u ./...
