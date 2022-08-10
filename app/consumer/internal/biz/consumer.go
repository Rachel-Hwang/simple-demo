package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	provider "helloworld/api/provider/service/v1"
)

type Consumer struct {
	ID int64
	Name string
}

type ConsumerRepo interface {
	Save(context.Context, *Consumer) (*Consumer, error)
	Update(context.Context, *Consumer) (*Consumer, error)
	FindByID(context.Context, int64) (*Consumer, error)
	ListByHello(context.Context, string) ([]*Consumer, error)
	ListAll(context.Context) ([]*Consumer, error)
}

type ConsumerUsecase struct {
	repo ConsumerRepo
	log  *log.Helper
	providerSrv provider.ProviderClient
}

// NewConsumerUsecase new a Consumer usecase.
func NewConsumerUsecase(repo ConsumerRepo, logger log.Logger) *ConsumerUsecase {
	return &ConsumerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateConsumer creates a Consumer, and returns the new Consumer.
func (uc *ConsumerUsecase) CreateConsumer(ctx context.Context, g *Consumer) (*Consumer, error) {
	uc.log.WithContext(ctx).Infof("CreateConsumer: %v", g.Name)
	return uc.repo.Save(ctx, g)
}


// CallProvider remote client call provider instance.
func (uc *ConsumerUsecase) CallProvider(ctx context.Context, g *Consumer) (*Consumer, error) {
	uc.log.WithContext(ctx).Infof("CreateConsumer: %v", g.Name)
	return uc.repo.Save(ctx, g)
}


