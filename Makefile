VERSION = 1.0.0

BUILDDIR = build
GOCMD = go
GOENVS = GOOS=js GOARCH=wasm

all: main

main:
	mkdir -p $(BUILDDIR)
	$(GOCMD) build -o $(BUILDDIR)/main-$(VERSION) cmd/main.go

test:
	$(GOCMD) test ./...

test/all:
	GOCACHE=off TEST_ALL=1  $(GOCMD) test ./...

clean:
	rm -rf $(BUILDDIR) $(WEBDIR)/js/main.wasm

install:

uninstall:


.PHONY: all main test clean install uninstall