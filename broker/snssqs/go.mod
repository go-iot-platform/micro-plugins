module github.com/go-iot-platform/micro-plugins/broker/snssqs

go 1.13

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/aws/aws-sdk-go v1.28.4
	golang.org/x/text v0.3.3
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
