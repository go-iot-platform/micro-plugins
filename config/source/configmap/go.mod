module github.com/go-iot-platform/micro-plugins/config/source/configmap

go 1.13

require (
	github.com/go-iot-platform/go-micro v0.0.0-20201113111737-6edce0effdfa
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-00010101000000-000000000000
	k8s.io/apimachinery v0.0.0-00010101000000-000000000000
	k8s.io/client-go v0.0.0-00010101000000-000000000000
	k8s.io/utils v0.0.0-20200109141947-94aeca20bf09 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190708174958-539a33f6e817
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190708180123-608cd7da68f7
	k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190708175518-244289f83105
)

replace github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
