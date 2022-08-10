package service

import (
	"context"
	"helloworld/app/provider/internal/biz"

	pb "helloworld/api/provider/service/v1"
)

type ProviderService struct {
	pb.UnimplementedProviderServer
	pu *biz.ProviderUsecase
}

func NewProviderService(pu *biz.ProviderUsecase) *ProviderService {
	return &ProviderService{
		pu:pu,
	}
}

func (s *ProviderService) CreateProvider(ctx context.Context, req *pb.CreateProviderRequest) (*pb.CreateProviderReply, error) {
	return &pb.CreateProviderReply{}, nil
}
func (s *ProviderService) UpdateProvider(ctx context.Context, req *pb.UpdateProviderRequest) (*pb.UpdateProviderReply, error) {
	return &pb.UpdateProviderReply{}, nil
}
func (s *ProviderService) DeleteProvider(ctx context.Context, req *pb.DeleteProviderRequest) (*pb.DeleteProviderReply, error) {
	return &pb.DeleteProviderReply{}, nil
}
func (s *ProviderService) GetProvider(ctx context.Context, req *pb.GetProviderRequest) (*pb.GetProviderReply, error) {
	s.pu.GetProvider(ctx,&biz.Provider{
		ID: req.Id,
	})
	return &pb.GetProviderReply{}, nil
}
func (s *ProviderService) ListProvider(ctx context.Context, req *pb.ListProviderRequest) (*pb.ListProviderReply, error) {
	return &pb.ListProviderReply{}, nil
}
