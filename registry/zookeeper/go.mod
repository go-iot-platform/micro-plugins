module github.com/go-iot-platform/micro-plugins/registry/zookeeper

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/go-zookeeper/zk v1.0.2
	github.com/google/uuid v1.1.2
	github.com/mitchellh/hashstructure v1.0.0
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d
	github.com/smartystreets/goconvey v1.6.4
)
