package programs

import "context"

type Remove struct {
	repo Repository
}

func (r Remove) Remove(ctx context.Context, hour *Hour, t *TypeProgram) error {
	return r.repo.Remove(ctx, hour, t)
}

func NewRemove(repo Repository) *Remove {
	return &Remove{repo: repo}
}
