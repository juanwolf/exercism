.PHONY: all
all: python go rust

.PHONY: go
go:
	golangci-lint run ./go/...
	go test -v --bench --benchmem ./go/...

.PHONY: rust
rust:
	./build-rust.sh

.PHONY: python
python:
	black --check --diff ./python/
	pytest ./python/

.PHONY: elisp
elisp:
	cd elisp && ./run-tests.sh
