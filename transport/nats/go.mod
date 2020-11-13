module github.com/go-iot-platform/micro-plugins/transport/nats

go 1.13

require (
	github.com/go-log/log v0.2.0
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/nats-io/nats.go v1.9.2
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
