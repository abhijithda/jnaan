TOP=$(shell pwd)

# Go build related variables
GOPATH=$(TOP)
GOSRC=$(GOPATH)/src
GOBIN=$(TOP)/bin

# GO_SRC_DIRS=$(sort $(dir $(shell cd $(GOSRC) && find . -name *.go)))
GO_SRC_DIRS := 0linkedlistloop
GO_SRC_DIRS += 1twosum

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
all: build test


.PHONY: build
build: .build_dep_go .build_go

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
analyze: .go_tools .gofmt .golint .govet

.PHONY: .golint
.golint:
	@echo Checking Go code for lint errors...
	$(GO_TOOLS_BIN)/golint -set_exit_status ./...

.PHONY: .gofmt
.gofmt:
	@echo Checking Go code for format errors...
	fmterrs=`gofmt -l . 2>&1`; \
	if [ "$$fmterrs" ]; then \
		echo "gofmt must be run on the following files:"; \
		echo "$$fmterrs"; \
		exit 1; \
	fi

.PHONY: .govet
.govet:
	@echo Vetting Go code for errors...
	go tool vet -all -shadow .


.PHONY: clean
clean: .clean_go

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

.PHONY: test
test: .unit_tests

.PHONY: .unit_tests
.unit_tests: .go_tools .test_go

.PHONY: .test_go
.test_go:
	@echo "Running Go unit tests...: $(GO_SRC_DIRS)"
	export GOPATH=$(GOPATH):$(GO_DEPS_PATH); \
	cd $(GOPATH)/src; \
	COVER_DIR=$(GOPATH)/cover; \
	rm -rf $${COVER_DIR}; \
	mkdir -p $${COVER_DIR}; \
	test_failed=0; \
	for d in $(GO_SRC_DIRS); do \
		echo "TEST_BEGIN: $${d}"; \
		go test -v --cover -covermode=count -coverprofile=$${COVER_DIR}/$$d.out $$d/... | \
			$(GO_TOOLS_BIN)/go-junit-report > TEST-$$d.xml; \
		echo "TEST_END($${ret}): $${d}"; \
	done ; \
	if [ $${test_failed} -ne 0 ]; then \
		@echo "Go unit tests failed."; \
		exit 1; \
	fi

.test_go_backup:
	@echo "Running Go unit tests...: $(GO_SRC_DIRS)"
	export GOPATH=$(GOPATH):$(GO_DEPS_PATH); \
	cd $(GOPATH)/src; \
	COVER_DIR=$(GOPATH)/cover; \
	rm -rf $${COVER_DIR}; \
	mkdir -p $${COVER_DIR}; \
	test_failed=0; \
	for d in $(GO_SRC_DIRS); do \
		echo "TEST_BEGIN: $$d"; \
		go test -v --cover -covermode=count -coverprofile=$${COVER_DIR}/$$d.out $$d/... | \
			$(GO_TOOLS_BIN)/go-junit-report > TEST-$$d.xml; \
		ret=$${PIPESTATUS[0]}; \
		if [ $${ret} -ne 0 ]; then \
			echo "Go unit test failed for $$d."; \
			test_failed=1; \
		fi ; \
		echo "TEST_END($${ret}): $$d"; \
		awk -f $(GOPATH)/utils/gocoverage-collate.awk $${COVER_DIR}/* > cover.out; \
		go tool cover -html=cover.out -o go-coverage-$$d.html; \
		$(GO_TOOLS_BIN)/gocov convert cover.out | $(GO_TOOLS_BIN)/gocov-xml > go-coverage-$$d.xml; \
		rm -rf $${COVER_DIR}/* cover.out; \
	done ; \
	if [ $${test_failed} -ne 0 ]; then \
		@echo "Go unit tests failed."; \
		exit 1; \
	fi

.NOTPARALLEL: