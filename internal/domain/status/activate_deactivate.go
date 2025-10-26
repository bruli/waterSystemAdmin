package status

import "context"

type ActivateDeactivate struct {
	repo ActivateRepository
}

func (d ActivateDeactivate) Activate(ctx context.Context) error {
	return d.repo.Activate(ctx, true)
}

func (d ActivateDeactivate) DeActivate(ctx context.Context) error {
	return d.repo.Activate(ctx, false)
}

func NewActivateDeactivate(repo ActivateRepository) *ActivateDeactivate {
	return &ActivateDeactivate{repo: repo}
}
