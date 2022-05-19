package app

import (
	"context"
	"homework-2/internal/models"
	pb "homework-2/pkg/api"
)

func (s *Server) GetSrcsByChat(ctx context.Context, req *pb.ChData) (*pb.SrcSlcByChat, error) {
	var user = models.User{
		ChatID: req.ChatID,
	}

	data, err := s.repo.GetSrcsByChat(ctx, user)
	if err != nil {
		err = NotFoundErr
	}

	return &pb.SrcSlcByChat{Sources: data.Sources}, err
}
