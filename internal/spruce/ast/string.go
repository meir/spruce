package ast

import (
	"context"

	"github.com/meir/spruce/internal/spruce"
)

func init() {
	spruce.RegisterNode(spruce.Node{
		States: []spruce.State{
			spruce.STATE_SCOPE,
			spruce.STATE_ROOT,
			spruce.STATE_INLINE_ATTRIBUTE,
		},
		Activate: func(ctx context.Context) (context.Context, bool) {
			tokenizer := spruce.GetTokenizer(ctx)
			t, err := tokenizer.Current()

			if err != nil {
				panic(err)
			}

			quote := ""
		QuoteSwitch:
			switch t.String() {
			case "`":
				for i := 0; i < 3; i++ {
					peek, err := tokenizer.Peek(i)
					if err != nil || peek.String() != "`" {
						quote = "`"
						break QuoteSwitch
					}
				}

				quote = "```"
				tokenizer.Skip(2)
				break
			case "'", "\"":
				quote = t.String()
				break
			}

			if quote == "" {
				return ctx, false
			}

			ctx = spruce.SetState(ctx, spruce.STATE_STRING)
			ctx = spruce.SetAST(ctx, &StringAST{
				Quote: quote,
			})
			return ctx, true
		},
	})
}

type StringAST struct {
	Quote   string
	content string
}

func (ast *StringAST) Build(ctx context.Context) (bool, error) {
	tokenizer := spruce.GetTokenizer(ctx)
	current, err := tokenizer.Current()

	if err != nil {
		return false, err
	}

QuoteSwitch:
	switch current.String() {
	case "\\":
		if err = tokenizer.Next(); err != nil {
			return false, err
		}
		current, err = tokenizer.Current()
		if err != nil {
			return false, err
		}
		break
	case "\n":
		if ast.Quote != "```" {
			return false, spruce.NewUnexpectedToken(current)
		}
	case "`":
		if ast.Quote == "```" {
			for i := 0; i < 3; i++ {
				peek, err := tokenizer.Peek(i)
				if err != nil || peek.String() != "`" {
					break QuoteSwitch
				}
			}
		}

		fallthrough
	case ast.Quote:
		return true, nil
	}

	ast.content += current.String()

	return false, nil
}

func (ast *StringAST) String(ctx context.Context) string {
	return ast.content
}
