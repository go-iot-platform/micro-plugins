module github.com/go-iot-platform/micro-plugins/broker/grpc

go 1.13

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	google.golang.org/grpc v1.27.0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
