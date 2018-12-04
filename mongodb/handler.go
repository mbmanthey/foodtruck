package mongodb

import (
	"context"

	"gopkg.in/mgo.v2"

	pb "github.com/mbmanthey/foodtruck/proto"
)

type Service struct {
	Session *mgo.Session
}

func (s *Service) GetRepo() Repository {
	return &TruckRepository{s.Session.Clone()}
}

func (s *Service) Create(ctx context.Context, req *pb.Truck) (*pb.Response, error) {
	defer s.GetRepo().Close()
	err := s.GetRepo().Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Truck: req}, nil
}

func (s *Service) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	sClone := &TruckRepository{s.Session.Clone()}
	defer sClone.Close()
	trucks, err := sClone.GetAll()
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: false, Trucks: trucks}, nil
}

func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	sClone := &TruckRepository{s.Session.Clone()}
	defer sClone.Close()
	truck, err := sClone.Get(req.ID)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: false, Truck: truck}, nil
}

func (s *Service) DeleteAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	sClone := &TruckRepository{s.Session.Clone()}
	defer sClone.Close()
	err := sClone.DeleteAll()
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: false, Trucks: nil}, nil
}
