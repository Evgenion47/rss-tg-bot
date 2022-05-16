package app

import (
	pb "homework-2/pkg/api"
)

type server struct {
	repo Repository
	pb.UnimplementedAwesomeBotIIIServer
}

func New(repo Repository) *server {
	return &server{repo: repo}
}

//func (s server) CreateUser(ctx context.Context, req *pb.UData) (*empty.Empty, error) {
//
//}
//
//func (s server) CreateSource(ctx context.Context, req *pb.DelCreateData) (*empty.Empty, error) {
//
//}
//
//func (s server) GetSrcs(ctx context.Context, req *pb.UData) (*pb.SrcList, error) {
//
//}
//
//func (s server) DeleteSource(ctx context.Context, req *pb.DelCreateData) (*empty.Empty, error) {
//
//}
