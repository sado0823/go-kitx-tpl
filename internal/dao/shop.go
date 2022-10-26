package dao

import (
	"context"

	"github.com/sado0823/go-kitx/kit/log"
	"github.com/sado0823/go-kitx/kit/tracing"
)

type (
	ShopRepo interface {
		ListAll(context.Context) ([]*Shop, error)
	}

	Shop struct {
		ID      int64  `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	shopRepo struct {
		data *dao
		log  *log.Helper
	}
)

func newShopRepo(data *dao, logger log.Logger) ShopRepo {
	return &shopRepo{data, log.NewHelper(log.WithFields(logger, "pkg", "ShopRepo"))}
}

func (r *shopRepo) ListAll(ctx context.Context) ([]*Shop, error) {
	r.log.WithContext(ctx).Warnf("shop list all")
	ctx, span := tracing.Get("go-kitx").Start(ctx, "shopRepo.ListAll")
	span.End()
	return []*Shop{
		{ID: 1, Name: "shop1", Address: "address1"},
		{ID: 2, Name: "shop2", Address: "address2"},
	}, nil
}
