package programs

import "context"

type FindAllPrograms struct {
	repo Repository
}

func (d FindAllPrograms) Find(ctx context.Context) (*Programs, error) {
	return d.repo.FindAll(ctx)
}

func NewFindAllPrograms(repo Repository) *FindAllPrograms {
	return &FindAllPrograms{repo: repo}
}
