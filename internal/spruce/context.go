package spruce

import (
	"context"
	"strings"
)

const ContextKey_Scope = "Scope"
const ContextKey_File = "File"
const ContextKey_Tokenizer = "Tokenizer"
const ContextKey_Lexer = "Lexer"
const ContextKey_State = "State"
const ContextKey_Children = "Children"
const ContextKey_AST = "AST"

func SetScope(ctx context.Context, scope *Scope) context.Context {
	return context.WithValue(ctx, ContextKey_Scope, scope)
}

func GetScope(ctx context.Context) *Scope {
	return ctx.Value(ContextKey_Scope).(*Scope)
}

func SetFile(ctx context.Context, file *File) context.Context {
	return context.WithValue(ctx, ContextKey_File, file)
}

func GetFile(ctx context.Context) *File {
	return ctx.Value(ContextKey_File).(*File)
}

func SetTokenizer(ctx context.Context, tokenizer *Tokenizer) context.Context {
	return context.WithValue(ctx, ContextKey_Tokenizer, tokenizer)
}

func GetTokenizer(ctx context.Context) *Tokenizer {
	return ctx.Value(ContextKey_Tokenizer).(*Tokenizer)
}

func SetLexer(ctx context.Context, lexer *Lexer) context.Context {
	return context.WithValue(ctx, ContextKey_Lexer, lexer)
}

func GetLexer(ctx context.Context) *Lexer {
	return ctx.Value(ContextKey_Lexer).(*Lexer)
}

func SetState(ctx context.Context, state State) context.Context {
	return context.WithValue(ctx, ContextKey_State, state)
}

func GetState(ctx context.Context) State {
	return ctx.Value(ContextKey_State).(State)
}

func SetChildren(ctx context.Context, children *[]context.Context) context.Context {
	return context.WithValue(ctx, ContextKey_Children, children)
}

func GetChildren(ctx context.Context) *[]context.Context {
	return ctx.Value(ContextKey_Children).(*[]context.Context)
}

func SetAST(ctx context.Context, ast AST) context.Context {
	return context.WithValue(ctx, ContextKey_AST, ast)
}

func GetAST(ctx context.Context) AST {
	if v, ok := ctx.Value(ContextKey_AST).(AST); ok {
		return v
	}
	return nil
}

func BuildChildren(ctx context.Context) string {
	builder := strings.Builder{}
	for _, child := range *GetChildren(ctx) {
		builder.WriteString(GetAST(child).String(child))
	}
	return builder.String()
}
