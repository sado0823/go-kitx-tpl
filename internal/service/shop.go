package service

import (
	"context"

	v1 "github.com/sado0823/go-kitx/tpl/api/helloworld/v1"
)

func (s *Service) ShopList(ctx context.Context, req *v1.ShopListRequest) (*v1.ShopListReply, error) {
	s.log.WithContext(ctx).Warnf("ShopList")
	shops, err := s.dao.ShopRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	resp := &v1.ShopListReply{Shops: make([]*v1.Shop, 0, len(shops))}
	for _, shop := range shops {
		resp.Shops = append(resp.Shops, &v1.Shop{Name: shop.Name, Address: shop.Address})
	}
	return resp, nil
}
