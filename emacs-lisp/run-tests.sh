#!/bin/bash

for d in $(find . -maxdepth 1 -mindepth 1 -type d | sort -u); do
    pushd $d > /dev/null;
    emacs -batch -l ert -l "$d-test.el" -f ert-run-tests-batch-and-exit
    popd > /dev/null;
done
