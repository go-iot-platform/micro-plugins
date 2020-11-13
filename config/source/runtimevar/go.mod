module github.com/go-iot-platform/micro-plugins/config/source/runtimevar

go 1.13

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	gocloud.dev v0.17.0
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
)

replace github.com/Azure/go-autorest/autorest => github.com/azure/go-autorest/autorest v0.3.0

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
