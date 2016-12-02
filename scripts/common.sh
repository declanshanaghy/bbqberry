#!/usr/bin/env bash

#OURGOFILES=($(go list -f '{{$p := .}}{{range $f := .GoFiles}}{{$p.Dir}}/{{$f}} {{end}} {{range $f := .TestGoFiles}}{{$p.Dir}}/{{$f}} {{end}}' ./... | xargs))
#echo $OURGOFILES

export OUR_PKGS=$(govendor list -no-status +local,^prog)
export OUR_DIRS="${OUR_PKGS//github.com\/declanshanaghy\/bbqberry\/}"
