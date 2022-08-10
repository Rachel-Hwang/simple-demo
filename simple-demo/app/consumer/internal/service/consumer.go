package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "helloworld/api/consumer/service/v1"
	"helloworld/app/consumer/internal/biz"
)

type ConsumerService struct {
	pb.UnimplementedConsumerServer
	consumerBiz  *biz.ConsumerUsecase
	log *log.Helper
}

func NewConsumerService(su *biz.ConsumerUsecase) *ConsumerService {
	return &ConsumerService{
		consumerBiz: su,
	}
}

func (s *ConsumerService) CreateConsumer(ctx context.Context, req *pb.CreateConsumerReq) (*pb.CreateConsumerReply, error) {
	s.consumerBiz.CreateConsumer(
		ctx,
		&biz.Consumer{ID: 1, Name: "Rachel"},
	)
	return &pb.CreateConsumerReply{}, nil
}
func (s *ConsumerService) UpdateConsumer(ctx context.Context, req *pb.UpdateConsumerReq) (*pb.UpdateConsumerReply, error) {
	return &pb.UpdateConsumerReply{}, nil
}
func (s *ConsumerService) DeleteConsumer(ctx context.Context, req *pb.DeleteConsumerReq) (*pb.DeleteConsumerReply, error) {
	return &pb.DeleteConsumerReply{}, nil
}
func (s *ConsumerService) GetConsumer(ctx context.Context, req *pb.GetConsumerReq) (*pb.GetConsumerReply, error) {
	return &pb.GetConsumerReply{}, nil
}
func (s *ConsumerService) ListConsumer(ctx context.Context, req *pb.ListConsumerReq) (*pb.ListConsumerReply, error) {
	return &pb.ListConsumerReply{}, nil
}

func (s *ConsumerService) CallProvider(ctx context.Context, req *pb.CallProviderReq) (*pb.CallProviderReply, error) {
	cu,err := s.consumerBiz.CallProvider(ctx,&biz.Consumer{
		ID: req.Id,
	})
	if err != nil {
		return nil, err
	}
	s.log.WithContext(ctx).Infof("CallProvider: %v", cu)
	return &pb.CallProviderReply{}, nil
}
