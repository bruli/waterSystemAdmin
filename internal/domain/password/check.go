package password

import "context"

type Check struct {
	repo Repository
}

func (c Check) Check(ctx context.Context, pass string) (bool, error) {
	passwd, err := c.repo.Read(ctx)
	if err != nil {
		return false, err
	}
	err = passwd.Compare(pass)
	return err == nil, nil
}

func NewCheck(repo Repository) *Check {
	return &Check{repo: repo}
}
