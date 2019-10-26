.PHONY: all
all: python go rust

.PHONY: go
go:
	golangci-lint run ./go/...
	go test -v ./go/...
	go test -v ./go/... --bench ".*" --benchmem ./go/...

.PHONY: rust
rust:
	cd rust && ./build.sh

.PHONY: python
python:
	black --check --diff --exclude=".*_test.py" ./python/
	pytest ./python/

.PHONY: elisp
elisp:
	cd elisp && ./run-tests.sh
