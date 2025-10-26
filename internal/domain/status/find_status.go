package status

import "context"

type FindStatus struct {
	repo Repository
}

func (f FindStatus) Find(ctx context.Context) (*Status, error) {
	return f.repo.Find(ctx)
}

func NewFindStatus(repo Repository) *FindStatus {
	return &FindStatus{repo: repo}
}
