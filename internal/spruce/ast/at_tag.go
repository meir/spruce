package ast

import (
	"context"

	"github.com/meir/spruce/internal/spruce"
)

func init() {
	spruce.RegisterNode(spruce.Node{
		States: []spruce.State{
			spruce.STATE_ROOT,
		},
		Activate: func(ctx context.Context) (context.Context, bool) {
			tokenizer := spruce.GetTokenizer(ctx)
			current, err := tokenizer.Current()

			if err != nil {
				panic(err)
			}

			if current.MatchString("@") {
				ctx = spruce.SetState(ctx, spruce.STATE_AT_TAG)
				ctx = spruce.SetAST(ctx, &AtTagAST{})
				return ctx, true
			}
			return ctx, false
		},
	})
}

type AtTagAST struct{}

func (a *AtTagAST) Build(ctx context.Context) (bool, error) {
	return true, nil
}

func (a *AtTagAST) String(ctx context.Context) string {
	// TODO: Run children to fill scopes
	return ""
}
