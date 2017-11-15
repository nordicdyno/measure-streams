#!/bin/bash
set -e
set -x
go build -o modeler .
echo "\$@=$@"
./modeler "$@"
rm modeler
