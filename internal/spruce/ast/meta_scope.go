package ast

import (
	"context"

	"github.com/meir/spruce/internal/spruce"
)

func init() {
	spruce.RegisterNode(spruce.Node{
		States: []spruce.State{
			spruce.STATE_META,
		},

		Activate: func(ctx context.Context) (context.Context, bool) {
			tokenizer := spruce.GetTokenizer(ctx)
			t, err := tokenizer.Current()

			if err != nil {
				panic(err)
			}

			if t.MatchString("{") {
				ctx = spruce.SetState(ctx, spruce.STATE_META_SCOPE)
				ctx = spruce.SetAST(ctx, &MetaScopeAST{})
				return ctx, true
			}
			return ctx, false
		},
	})
}

type MetaScopeAST struct{}

func (s *MetaScopeAST) Build(ctx context.Context) (bool, error) {
	tokenizer := spruce.GetTokenizer(ctx)
	t, err := tokenizer.Current()
	if err != nil {
		return false, err
	}

	return t.MatchString("}"), nil
}

func (s *MetaScopeAST) String(ctx context.Context) string {
	return spruce.BuildChildren(ctx)
}
