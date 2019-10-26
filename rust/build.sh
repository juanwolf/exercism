#!/bin/bash

for f in $(find . -name Cargo.toml -printf '%h\n' | sort -u); do
    pushd $f > /dev/null;
    cargo test;
    popd > /dev/null;
done
