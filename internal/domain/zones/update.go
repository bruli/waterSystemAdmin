package zones

import "context"

type Update struct {
	repo Repository
}

func (c Update) Update(ctx context.Context, z *Zone) error {
	return c.repo.Update(ctx, z)
}

func NewUpdate(repo Repository) *Update {
	return &Update{repo: repo}
}
