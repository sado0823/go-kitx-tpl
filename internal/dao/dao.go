package dao

import (
	"context"

	"github.com/sado0823/go-kitx/errorx"
	"github.com/sado0823/go-kitx/kit/log"
	"github.com/sado0823/go-kitx/kit/store/mysql"
	"github.com/sado0823/go-kitx/kit/store/redis"

	"github.com/sado0823/go-kitx/tpl/internal/conf"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDao)

type (
	Dao struct {
		BookRepo BookRepo
		ShopRepo ShopRepo
	}

	dao struct {
		conf  *conf.Data
		redis *redis.Redis
		mysql *mysql.Mysql
	}
)

func NewDao(c *conf.Data, logger log.Logger) (*Dao, func(), error) {

	data := &dao{conf: c}
	toNew := []func(c *conf.Data, data *dao) error{
		newRedis, newMysql,
	}

	for _, fn := range toNew {
		if err := fn(c, data); err != nil {
			return nil, nil, err
		}
	}

	return &Dao{
			BookRepo: newBookRepo(data, logger),
			ShopRepo: newShopRepo(data, logger),
		}, func() {
			err := data.Close()
			log.NewHelper(log.WithFields(logger, "pkg", "data")).Warnf("closing the data resources, err: %+v", err)
		}, nil
}

func (d *dao) Close() error {
	return new(errorx.Batch).Add(
		d.redis.Close(),
		d.mysql.Close(),
	).Err()
}

func newRedis(c *conf.Data, data *dao) error {
	data.redis = redis.New(c.Redis.Addr)
	return data.redis.Ping(context.Background())
}

func newMysql(c *conf.Data, data *dao) (err error) {
	data.mysql, err = mysql.New(c.Database.Source)
	return err
}
