package password

import "context"

type Repository interface {
	Exists(ctx context.Context) (bool, error)
	Save(ctx context.Context, pass *Password) error
	Read(ctx context.Context) (*Password, error)
}
