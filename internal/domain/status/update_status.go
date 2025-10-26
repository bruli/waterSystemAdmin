package status

import "context"

type UpdateStatus struct {
	repo Repository
}

func (u UpdateStatus) Update(ctx context.Context) error {
	return u.repo.Update(ctx)
}

func NewUpdateStatus(repo Repository) *UpdateStatus {
	return &UpdateStatus{repo: repo}
}
