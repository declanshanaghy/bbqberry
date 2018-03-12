#!/usr/bin/env bash

ftp -n <<EOF
open bbqberry-gaff
user pi berry
bin
put tmp/bin/bbqberry bbqberry
chmod 755 bbqberry
EOF