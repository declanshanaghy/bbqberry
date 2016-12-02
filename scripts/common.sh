#!/usr/bin/env bash

#OURGOFILES=($(go list -f '{{$p := .}}{{range $f := .GoFiles}}{{$p.Dir}}/{{$f}} {{end}} {{range $f := .TestGoFiles}}{{$p.Dir}}/{{$f}} {{end}}' ./... | xargs))
#echo $OURGOFILES

export OUR_PKGS=$(govendor list -no-status +local,^prog | grep -v client | grep -v restapi | grep -v models)
export OUR_DIRS="${OUR_PKGS//github.com\/declanshanaghy\/bbqberry\/}"
