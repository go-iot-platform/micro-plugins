package grpc

import (
	"crypto/tls"

	"github.com/go-iot-platform/go-micro/service"
	gc "github.com/go-iot-platform/micro-plugins/client/grpc"
	gs "github.com/go-iot-platform/micro-plugins/server/grpc"
)

// WithTLS sets the TLS config for the service
func WithTLS(t *tls.Config) service.Option {
	return func(o *service.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}
