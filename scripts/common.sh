#!/usr/bin/env bash

export OUR_PKGS=$(govendor list -no-status +local,^prog | grep -v client | grep -v restapi | grep -v models | grep -v mock | grep -v test)
export OUR_DIRS="${OUR_PKGS//github.com\/declanshanaghy\/bbqberry\/}"
