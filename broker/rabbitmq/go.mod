module github.com/go-iot-platform/micro-plugins/broker/rabbitmq

go 1.13

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/google/uuid v1.1.2
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
