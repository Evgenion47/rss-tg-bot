package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework-2/internal/models"
	pb "homework-2/pkg/api"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.ChData) (*empty.Empty, error) {
	var user = models.User{
		ChatID: req.ChatID,
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		err = AlreadyExistErr
	}

	return &emptypb.Empty{}, err
}
