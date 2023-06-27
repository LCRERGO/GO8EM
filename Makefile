VERSION = 1.0.0

BINNAME = go8em
BUILDDIR = build
GOCMD = env go
GOENVS = GOOS=js GOARCH=wasm

all: main

main:
	mkdir -p $(BUILDDIR)
	$(GOCMD) build -o $(BUILDDIR)/$(BINNAME) cmd/main.go

check:
	$(GOCMD) test ./...

check/all:
	TEST_ALL=1  $(GOCMD) test -count=1 ./...

clean:
	rm -rf $(BUILDDIR)

install:

uninstall:


.PHONY: all main check clean install uninstall
