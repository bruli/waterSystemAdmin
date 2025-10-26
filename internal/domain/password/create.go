package password

import "context"

type Create struct {
	repo Repository
}

func (c Create) Create(ctx context.Context, pass string) error {
	passw, err := NewPassword(pass)
	if err != nil {
		return err
	}
	return c.repo.Save(ctx, passw)
}

func NewCreate(repo Repository) *Create {
	return &Create{repo: repo}
}
