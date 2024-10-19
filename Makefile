# Makefile for Go project

BINARY_NAME=ourbible
BASEDIR ?= $(PWD)
PREFIX ?= /usr/local

# Check the operating system
ifeq ($(OS),Windows_NT)
    BINARY_NAME=$(BINARY_NAME).exe
    BUILD_COMMAND=go build -o build/$(BINARY_NAME) -ldflags "-H=windowsgui" ./cmd
else
    BUILD_COMMAND=go build -o build/$(BINARY_NAME) ./cmd
endif

# Default target
.PHONY: all
all: build

# Build target
.PHONY: build
build:
	go mod tidy
	$(BUILD_COMMAND)

# Install target
.PHONY: install
install: build
	install -m 755 $(BINARY_NAME) $(PREFIX)/bin/
	install -m 755 build/$(BINARY_NAME) $(BASEDIR)/$(PREFIX)/bin/$(BINARY_NAME)
	install -d $(BASEDIR)/usr/share/applications
  install -d $(BASEDIR)/$(PREFIX)/$(BINARY_NAME)/static
  install -d $(BASEDIR)/$(PREFIX)/$(BINARY_NAME)/database
  mv "static/"* "$(BASEDIR)/$(PREFIX)/$(BINARY_NAME)/static"
  mv "database/"* "$(BASEDIR)/$(PREFIX)/$(BINARY_NAME)/database"
  mv "$(BINARY_NAME).desktop" "$(BASEDIR)/usr/share/applications"

# Clean target
.PHONY: clean
clean:
	rm -f build/$(BINARY_NAME)
