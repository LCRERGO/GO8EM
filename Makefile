VERSION = 1.0.0

BUILDDIR = build
GOCMD = go
GOENVS = GOOS=js GOARCH=wasm

all: main

main:
	mkdir -p $(BUILDDIR)
	$(GOCMD) build -o $(BUILDDIR)/main-$(VERSION) cmd/main.go

clean:
	rm -rf $(BUILDDIR) $(WEBDIR)/js/main.wasm

install:

uninstall:
