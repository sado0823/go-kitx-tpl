package service

import (
	"context"
	"encoding/json"

	v1 "github.com/sado0823/go-kitx/tpl/api/helloworld/v1"
	"github.com/sado0823/go-kitx/tpl/internal/dao"
)

func (s *Service) AddBook(ctx context.Context, req *v1.AddBookRequest) (*v1.AddBookReply, error) {
	msg, err := s.dao.BookRepo.Save(ctx, &dao.Book{
		Name:  req.GetName(),
		Price: req.GetPrice(),
	})
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &v1.AddBookReply{
		Message: msg + string(marshal),
		Price:   req.GetPrice(),
	}, nil
}
