module github.com/go-iot-platform/micro-plugins/codec/msgpackrpc

go 1.13

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/tinylib/msgp v1.1.1
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
