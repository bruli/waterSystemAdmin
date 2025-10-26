package zones

import "context"

type Create struct {
	repo Repository
}

func (c Create) Create(ctx context.Context, z *Zone) error {
	return c.repo.Create(ctx, z)
}

func NewCreate(repo Repository) *Create {
	return &Create{repo: repo}
}
