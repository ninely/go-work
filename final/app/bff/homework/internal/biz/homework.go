package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/ninely/go-work/final/api/homework/v1"
	"golang.org/x/sync/errgroup"
)

var (
	ErrNoResult = errors.NotFound(v1.ErrorReason_NO_RESULT.String(), "result not found")
)

type Record struct {
	ID       uint64
	UID      uint64
	Question *Question
	Solution *Solution
}

type Question struct {
	Content string
	FileUrl string
}

type Solution struct {
	SearchSolution *SearchSolution
	SolveSolution  *SolveSolution
}

type SearchSolution struct {
	Solution string
}

type SolveSolution struct {
	Content   string
	Solution  string
	Knowledge int32
	Steps     []*SolutionStep
}

type SolutionStep struct {
	Name   string
	Output string
}

type RecordRepo interface {
	CreateRecord(context.Context, *Record) (*Record, error)
	UpdateRecord(context.Context, *Record) error
	ListRecord(context.Context, uint64) ([]*Record, error)
}

type SearchRepo interface {
	Search(context.Context, string) (*SearchSolution, error)
}

type SolveRepo interface {
	Solve(context.Context, *Question) (*SolveSolution, error)
}

type HomeworkUseCase struct {
	recordRepo RecordRepo
	searchRepo SearchRepo
	solveRepo  SolveRepo
	log        *log.Helper
}

func NewHomeworkUseCase(recordRepo RecordRepo, searchRepo SearchRepo, solveRepo SolveRepo, logger log.Logger) *HomeworkUseCase {
	return &HomeworkUseCase{
		recordRepo: recordRepo,
		searchRepo: searchRepo,
		solveRepo:  solveRepo,
		log:        log.NewHelper(logger),
	}
}

func (uc *HomeworkUseCase) GetRecords(ctx context.Context, uid uint64) ([]*Record, error) {
	result, err := uc.recordRepo.ListRecord(ctx, uid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *HomeworkUseCase) FindSolution(ctx context.Context, uid uint64, q *Question) (*Solution, error) {
	var (
		search   *SearchSolution
		solve    *SolveSolution
		solution *Solution
	)

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var err error
		if solve, err = uc.solveRepo.Solve(egCtx, q); err != nil {
			uc.log.WithContext(egCtx).Errorf("Solve failed: %v", q)
			return err
		}
		return nil
	})
	eg.Go(func() error {
		var err error
		if search, err = uc.searchRepo.Search(egCtx, q.FileUrl); err != nil {
			uc.log.WithContext(egCtx).Errorf("Search failed: %s", q.FileUrl)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, v1.ErrorUnspecified("find solution failed")
	}

	solution = &Solution{
		SearchSolution: search,
		SolveSolution:  solve,
	}

	if _, err := uc.recordRepo.CreateRecord(ctx, &Record{
		UID:      uid,
		Question: q,
		Solution: solution,
	}); err != nil {
		uc.log.WithContext(ctx).Errorf("create record failed: %v", err)
		return nil, v1.ErrorUnspecified("create record failed")
	}

	return solution, nil
}
