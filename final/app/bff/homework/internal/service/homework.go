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
	_, err := s.uc.GetRecords(ctx, &biz.Homework{})
	if err != nil {
		return nil, err
	}
	return &v1.GetRecordsReply{}, nil
}

func (s *HomeworkService) FindDetailSolution(ctx context.Context, in *v1.FindDetailSolutionRequest) (*v1.FindDetailSolutionReply, error) {
	_, err := s.uc.FindDetailSolution(ctx, &biz.Homework{})
	if err != nil {
		return nil, err
	}
	return &v1.FindDetailSolutionReply{}, nil
}
