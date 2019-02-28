# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.DEFAULT_GOAL := help
.PHONY: help
help:	## Display this help message.
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

TOP=$(shell pwd)

# Go build related variables
GOPATH=$(TOP)
GOSRC=$(GOPATH)/src
GOBIN=$(TOP)/bin

# GO_SRC_DIRS=$(sort $(dir $(shell cd $(GOSRC) && find . -name *.go)))
GO_SRC_DIRS := $(shell \
	cd $(GOSRC) && find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
# GO_SRC_DIRS := 0linkedlistloop
# GO_SRC_DIRS += 1twosum

GO_DEPS_PATH=$(TOP)/.godeps
GO_DEPS_BIN=$(GO_TOOLS_PATH)/bin
GO_SRC_DEPS := gopkg.in/yaml.v2

GO_TOOLS_PATH=$(TOP)/.gotools
GO_TOOLS_BIN=$(GO_TOOLS_PATH)/bin
GO_TOOLS := github.com/axw/gocov/gocov
GO_TOOLS += github.com/AlekSi/gocov-xml
GO_TOOLS += github.com/jstemmer/go-junit-report
GO_TOOLS += golang.org/x/lint/golint
GO_TOOLS += github.com/golang/lint/golint


.PHONY: all
all: build test		## Perform build & test


.PHONY: build
build: .build_dep_go .build_go	## Build source code

.PHONY: .build_dep_go
.build_dep_go:
	# Since go build determines and build only updated sources, no need to run clean all go binaries before running build_go
	@echo "Getting any Go source dependencies...: ";
	export GOPATH=$(GO_DEPS_PATH); \
	export GOBIN=$(GO_DEPS_BIN); \
	for d in $(GO_SRC_DEPS); do \
		echo "Getting dependency: $$d..."; \
		go get -insecure $$d; \
	done;

.PHONY: .build_go
.build_go:
	# Since go build determines and build only updated sources, no need to run clean all go binaries before running build_go
	@echo "Building Go binaries...: $(GO_SRC_DIRS) ";
	export GOPATH=$(GOPATH):$(GO_DEPS_PATH); \
	export GOBIN=$(GOBIN); \
	cd $(GOSRC); \
	build_failed=0; \
	for d in $(GO_SRC_DIRS); do \
		echo "BUILD_BEGIN: $$d"; \
		go install $$d; \
		ret=$$?; \
		if [ $${ret} -ne 0 ]; then \
			echo "Failed to build: $$d."; \
			build_failed=1; \
		fi ; \
		echo "BUILD_END($${ret}): $$d"; \
	done ; \
	if [ $${build_failed} -ne 0 ]; then \
		@echo "Failed to build Go binaries."; \
		exit 1; \
	fi

.PHONY: .go_tools
.go_tools:
	@echo "Getting required Go tools...: ";
	export GOPATH=$(GO_TOOLS_PATH); \
	export GOBIN=$(GO_TOOLS_BIN); \
	for d in $(GO_TOOLS); do \
		echo "Getting: $$d..."; \
		go get -insecure $$d; \
	done;


.PHONY: analyze
analyze: .go_tools gofmt golint govet go-race  ## Analyze source code for different errors through gofmt, golint, govet, go-race

.PHONY: golint
golint:	## golint
	@echo Checking Go code for lint errors...
	$(GO_TOOLS_BIN)/golint -set_exit_status ./...

.PHONY: gofmt
gofmt:	## gofmt
	@echo Checking Go code for format errors...
	fmterrs=`gofmt -l . 2>&1`; \
	if [ "$$fmterrs" ]; then \
		echo "gofmt must be run on the following files:"; \
		echo "$$fmterrs"; \
		exit 1; \
	fi

.PHONY: govet
govet:	## Vet Go source code for errors
	@echo Vetting Go code for errors...
	go tool vet -all -shadow .


.PHONY: clean
clean: .clean_go	## clean

.PHONY: .clean_go
.clean_go:
	@echo "Cleaning Go binaries...: $(GO_SRC_DIRS) ";
	export GOPATH=$(GOPATH); \
	export GOBIN=$(GOBIN); \
	cd $(GOSRC); \
	for d in $(GO_SRC_DIRS); do \
		go clean -i $$d; \
		rm -f $$d/*.log; \
	done;

.PHONY: .clean_dep_go
.clean_dep_go:
	# NOTE: Not cleaning dependencies as part of `clean` target, as otherwise
	# 		it will need to be downloaded everytime while building locally.
	@echo "Cleaning Go source and test dependencies...: "; \
	export GOBIN=$(GOBIN); \
	cd $(GOSRC); \
	for d in $(GO_SRC_DEPS); do \
		go clean -i; \
		echo "Deleting $$d..."; set -x;\
		rm -rfv `dirname $$d`; \
	done;
	@rm -rf $(GO_TOOLS_PATH)

.PHONY: show
show:	## show source directories
	@echo "SRC  = $(GO_SRC_DIRS)"

.PHONY: test
test: .unit_tests	## Run all tests

.PHONY: .unit_tests
.unit_tests: .go_tools test_go

.PHONY: test_go
test_go:	## Run go tests
	@echo "Running Go unit tests...: $(GO_SRC_DIRS)"
	export GOPATH=$(GOPATH):$(GO_DEPS_PATH); \
	cd $(GOPATH)/src; \
	echo "" > coverage.txt; \
	for d in $(GO_SRC_DIRS); do \
		echo "TEST_BEGIN: $${d}"; \
		go test -v --cover -covermode=count -coverprofile=profile.out $${d}/...; \
		ret=$$?; \
        cat profile.out >> coverage.txt; \
        rm profile.out; \
		echo "TEST_END($${ret}): $${d}"; \
	done;

.NOTPARALLEL: