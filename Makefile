VERSION = 1.0.0

BUILDDIR = build
GOCMD = env go
GOENVS = GOOS=js GOARCH=wasm

all: main

main:
	mkdir -p $(BUILDDIR)
	$(GOCMD) build -o $(BUILDDIR)/main-$(VERSION) cmd/main.go

check:
	$(GOCMD) test ./...

check/all:
	GOCACHE=off TEST_ALL=1  $(GOCMD) test ./...

clean:
	rm -rf $(BUILDDIR)

install:

uninstall:


.PHONY: all main check clean install uninstall