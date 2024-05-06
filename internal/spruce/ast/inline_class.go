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

			if !current.Match("\\.") {
				return ctx, false
			}

			key, err := tokenizer.Peek(1)
			if err != nil {
				return ctx, false
			}

			if !key.Match(`[a-zA-Z_][a-zA-Z0-9_-]*`) {
				return ctx, false
			}

			ctx = spruce.SetAST(ctx, &InlineClassAST{
				Key: key.String(),
			})
			return ctx, true
		},
	})
}

type InlineClassAST struct {
	Key string
}

func (i *InlineClassAST) Build(ctx context.Context) (bool, error) {
	return true, nil
}

func (i *InlineClassAST) String(ctx context.Context) string {
	scope := spruce.GetScope(ctx)
	if class := scope.Get("class"); class != nil {
		kv := class.Value().([]any)
		value := kv[1]

		switch value.(type) {
		case string:
			kv[1] = value.(string) + " " + i.Key
		default:
			panic("expected value type 'string' got value %v")
		}

		class.Set(kv)
		return ""
	}

	scope.Define("class", types.NewAttr("class", i.Key))
	return ""
}
