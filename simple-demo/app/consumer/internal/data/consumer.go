package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	provider "helloworld/api/provider/service/v1"
	"helloworld/app/consumer/internal/biz"
)

type consumerRepo struct {
	data *Data
	providerSrv provider.ProviderClient
	log  *log.Helper
}


func (r *consumerRepo) Save(ctx context.Context, g *biz.Consumer) (*biz.Consumer, error) {
	return g, nil
}

func (r *consumerRepo) Update(ctx context.Context, g *biz.Consumer) (*biz.Consumer, error) {
	return g, nil
}

func (r *consumerRepo) FindByID(context.Context, int64) (*biz.Consumer, error) {
	return nil, nil
}

func (r *consumerRepo) ListByHello(context.Context, string) ([]*biz.Consumer, error) {
	return nil, nil
}

func (r *consumerRepo) ListAll(context.Context) ([]*biz.Consumer, error) {
	return nil, nil
}
