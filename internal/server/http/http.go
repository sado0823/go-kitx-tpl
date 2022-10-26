package http

import (
	"context"

	"github.com/google/wire"

	"github.com/sado0823/go-kitx/kit/log"
	"github.com/sado0823/go-kitx/kit/tracing"
	v1 "github.com/sado0823/go-kitx/tpl/api/helloworld/v1"
	"github.com/sado0823/go-kitx/tpl/internal/conf"
	"github.com/sado0823/go-kitx/tpl/internal/service"
	"github.com/sado0823/go-kitx/transport/http"
	"github.com/sado0823/go-kitx/transport/pbchain"
)

var ProviderSet = wire.NewSet(NewServer)

func NewServer(c *conf.Server, svc *service.Service, logger log.Logger) *http.Server {
	var (
		opts = []http.ServerOption{
			http.WithServerPBChain(
				pbchain.Recovery(),
				tracing.Server(),
				pbchain.LoggingServer(),
			),
		}
		matches = []bool{
			c.Http.Network != "",
			c.Http.Addr != "",
			c.Http.Timeout != nil,
		}
		match2Opts = []http.ServerOption{
			http.WithServerNetwork(c.Http.Network),
			http.WithServerAddress(c.Http.Addr),
			http.WithServerTimeout(c.Http.Timeout.AsDuration()),
		}
	)

	for optIndex, match := range matches {
		if match {
			opts = append(opts, match2Opts[optIndex])
		}
	}

	srv := http.NewServer(opts...)
	router := srv.Route("")

	// register from grpc
	v1.RegisterGreeterHTTPServer(router, svc)

	// registry from custom
	router.GET("/ping", _ping(svc))

	return srv
}

func _ping(svc *service.Service) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in = make(map[string]interface{})
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}

		// can not remove this handler wrapper !!! or http pbchain would be invalid
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			// do service logic
			v := req.(map[string]interface{})
			v["wrapper"] = "pong"
			return v, nil
		})

		out, err := h(ctx, in)
		if err != nil {
			return err
		}

		return ctx.JSON(200, out)
	}
}
