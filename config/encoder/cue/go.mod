module github.com/go-iot-platform/micro-plugins/config/encoder/cue

go 1.13

require (
	cuelang.org/go v0.0.15
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/ghodss/yaml v1.0.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
