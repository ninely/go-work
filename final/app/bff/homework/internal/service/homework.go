package service

import (
	"context"

	v1 "github.com/ninely/go-work/final/api/homework/v1"
	"github.com/ninely/go-work/final/app/bff/homework/internal/biz"
)

type HomeworkService struct {
	v1.UnimplementedHomeworkServer

	uc *biz.HomeworkUseCase
}

func NewHomeworkService(uc *biz.HomeworkUseCase) *HomeworkService {
	return &HomeworkService{uc: uc}
}

func (s *HomeworkService) GetRecords(ctx context.Context, in *v1.GetRecordsRequest) (*v1.GetRecordsReply, error) {
	_, err := s.uc.GetRecords(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.GetRecordsReply{}, nil
}

func (s *HomeworkService) FindDetailSolution(ctx context.Context, in *v1.FindSolutionRequest) (*v1.FindSolutionReply, error) {
	_, err := s.uc.FindSolution(ctx, in.GetUserId(), &biz.Question{
		Content: in.Content,
		FileUrl: in.FileUrl,
	})
	if err != nil {
		return nil, err
	}
	return &v1.FindSolutionReply{}, nil
}
