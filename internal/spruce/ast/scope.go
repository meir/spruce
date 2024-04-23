package ast

import (
	"context"

	"github.com/meir/spruce/internal/spruce"
)

func init() {
	spruce.RegisterNode(spruce.Node{
		States: []spruce.State{
			spruce.STATE_ROOT,
			spruce.STATE_SCOPE,
		},

		Activate: func(ctx context.Context) (context.Context, bool) {
			tokenizer := spruce.GetTokenizer(ctx)
			t, err := tokenizer.Current()

			if err != nil {
				panic(err)
			}

			if t.MatchString("{") {
				ctx = spruce.SetState(ctx, spruce.STATE_SCOPE)
				ctx = spruce.SetScope(ctx, spruce.NewScope(spruce.GetScope(ctx)))
				ctx = spruce.SetAST(ctx, &ScopeAST{})
				return ctx, true
			}
			return ctx, false
		},
	})
}

type ScopeAST struct{}

func (s *ScopeAST) Build(ctx context.Context) (bool, error) {
	tokenizer := spruce.GetTokenizer(ctx)
	t, err := tokenizer.Current()
	if err != nil {
		return false, err
	}

	return t.MatchString("}"), nil
}

func (s *ScopeAST) String(ctx context.Context) string {
	return spruce.BuildChildren(ctx)
}
