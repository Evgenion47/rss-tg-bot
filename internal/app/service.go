package app

import (
	pb "awesomeProjectIII/pkg/api"
)

type server struct {
	pb.UnimplementedAwesomeBotIIIServer
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
