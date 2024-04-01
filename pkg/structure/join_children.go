package structure

import "strings"

func JoinChildren(asts []*ASTWrapper) string {
	str := []string{}
	for _, child := range asts {
		str = append(str, child.Ast.String(child.Children))
	}
	return strings.Join(str, "")
}
