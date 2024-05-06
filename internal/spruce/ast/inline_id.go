package ast

import (
	"context"

	"github.com/meir/spruce/internal/spruce"
	"github.com/meir/spruce/internal/spruce/types"
)

func init() {
	spruce.RegisterNode(spruce.Node{
		States: []spruce.State{
			spruce.STATE_IDENTIFIER,
		},

		Activate: func(ctx context.Context) (context.Context, bool) {
			tokenizer := spruce.GetTokenizer(ctx)
			current, err := tokenizer.Current()
			if err != nil {
				return ctx, false
			}

			if !current.Match("#") {
				return ctx, false
			}

			key, err := tokenizer.Peek(1)
			if err != nil {
				return ctx, false
			}

			if !key.Match(`[a-zA-Z_][a-zA-Z0-9_-]*`) {
				return ctx, false
			}

			ctx = spruce.SetAST(ctx, &InlineIdAST{
				Key: key.String(),
			})
			return ctx, true
		},
	})
}

type InlineIdAST struct {
	Key string
}

func (i *InlineIdAST) Build(ctx context.Context) (bool, error) {
	return true, nil
}

func (i *InlineIdAST) String(ctx context.Context) string {
	scope := spruce.GetScope(ctx)
	if id := scope.Get("id"); id != nil {
		kv := id.Value().([]any)
		kv[1] = i.Key
		id.Set(kv)
		return ""
	}

	scope.Define("id", types.NewAttr("id", i.Key))
	return ""
}
