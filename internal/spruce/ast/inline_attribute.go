package ast

import (
	"context"
	"regexp"

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

			if !current.Match(`[a-zA-Z_][a-zA-Z0-9_-]*`) {
				return ctx, false
			}

			t, err := tokenizer.PeekCheck(1, regexp.MustCompile("[^ ]+"))
			if err != nil {
				return ctx, false
			}

			if t.Match("=") {
				ctx = spruce.SetState(ctx, spruce.STATE_INLINE_ATTRIBUTE)
				ctx = spruce.SetAST(ctx, &InlineAttributeAST{
					Key: current.String(),
				})
				return ctx, true
			}

			return ctx, false
		},
	})
}

type InlineAttributeAST struct {
	Key string
}

func (i *InlineAttributeAST) Build(ctx context.Context) (bool, error) {
	children := spruce.GetChildren(ctx)
	return len(*children) > 0, nil
}

func (i *InlineAttributeAST) String(ctx context.Context) string {
	scope := spruce.GetScope(ctx)
	scope.Define(i.Key, types.NewAttr(i.Key, spruce.BuildChildren(ctx)))
	return ""
}
