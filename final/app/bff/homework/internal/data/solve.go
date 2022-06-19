package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ninely/go-work/final/app/bff/homework/internal/biz"
)

type solveRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.SolveRepo = (*solveRepo)(nil)

func NewSolveRepo(data *Data, logger log.Logger) *solveRepo {
	return &solveRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *solveRepo) Solve(ctx context.Context, question *biz.Question) (*biz.SolveSolution, error) {
	return nil, nil
}
