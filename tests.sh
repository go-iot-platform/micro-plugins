#!/bin/bash -e

set -x

mod="github.com/go-iot-platform/micro-plugins"
PKGS=""
for d in $(find * -name 'go.mod'); do
  pushd $(dirname $d) >/dev/null
  go mod download
  #go test -race -v ./... || :
  go test -v ./...
  popd >/dev/null
#  PKGS=" $PKGS ${mod}/$(dirname $d)"
done

#go test -race -v $PKGS || :
#go test -v $PKGS
