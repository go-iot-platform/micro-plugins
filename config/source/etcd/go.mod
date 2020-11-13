module github.com/go-iot-platform/micro-plugins/config/source/etcd

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200821141407-46a0a44f9539
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	go.uber.org/zap v1.16.0 // indirect
)
