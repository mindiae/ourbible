# Makefile for Go project

BINARY_NAME=ourbible
BASEDIR ?= $(PWD)
PREFIX ?= /usr/local

# Check the operating system
ifeq ($(OS),Windows_NT)
    BINARY_NAME=$(BINARY_NAME).exe
    BUILD_COMMAND=go build -o build/$(BINARY_NAME) main.go
else
    BUILD_COMMAND=go build -o build/$(BINARY_NAME) main.go
endif

# Default target
.PHONY: all
all: build

# Build target
.PHONY: build
build:
	$(BUILD_COMMAND)

# Install target
.PHONY: install
install: build
	install -m 755 $(BINARY_NAME) $(PREFIX)/bin/

# Clean target
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
