package di

import (
	"github.com/sado0823/go-kitx"
	"github.com/sado0823/go-kitx/kit/log"
	"github.com/sado0823/go-kitx/transport/grpc"
	"github.com/sado0823/go-kitx/transport/http"
)

type Base struct {
	ID      string
	Name    string
	Version string
}

func newApp(base *Base, logger log.Logger, gs *grpc.Server, hs *http.Server) *kitx.App {
	return kitx.New(
		kitx.WithID(base.ID),
		kitx.WithName(base.Name),
		kitx.WithVersion(base.Version),
		kitx.WithMetadata(map[string]string{}),
		kitx.WithLogger(logger),
		kitx.WithServer(
			gs,
			hs,
		),
	)
}
