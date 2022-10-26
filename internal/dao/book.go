package dao

import (
	"context"

	"github.com/sado0823/go-kitx/kit/log"
)

type (
	BookRepo interface {
		Save(ctx context.Context, book *Book) (string, error)
	}

	Book struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Price int64  `json:"price"`
	}

	bookRepo struct {
		data *dao
		log  *log.Helper
	}
)

func newBookRepo(data *dao, logger log.Logger) BookRepo {
	return &bookRepo{data, log.NewHelper(log.WithFields(logger, "pkg", "BookRepo"))}
}

func (b *bookRepo) Save(ctx context.Context, book *Book) (string, error) {
	b.log.WithContext(ctx).Warnf("save book with: %#v", book)
	return "save book ok", nil
}
