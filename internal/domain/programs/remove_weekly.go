package programs

import "context"

type RemoveWeekly struct {
	repo WeeklyRepository
}

func (r RemoveWeekly) Remove(ctx context.Context, day *WeekDay) error {
	return r.repo.Remove(ctx, day)
}

func NewRemoveWeekly(repo WeeklyRepository) *RemoveWeekly {
	return &RemoveWeekly{repo: repo}
}
