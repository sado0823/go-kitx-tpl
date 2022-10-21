package service

import (
	"context"

	"github.com/sado0823/go-kitx/errorx"
	"github.com/sado0823/go-kitx/kit/log"
	v1 "github.com/sado0823/go-kitx/tpl/api/helloworld/v1"
	"github.com/sado0823/go-kitx/tpl/internal/conf"
	"github.com/sado0823/go-kitx/tpl/internal/dao"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(New)

type Service struct {
	conf *conf.Data
	dao  *dao.Dao
	log  *log.Helper

	v1.UnimplementedGreeterServer
}

func New(d *dao.Dao, c *conf.Data, logger log.Logger) (s *Service, cf func(), err error) {
	s = &Service{
		conf: c,
		dao:  d,
		log:  log.NewHelper(log.WithFields(logger, "pkg", "Service")),
	}
	return s, func() {
		s.Close()
	}, nil
}

func (s *Service) Close() {

}

func (s *Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	save, err := s.dao.ShopRepo.Save(ctx, &dao.Shop{
		Name: in.GetName(),
	})

	return &v1.HelloReply{Message: "Hello " + save.Name},
		errorx.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found").WithCause(err)
}
