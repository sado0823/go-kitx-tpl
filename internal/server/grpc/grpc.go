package grpc

import (
	"github.com/google/wire"

	"github.com/sado0823/go-kitx/kit/log"
	"github.com/sado0823/go-kitx/transport/pbchain"

	v1 "github.com/sado0823/go-kitx/tpl/api/helloworld/v1"
	"github.com/sado0823/go-kitx/tpl/internal/conf"
	"github.com/sado0823/go-kitx/tpl/internal/service"
	"github.com/sado0823/go-kitx/transport/grpc"
)

var ProviderSet = wire.NewSet(NewServer)

func NewServer(c *conf.Server, svc *service.Service, logger log.Logger) *grpc.Server {
	var (
		opts = []grpc.ServerOption{
			grpc.WithServerPBChain(
				pbchain.Recovery(),
				pbchain.LoggingServer(logger),
			),
		}
		matches = []bool{
			c.Grpc.Network != "",
			c.Grpc.Addr != "",
			c.Grpc.Timeout != nil,
		}
		match2Opts = []grpc.ServerOption{
			grpc.WithServerNetwork(c.Grpc.Network),
			grpc.WithServerAddress(c.Grpc.Addr),
			grpc.WithServerTimeout(c.Grpc.Timeout.AsDuration()),
		}
	)

	for optIndex, match := range matches {
		if match {
			opts = append(opts, match2Opts[optIndex])
		}
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, svc)
	return srv
}
