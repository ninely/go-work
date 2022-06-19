package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/ninely/go-work/final/api/homework/v1"
)

var (
	ErrNoResult = errors.NotFound(v1.ErrorReason_NO_RESULT.String(), "result not found")
)

type Homework struct {
	Hello string
}

type HomeworkRepo interface {
	Save(context.Context, *Homework) (*Homework, error)
	Update(context.Context, *Homework) (*Homework, error)
	FindByID(context.Context, int64) (*Homework, error)
	ListByHello(context.Context, string) ([]*Homework, error)
	ListAll(context.Context) ([]*Homework, error)
}

type HomeworkUseCase struct {
	repo HomeworkRepo
	log  *log.Helper
}

func NewHomeworkUseCase(repo HomeworkRepo, logger log.Logger) *HomeworkUseCase {
	return &HomeworkUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *HomeworkUseCase) GetRecords(ctx context.Context, g *Homework) (*Homework, error) {
	uc.log.WithContext(ctx).Infof("GetRecords: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *HomeworkUseCase) FindDetailSolution(ctx context.Context, g *Homework) (*Homework, error) {
	uc.log.WithContext(ctx).Infof("FindDetailSolution: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
