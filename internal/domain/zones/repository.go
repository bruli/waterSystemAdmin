package zones

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Zone, error)
	Create(ctx context.Context, z *Zone) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, z *Zone) error
}
