.PHONY: all
all: python go rust

.PHONY: go
go:
	go test -v --bench --benchmem ./go/...

.PHONY: rust
rust:
	./build-rust.sh

.PHONY: python
python:
	pytest ./python/

.PHONY: elisp
elisp:
	cd elisp && ./run-tests.sh
