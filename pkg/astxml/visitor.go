package astxml

import (
	"fmt"
	"go/ast"
	"strconv"
)

type visitor struct {
	parent Node
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	var xn Node
	if n != nil {
		xn = mapNode(n)
		v.parent.AddChild(xn)
	} else {
		xn = v.parent
	}
	return &visitor{parent: xn}
}

// mapNode converts AST elements into XML serializable types
func mapNode(node ast.Node) (xn Node) {
	switch n := node.(type) {
	case *ast.Comment:
		xn = &Comment{Text: n.Text}
	case *ast.CommentGroup:
		xn = &CommentGroup{}
	case *ast.Field:
		xn = &Field{}
	case *ast.FieldList:
		xn = &FieldList{}
	case *ast.BadExpr:
		xn = &BadExpr{}
	case *ast.Ident:
		id := &Ident{Name: n.Name}
		if n.Obj != nil {
			id.Object = &Object{
				Name: n.Obj.Name,
				Kind: n.Obj.Kind.String(),
				Data: n.Obj.Data,
			}
		}
		xn = id
	case *ast.BasicLit:
		var val string
		if str, err := strconv.Unquote(n.Value); err == nil {
			val = str
		}
		xn = &BasicLit{
			Kind:  n.Kind.String(),
			Value: val,
		}
	case *ast.Ellipsis:
		xn = &Ellipsis{}
	case *ast.FuncLit:
		xn = &FuncLit{}
	case *ast.CompositeLit:
		xn = &CompositeLit{}
	case *ast.ParenExpr:
		xn = &ParenExpr{}
	case *ast.SelectorExpr:
		xn = &SelectorExpr{}
	case *ast.IndexExpr:
		xn = &IndexExpr{}
	case *ast.SliceExpr:
		xn = &SliceExpr{}
	case *ast.TypeAssertExpr:
		xn = &TypeAssertExpr{}
	case *ast.CallExpr:
		xn = &CallExpr{}
	case *ast.StarExpr:
		xn = &StarExpr{}
	case *ast.UnaryExpr:
		xn = &UnaryExpr{
			Op: n.Op.String(),
		}
	case *ast.BinaryExpr:
		xn = &BinaryExpr{
			Op: n.Op.String(),
		}
	case *ast.KeyValueExpr:
		xn = &KeyValueExpr{}
	case *ast.ArrayType:
		xn = &ArrayType{}
	case *ast.StructType:
		xn = &StructType{}
	case *ast.FuncType:
		xn = &FuncType{}
	case *ast.InterfaceType:
		xn = &InterfaceType{}
	case *ast.MapType:
		xn = &MapType{}
	case *ast.ChanType:
		xn = &ChanType{
			Dir: mapChanDir(n.Dir),
		}
	case *ast.BadStmt:
		xn = &BadStmt{}
	case *ast.DeclStmt:
		xn = &DeclStmt{}
	case *ast.EmptyStmt:
		xn = &EmptyStmt{}
	case *ast.LabeledStmt:
		xn = &LabeledStmt{}
	case *ast.ExprStmt:
		xn = &ExprStmt{}
	case *ast.SendStmt:
		xn = &SendStmt{}
	case *ast.IncDecStmt:
		xn = &IncDecStmt{}
	case *ast.AssignStmt:
		xn = &AssignStmt{}
	case *ast.GoStmt:
		xn = &GoStmt{}
	case *ast.DeferStmt:
		xn = &DeferStmt{}
	case *ast.ReturnStmt:
		xn = &ReturnStmt{}
	case *ast.BranchStmt:
		xn = &BranchStmt{}
	case *ast.BlockStmt:
		xn = &BlockStmt{}
	case *ast.IfStmt:
		xn = &IfStmt{}
	case *ast.CaseClause:
		xn = &CaseClause{}
	case *ast.SwitchStmt:
		xn = &SwitchStmt{}
	case *ast.TypeSwitchStmt:
		xn = &TypeSwitchStmt{}
	case *ast.CommClause:
		xn = &CommClause{}
	case *ast.SelectStmt:
		xn = &SelectStmt{}
	case *ast.ForStmt:
		xn = &ForStmt{}
	case *ast.RangeStmt:
		xn = &RangeStmt{}
	case *ast.ImportSpec:
		xn = &ImportSpec{}
	case *ast.ValueSpec:
		xn = &ValueSpec{}
	case *ast.TypeSpec:
		xn = &TypeSpec{}
	case *ast.BadDecl:
		xn = &BadDecl{}
	case *ast.GenDecl:
		xn = &GenDecl{}
	case *ast.FuncDecl:
		xn = &FuncDecl{}
	case *ast.File:
		xn = &File{}
	default:
		panic(fmt.Sprintf("unexpected node type %T", n))
	}

	xn.SetPosition(node.Pos(), node.End())

	return xn
}

func mapChanDir(dir ast.ChanDir) string {
	switch dir {
	case ast.SEND:
		return "send"
	case ast.RECV:
		return "recv"
	default:
		return ""
	}
}
