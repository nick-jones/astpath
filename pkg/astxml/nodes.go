package astxml

import (
	"encoding/xml"
	"go/token"
)

type Node interface {
	AddChild(Node)
	SetPosition(start, end token.Pos)
}

type BaseNode struct {
	Children []Node
	PosStart token.Pos `xml:"pos-start,attr,omitempty"`
	PosEnd   token.Pos `xml:"pos-end,attr,omitempty"`
}

func (bn *BaseNode) AddChild(n Node) {
	bn.Children = append(bn.Children, n)
}

func (bn *BaseNode) SetPosition(start, end token.Pos) {
	bn.PosStart = start
	bn.PosEnd = end
}

type AST struct {
	XMLName xml.Name `xml:"AST"`
	BaseNode
}

type Comment struct {
	XMLName xml.Name `xml:"Comment"`
	Text    string
	BaseNode
}

type CommentGroup struct {
	XMLName xml.Name `xml:"CommentGroup"`
	BaseNode
}

type Field struct {
	XMLName xml.Name `xml:"Field"`
	BaseNode
}

type FieldList struct {
	XMLName xml.Name `xml:"FieldList"`
	BaseNode
}

type BadExpr struct {
	XMLName xml.Name `xml:"BadExpr"`
	BaseNode
}

type Ident struct {
	XMLName xml.Name `xml:"Ident"`
	Name    string   `xml:"name,attr"`
	Object  *Object
	BaseNode
}

type Object struct {
	XMLName xml.Name `xml:"Object"`
	Name    string   `xml:"name,attr"`
	Kind    string   `xml:"kind,attr"`
	Data    interface{}
}

type BasicLit struct {
	XMLName xml.Name `xml:"BasicLit"`
	Kind    string   `xml:"kind,attr"`
	Value   string   `xml:"value,attr"`
	BaseNode
}

type Ellipsis struct {
	XMLName xml.Name `xml:"Ellipsis"`
	BaseNode
}

type FuncLit struct {
	XMLName xml.Name `xml:"FuncLit"`
	BaseNode
}

type CompositeLit struct {
	XMLName xml.Name `xml:"CompositeLit"`
	BaseNode
}

type ParenExpr struct {
	XMLName xml.Name `xml:"ParenExpr"`
	BaseNode
}

type SelectorExpr struct {
	XMLName xml.Name `xml:"SelectorExpr"`
	BaseNode
}

type IndexExpr struct {
	XMLName xml.Name `xml:"IndexExpr"`
	BaseNode
}

type SliceExpr struct {
	XMLName xml.Name `xml:"SliceExpr"`
	BaseNode
}

type TypeAssertExpr struct {
	XMLName xml.Name `xml:"TypeAssertExpr"`
	BaseNode
}

type CallExpr struct {
	XMLName xml.Name `xml:"CallExpr"`
	BaseNode
}

type StarExpr struct {
	XMLName xml.Name `xml:"StarExpr"`
	BaseNode
}

type UnaryExpr struct {
	XMLName xml.Name `xml:"UnaryExpr"`
	Op      string   `xml:"op,attr"`
	BaseNode
}

type BinaryExpr struct {
	XMLName xml.Name `xml:"BinaryExpr"`
	Op      string   `xml:"op,attr"`
	BaseNode
}

type KeyValueExpr struct {
	XMLName xml.Name `xml:"KeyValueExpr"`
	BaseNode
}

type ArrayType struct {
	XMLName xml.Name `xml:"ArrayType"`
	BaseNode
}

type StructType struct {
	XMLName xml.Name `xml:"StructType"`
	BaseNode
}

type FuncType struct {
	XMLName xml.Name `xml:"FuncType"`
	BaseNode
}

type InterfaceType struct {
	XMLName xml.Name `xml:"InterfaceType"`
	BaseNode
}

type MapType struct {
	XMLName xml.Name `xml:"MapType"`
	BaseNode
}

type ChanType struct {
	XMLName xml.Name `xml:"ChanType"`
	Dir     string   `xml:"dir,attr"`
	BaseNode
}

type BadStmt struct {
	XMLName xml.Name `xml:"BadStmt"`
	BaseNode
}

type DeclStmt struct {
	XMLName xml.Name `xml:"DeclStmt"`
	BaseNode
}

type EmptyStmt struct {
	XMLName xml.Name `xml:"EmptyStmt"`
	BaseNode
}

type LabeledStmt struct {
	XMLName xml.Name `xml:"LabeledStmt"`
	BaseNode
}

type ExprStmt struct {
	XMLName xml.Name `xml:"ExprStmt"`
	BaseNode
}

type SendStmt struct {
	XMLName xml.Name `xml:"SendStmt"`
	BaseNode
}

type IncDecStmt struct {
	XMLName xml.Name `xml:"IncDecStmt"`
	BaseNode
}

type AssignStmt struct {
	XMLName xml.Name `xml:"AssignStmt"`
	BaseNode
}

type GoStmt struct {
	XMLName xml.Name `xml:"GoStmt"`
	BaseNode
}

type DeferStmt struct {
	XMLName xml.Name `xml:"DeferStmt"`
	BaseNode
}

type ReturnStmt struct {
	XMLName xml.Name `xml:"ReturnStmt"`
	BaseNode
}

type BranchStmt struct {
	XMLName xml.Name `xml:"BranchStmt"`
	BaseNode
}

type BlockStmt struct {
	XMLName xml.Name `xml:"BlockStmt"`
	BaseNode
}

type IfStmt struct {
	XMLName xml.Name `xml:"IfStmt"`
	BaseNode
}

type CaseClause struct {
	XMLName xml.Name `xml:"CaseClause"`
	BaseNode
}

type SwitchStmt struct {
	XMLName xml.Name `xml:"SwitchStmt"`
	BaseNode
}

type TypeSwitchStmt struct {
	XMLName xml.Name `xml:"TypeSwitchStmt"`
	BaseNode
}

type CommClause struct {
	XMLName xml.Name `xml:"CommClause"`
	BaseNode
}

type SelectStmt struct {
	XMLName xml.Name `xml:"SelectStmt"`
	BaseNode
}

type ForStmt struct {
	XMLName xml.Name `xml:"ForStmt"`
	BaseNode
}

type RangeStmt struct {
	XMLName xml.Name `xml:"RangeStmt"`
	BaseNode
}

type ImportSpec struct {
	XMLName xml.Name `xml:"ImportSpec"`
	BaseNode
}

type ValueSpec struct {
	XMLName xml.Name `xml:"ValueSpec"`
	BaseNode
}

type TypeSpec struct {
	XMLName xml.Name `xml:"TypeSpec"`
	BaseNode
}

type BadDecl struct {
	XMLName xml.Name `xml:"BadDecl"`
	BaseNode
}

type GenDecl struct {
	XMLName xml.Name `xml:"GenDecl"`
	BaseNode
}

type FuncDecl struct {
	XMLName xml.Name `xml:"FuncDecl"`
	BaseNode
}

type File struct {
	XMLName xml.Name `xml:"File"`
	BaseNode
}
