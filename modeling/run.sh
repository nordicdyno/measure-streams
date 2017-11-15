#!/bin/bash
set -e
set -x
go build -o modeler github.com/nordicdyno/measure-streams/modeling
echo "\$@=$@"
./modeler "$@"
