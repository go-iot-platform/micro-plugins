#!/bin/bash -e

tag=$1

for m in $(find ./ -name 'go.mod'); do
  d=$(dirname $m);
  pushd $d;
  grep -q github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
  go mod tidy
  popd;
done
