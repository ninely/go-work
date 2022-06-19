package data

import (
	"context"
	"encoding/json"
	v1 "github.com/ninely/go-work/final/api/record/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ninely/go-work/final/app/bff/homework/internal/biz"
)

type recordRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.RecordRepo = (*recordRepo)(nil)

func NewRecordRepo(data *Data, logger log.Logger) *recordRepo {
	return &recordRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *recordRepo) CreateRecord(ctx context.Context, record *biz.Record) (*biz.Record, error) {
	var (
		b   []byte
		err error
	)
	if record.Solution.SolveSolution != nil {

	}
	if b, err = json.Marshal(record.Solution.SolveSolution); err != nil {
		return nil, err
	}
	result, err := r.data.recordClient.CreateRecord(ctx, &v1.CreateRecordRequest{
		UserId: record.UID,
		Question: &v1.Question{
			Content: record.Question.Content,
			FileUrl: record.Question.FileUrl,
		},
		Solution: &v1.Solution{
			SearchSolution: record.Solution.SearchSolution.Solution,
			SolveSolution:  string(b),
		},
	})
	if err != nil {
		return nil, err
	}
	return &biz.Record{ID: result.Id}, nil
}

func (r *recordRepo) UpdateRecord(ctx context.Context, record *biz.Record) error {
	return nil
}

func (r *recordRepo) ListRecord(ctx context.Context, uid uint64) ([]*biz.Record, error) {
	return nil, nil
}
