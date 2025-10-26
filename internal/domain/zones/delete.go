package zones

import "context"

type Delete struct {
	repo Repository
}

func (d Delete) Delete(ctx context.Context, id string) error {
	return d.repo.Delete(ctx, id)
}

func NewDelete(repo Repository) *Delete {
	return &Delete{repo: repo}
}
