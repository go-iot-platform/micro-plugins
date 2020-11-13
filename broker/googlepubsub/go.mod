module github.com/go-iot-platform/micro-plugins/broker/googlepubsub

go 1.13

require (
	cloud.google.com/go/pubsub v1.1.0
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/google/uuid v1.1.2
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
	google.golang.org/api v0.15.0
	google.golang.org/grpc v1.27.0
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
