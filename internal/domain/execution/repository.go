package execution

import "context"

type Repository interface {
	SendExecution(ctx context.Context, exe *Execution) error
}
