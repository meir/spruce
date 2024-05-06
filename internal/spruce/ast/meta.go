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
			tokenizer := spruce.GetTokenizer(ctx)
			current, err := tokenizer.Current()

			if err != nil {
				panic(err)
			}

			if current.MatchString("meta") {
				ctx = spruce.SetState(ctx, spruce.STATE_META)
				ctx = spruce.SetAST(ctx, &MetaAST{})
				return ctx, true
			}
			return ctx, false
		},
	})
}

type MetaAST struct{}

func (m *MetaAST) Build(ctx context.Context) (bool, error) {
	children := spruce.GetChildren(ctx)
	return len(*children) > 0, nil
}

func (m *MetaAST) String(ctx context.Context) string {
	return ""
}
