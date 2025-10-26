package programs

import "context"

type CreateWeekly struct {
	repo WeeklyRepository
}

func (c CreateWeekly) Create(ctx context.Context, p *Weekly) error {
	return c.repo.Save(ctx, p)
}

func NewCreateWeekly(repo WeeklyRepository) *CreateWeekly {
	return &CreateWeekly{repo: repo}
}
