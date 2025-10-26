package logs

import "context"

type Repository interface {
	Find(ctx context.Context) ([]Log, error)
}
