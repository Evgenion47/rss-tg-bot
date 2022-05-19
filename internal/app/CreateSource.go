package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework-2/internal/models"
	pb "homework-2/pkg/api"
)

func (s *Server) CreateSource(ctx context.Context, req *pb.ChSrcData) (*empty.Empty, error) {
	var data = models.DCData{
		ChatID: req.ChatID,
		Source: req.Source,
	}

	err := s.repo.CreateSource(ctx, data)
	if err != nil && err != UserDoesNotExist {
		err = AlreadyExistErr
	}

	return &emptypb.Empty{}, err
}
