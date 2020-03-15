# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.DEFAULT_GOAL := help
.PHONY: help
help:	## Display this help message.
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

TOP=$(shell pwd)

.PHONY: all
all: build test		## Perform build & test

.PHONY: build
build:	## Build source code
	# As go build determines and builds only updated sources, 
	# no need to run clean all go binaries before running build
	@echo "Building Go binaries...";
	go build -v ./...; \
	ret=$$?; \
	if [ $${ret} -ne 0 ]; then \
		@echo "Failed to build Go binaries."; \
		exit 1; \
	fi

.PHONY: analyze
analyze: gofmt govet go-race  ## Analyze source code for different errors through gofmt, golint, govet, go-race

.PHONY: gofmt
gofmt:	## gofmt
	@echo Checking Go code for format errors...
	fmterrs=`gofmt -l . 2>&1`; \
	if [ "$$fmterrs" ]; then \
		echo "gofmt must be run on the following files:"; \
		echo "$$fmterrs"; \
		exit 1; \
	fi

go-race:	## gorace
	@echo Checking for any race conditions...
	go test -mod=vendor -v -race ./...;

.PHONY: govet
govet:	## Vet Go source code for errors
	@echo Vetting Go code for errors...
	go vet -all ./...


.PHONY: clean
clean:	## clean
	@echo "Cleaning Go binaries...: $(GO_SRC_DIRS) ";
	go clean -i ./...;
	rm -f */*.log;
	rm -f coverage.txt;

.PHONY: test
test:	## Run all tests
	@echo "Running Go unit tests..."
	go test -v --cover -covermode=count -coverprofile=coverage.txt ./...;
	# /usr/bin/sed -i -e s%`go list -m all`/%%g coverage.txt

.NOTPARALLEL: