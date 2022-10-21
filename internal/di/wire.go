//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/sado0823/go-kitx/tpl/internal/conf"
	"github.com/sado0823/go-kitx/tpl/internal/dao"
	"github.com/sado0823/go-kitx/tpl/internal/server/grpc"
	"github.com/sado0823/go-kitx/tpl/internal/server/http"
	"github.com/sado0823/go-kitx/tpl/internal/service"

	"github.com/sado0823/go-kitx"
	"github.com/sado0823/go-kitx/kit/log"

	"github.com/google/wire"
)

// wireApp init kitx application.
func WireApp(*Base, *conf.Server, *conf.Data, log.Logger) (*kitx.App, func(), error) {
	panic(wire.Build(dao.ProviderSet, service.ProviderSet, http.ProviderSet, grpc.ProviderSet, newApp))
}
