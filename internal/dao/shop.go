package dao

import (
	"context"

	"github.com/sado0823/go-kitx/kit/log"
)

type (
	ShopRepo interface {
		Save(ctx context.Context, shop *Shop) (*Shop, error)
		ListAll(context.Context) ([]*Shop, error)
	}

	Shop struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	shopRepo struct {
		data *dao
		log  *log.Helper
	}
)

func newShopRepo(data *dao, logger log.Logger) ShopRepo {
	return &shopRepo{data, log.NewHelper(log.WithFields(logger, "pkg", "shopRepo"))}
}

func (r *shopRepo) Save(ctx context.Context, g *Shop) (*Shop, error) {
	r.log.Errorf("i am greeterRepo!!!")
	return g, nil
}

func (r *shopRepo) ListAll(context.Context) ([]*Shop, error) {
	return nil, nil
}
