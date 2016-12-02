#!/usr/bin/env bash

source ./scripts/common.sh

go vet -x $OUR_PKGS

for pkg in $OUR_PKGS
do
    golint -set_exit_status $pkg
    result=$?
	if [ "$result" != "0" ]; then
	    exit $result
	fi
done
