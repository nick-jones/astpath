package query

import "github.com/antchfx/xpath"

type Expr = xpath.Expr

func Compile(expr string) (*Expr, error) {
	return xpath.Compile(expr)
}
