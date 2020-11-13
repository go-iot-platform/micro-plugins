module github.com/go-iot-platform/micro-plugins/store/mysql

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/pkg/errors v0.9.1
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
