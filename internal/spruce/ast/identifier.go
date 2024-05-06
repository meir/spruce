package ast

import (
	"context"
	"fmt"
	"strings"

	"github.com/meir/spruce/internal/spruce"
	"github.com/meir/spruce/internal/spruce/types"
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

			if t.Match(`[a-zA-Z_][a-zA-Z0-9_-]*`) {
				ctx = spruce.SetState(ctx, spruce.STATE_IDENTIFIER)
				ctx = spruce.SetAST(ctx, &IdentifierAST{
					Str: t.String(),
				})
				return ctx, true
			}

			return ctx, false
		},
	})
}

type IdentifierAST struct {
	Str string
}

func (i *IdentifierAST) Build(ctx context.Context) (bool, error) {
	children := spruce.GetChildren(ctx)
	if len(*children) == 0 {
		return false, nil
	}

	lastChild := (*children)[len(*children)-1]

	switch spruce.GetAST(lastChild).(type) {
	case *ScopeAST:
		return true, nil
	}

	return false, nil
}

func (i *IdentifierAST) String(ctx context.Context) string {
	children := spruce.GetChildren(ctx)
	if len(*children) == 0 {
		panic("identifier expects atleast one valid child")
	}

	child := (*children)[len(*children)-1]

	switch spruce.GetAST(child).(type) {
	case *ScopeAST:
		scope := spruce.GetScope(ctx)
		v := scope.Get(i.Str)
		if v == nil {
			// identifier is not a variable, so must be an html tag
			content := spruce.BuildChildren(ctx)

			// get all attributes
			attrs := scope.GetByType(&types.Attr{})
			attributes := []string{}
			for _, a := range attrs {
				attributes = append(attributes, a.String())
			}

			if len(attributes) > 0 {
				return fmt.Sprintf("<%s %s>%s</%s>", i.Str, strings.Join(attributes, " "), content, i.Str)
			}

			return fmt.Sprintf("<%s>%s</%s>", i.Str, content, i.Str)
		} else {
			// identifier is a template

			return v.String()
		}
	}

	return i.Str
}
