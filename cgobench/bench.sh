#!/bin/sh

set -xe

cd $(dirname $0)
go test -bench .
go test -bench . -gcflags '-l'
