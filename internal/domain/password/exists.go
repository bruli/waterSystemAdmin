package password

import "context"

type Exists struct {
	repo Repository
}

func (e Exists) Exists(ctx context.Context) (bool, error) {
	return e.repo.Exists(ctx)
}

func NewExists(repo Repository) *Exists {
	return &Exists{repo: repo}
}
