package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	//ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Provider is a Provider model.
type Provider struct {
	ID int64
	Name string
}

// ProviderRepo is a Provider repo.
type ProviderRepo interface {
	Save(context.Context, *Provider) (*Provider, error)
	Update(context.Context, *Provider) (*Provider, error)
	FindByID(context.Context, int64) (*Provider, error)
	ListByHello(context.Context, string) ([]*Provider, error)
	ListAll(context.Context) ([]*Provider, error)
}

// ProviderUsecase is a Provider usecase.
type ProviderUsecase struct {
	repo ProviderRepo
	log  *log.Helper
}

// NewProviderUsecase new a Provider usecase.
func NewProviderUsecase(repo ProviderRepo, logger log.Logger) *ProviderUsecase {
	return &ProviderUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateProvider creates a Provider, and returns the new Provider.
func (uc *ProviderUsecase) CreateProvider(ctx context.Context, g *Provider) (*Provider, error) {
	uc.log.WithContext(ctx).Infof("CreateProvider: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

// GetProvider returns the new Provider.
func (uc *ProviderUsecase) GetProvider(ctx context.Context, g *Provider) (*Provider, error) {
	uc.log.WithContext(ctx).Infof("GetProvider: %v", g.Name)
	return uc.repo.FindByID(ctx,g.ID)
}

