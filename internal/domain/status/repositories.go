package status

import "context"

type Repository interface {
	Find(ctx context.Context) (*Status, error)
	Update(ctx context.Context) error
}

type ActivateRepository interface {
	Activate(ctx context.Context, activate bool) error
}
