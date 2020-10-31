module github.com/asim/nitro-plugins/proxy/grpc/v3

go 1.15

replace github.com/asim/nitro-plugins/client/grpc/v3 => ../../client/grpc

require (
	github.com/asim/nitro/v3 v3.3.0
	github.com/asim/nitro-plugins/client/grpc/v3 v3.0.4-0.20201022135331-c0b30a624ebb
	google.golang.org/grpc v1.31.0
	google.golang.org/grpc/examples v0.0.0-20201021230544-4e8458e5c638 // indirect
)