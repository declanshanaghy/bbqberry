#!/usr/bin/env bash -x
set -e

source ./scripts/common.sh

gofmt -s -w $OUR_DIRS
goimports -w $OUR_DIRS