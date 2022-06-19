package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ninely/go-work/final/app/bff/homework/internal/biz"
)

type searchRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.SearchRepo = (*searchRepo)(nil)

func NewSearchRepo(data *Data, logger log.Logger) *searchRepo {
	return &searchRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *searchRepo) Search(ctx context.Context, fileUrl string) (*biz.SearchSolution, error) {
	return nil, nil
}
