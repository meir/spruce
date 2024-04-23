package spruce

import "context"

type AST interface {
	Build(ctx context.Context) (bool, error)
	String(ctx context.Context) string
}
