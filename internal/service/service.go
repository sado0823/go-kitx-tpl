package service

import (
	"context"
	"errors"

	"github.com/sado0823/go-kitx/errorx"
	"github.com/sado0823/go-kitx/kit/log"
	"github.com/sado0823/go-kitx/kit/tracing"
	v1 "github.com/sado0823/go-kitx/tpl/api/helloworld/v1"
	"github.com/sado0823/go-kitx/tpl/internal/conf"
	"github.com/sado0823/go-kitx/tpl/internal/dao"

	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/emptypb"
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
	s.log.WithContext(ctx).Warnf("hello service log")
	ctx, span := tracing.Get("go-kitx").Start(ctx, "service.SayHello")
	span.End()

	_, _ = s.dao.ShopRepo.ListAll(ctx)
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *Service) GotError(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	s.log.WithContext(ctx).Error("got error")
	return nil, errorx.Unauthorized(v1.ErrorReason_USER_NOT_FOUND.String(), "用户不存在").
		WithCause(errors.New("customer error")).WithMetadata(map[string]string{"user_id": "0"})
}
