package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go/aws/session"
	pb "github.com/mbmanthey/foodtruck/proto"
)

type Service struct {
	Session *session.Session
}

func (s *Service) GetRepo() Repository {
	return &TruckRepository{s.Session}
}

func (s *Service) Create(ctx context.Context, req *pb.Truck) (*pb.Response, error) {
	err := s.GetRepo().Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Truck: req}, nil
}

func (s *Service) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	trucks, err := s.GetRepo().GetAll()
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: false, Trucks: trucks}, nil
}

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	truck, err := s.GetRepo().Get(req.ID)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: false, Truck: truck}, nil
}

func (s *Service) Delete(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	err := s.GetRepo().Delete(req.ID)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: false, Trucks: nil}, nil
}
