package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"helloworld/app/provider/internal/biz"
)

type providerRepo struct {
	data *Data
	log  *log.Helper
}

// NewProviderRepo .
func NewProviderRepo(data *Data, logger log.Logger) biz.ProviderRepo {
	return &providerRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *providerRepo) Save(ctx context.Context, g *biz.Provider) (*biz.Provider, error) {
	return g, nil
}

func (r *providerRepo) Update(ctx context.Context, g *biz.Provider) (*biz.Provider, error) {
	return g, nil
}

func (r *providerRepo) FindByID(context.Context, int64) (*biz.Provider, error) {
	return nil, nil
}

func (r *providerRepo) ListByHello(context.Context, string) ([]*biz.Provider, error) {
	return nil, nil
}

func (r *providerRepo) ListAll(context.Context) ([]*biz.Provider, error) {
	return nil, nil
}
