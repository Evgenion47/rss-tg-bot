package app

import (
	"context"
	"homework-2/internal/models"
	pb "homework-2/pkg/api"
)

func (s *server) GetRSSBySource(ctx context.Context, req *pb.ChSrcData) (*pb.RSSSlc, error) {
	var data = models.DCData{
		ChatID: req.ChatID,
		Source: req.Source,
	}
	mess, err := s.repo.GetRSSBySource(ctx, data)
	if err != nil {
		err = NotFoundErr
	}
	return &pb.RSSSlc{News: mess.News}, err
}
