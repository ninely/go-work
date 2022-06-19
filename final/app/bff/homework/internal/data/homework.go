package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ninely/go-work/final/app/bff/homework/internal/biz"
)

type homeworkRepo struct {
	data *Data
	log  *log.Helper
}

// NewHomeworkRepo .
func NewHomeworkRepo(data *Data, logger log.Logger) biz.HomeworkRepo {
	return &homeworkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *homeworkRepo) Save(ctx context.Context, g *biz.Homework) (*biz.Homework, error) {
	return g, nil
}

func (r *homeworkRepo) Update(ctx context.Context, g *biz.Homework) (*biz.Homework, error) {
	return g, nil
}

func (r *homeworkRepo) FindByID(context.Context, int64) (*biz.Homework, error) {
	return nil, nil
}

func (r *homeworkRepo) ListByHello(context.Context, string) ([]*biz.Homework, error) {
	return nil, nil
}

func (r *homeworkRepo) ListAll(context.Context) ([]*biz.Homework, error) {
	return nil, nil
}
