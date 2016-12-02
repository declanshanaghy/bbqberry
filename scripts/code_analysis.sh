#!/usr/bin/env bash
set -e

source ./scripts/common.sh

misspell -error $OUR_DIRS

ineffassign .

gocyclo -over 10 $OUR_DIRS

go vet $OUR_PKGS

for pkg in $OUR_PKGS
do
    golint -set_exit_status $pkg
done
