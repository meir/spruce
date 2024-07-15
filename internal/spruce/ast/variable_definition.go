package ast

import (
	"context"

	"github.com/meir/spruce/internal/spruce"
	"github.com/meir/spruce/internal/spruce/types"
)

func init() {
	spruce.RegisterNode(spruce.Node{
		Priority: 1,
		States: []spruce.State{
			spruce.STATE_ROOT,
			spruce.STATE_SCOPE,
		},
		Activate: func(ctx context.Context) (context.Context, bool) {
			tokenizer := spruce.GetTokenizer(ctx)
			current, err := tokenizer.Current()

			if err != nil {
				panic(err)
			}

			var t spruce.Variable
			switch current.String() {
			case "attr":
				t = types.NewAttr("", "")
			case "bool":
				t = types.NewBool(false)
			case "int":
				t = types.NewInt(0)
				// case "slice":
				//   t = types.NewSlice()
			case "str":
				t = types.NewString("")
			}

			if t == nil {
				return ctx, false
			}

			ctx = spruce.SetState(ctx, spruce.STATE_VARIABLE_DEFINITION)
			ctx = spruce.SetAST(ctx, &VariableDefinitionAST{
				v: t,
			})
			return ctx, true
		},
	})
}

type VariableDefinitionAST struct {
	v spruce.Variable
}

func (a *VariableDefinitionAST) Build(ctx context.Context) (bool, error) {
	return true, nil
}

func (a *VariableDefinitionAST) String(ctx context.Context) string {
	// TODO: Run children to fill scopes
	return ""
}
