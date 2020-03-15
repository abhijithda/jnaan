# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.DEFAULT_GOAL := help
.PHONY: help
help:	## Display this help message.
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

TOP=$(shell pwd)

SUBDIRS := lcgo algos/golang

.PHONY: all
all: build test		## Perform build & test

.PHONY: build
build: 	## Build source code
	for d in $(SUBDIRS); do\
		make -C $$d build; \
	done

.PHONY: clean
clean:	## Clean build artifacts
	for d in $(SUBDIRS); do\
		make -C $$d clean; \
	done
	
.PHONY: test
test:	## Test source code
	for d in $(SUBDIRS); do\
		make -C $$d test; \
	done
	