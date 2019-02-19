TOP=$(shell pwd)

SUBDIRS := algos/golang

.PHONY: all
all: build test

.PHONY: build
build: 
	for d in $(SUBDIRS); do\
		make -C $$d build; \
	done

.PHONY: test
test:
	for d in $(SUBDIRS); do\
		make -C $$d test; \
	done
	