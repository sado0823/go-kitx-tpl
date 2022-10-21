package dao

import (
	"context"
	"errors"

	"github.com/sado0823/go-kitx/kit/log"
)

type (
	BookRepo interface {
		Save(ctx context.Context, book *Book) error
		FindByID(ctx context.Context, id int64) (*Book, error)
	}

	Book struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	bookRepo struct {
		data *dao
		log  *log.Helper
	}
)

func newBookRepo(data *dao, logger log.Logger) BookRepo {
	return &bookRepo{data, log.NewHelper(log.WithFields(logger, "pkg", "bookRepo"))}
}

func (b *bookRepo) Save(ctx context.Context, book *Book) error {
	return errors.New("book save error")
}

func (b *bookRepo) FindByID(ctx context.Context, id int64) (*Book, error) {
	return &Book{ID: id, Name: "book repo"}, nil
}
