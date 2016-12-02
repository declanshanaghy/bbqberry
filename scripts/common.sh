#!/usr/bin/env bash

pkg_root=$(cat ./vendor/vendor.json | grep rootPath | awk '{split($0,array,": ")} END{print array[2]}' | awk '{split($0,array,"\"")} END{print array[2]}')

export OUR_PKGS=$(govendor list -no-status +local,^prog | grep -v client | grep -v restapi | grep -v models | grep -v mock | grep -v test)
export OUR_DIRS="${OUR_PKGS//$pkg_root\/}"

