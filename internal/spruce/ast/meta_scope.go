package ast

import (
	"context"

	"github.com/meir/spruce/internal/spruce"
)

func init() {
	spruce.RegisterNode(spruce.Node{
		States: []spruce.State{
			spruce.STATE_AT_TAG,
		},

		Activate: func(ctx context.Context) (context.Context, bool) {

			return ctx, true
		},
	})
}

type MetaScopeAST struct{}

func (m *MetaScopeAST) Build(ctx context.Context) (bool, error) {

	return false, nil
}

func (m *MetaScopeAST) String(ctx context.Context) string {
	return ""
}
