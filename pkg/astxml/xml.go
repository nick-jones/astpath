package astxml

import (
	"encoding/xml"
	"fmt"
	"go/ast"
)

func Marshal(n ast.Node) ([]byte, error) {
	node, err := convertNode(n)
	if err != nil {
		return nil, err
	}
	return xml.Marshal(node)
}

func MarshalIndent(n ast.Node, prefix, indent string) ([]byte, error) {
	node, err := convertNode(n)
	if err != nil {
		return nil, err
	}
	return xml.MarshalIndent(node, prefix, indent)
}

func convertNode(n ast.Node) (Node, error) {
	root := &AST{}
	ast.Walk(&visitor{parent: root}, n)

	if len(root.Children) != 1 {
		return nil, fmt.Errorf("failed to parse file, unexpected number of children: %d", len(root.Children))
	}

	return root.Children[0], nil
}
