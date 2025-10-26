package logs

import "context"

type FindLogs struct {
	repo Repository
}

func (l FindLogs) Find(ctx context.Context) ([]Log, error) {
	return l.repo.Find(ctx)
}

func NewFindLogs(repo Repository) *FindLogs {
	return &FindLogs{repo: repo}
}
