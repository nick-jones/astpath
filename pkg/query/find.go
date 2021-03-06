package query

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"strconv"

	"github.com/antchfx/xmlquery"
	"github.com/nick-jones/astpath/internal/readutil"
	"github.com/nick-jones/astpath/pkg/astxml"
)

type Result struct {
	XML       string
	XMLInner  string
	Source    string
	LineValue string
	token.Position
}

// FindAll searches for all instances of the supplied expression
func FindAll(r io.ReaderAt, filename string, expr *Expr) ([]Result, error) {
	fset := token.NewFileSet()
	src, err := parser.ParseFile(fset, filename, r, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	xml, err := astToXML(src)
	if err != nil {
		return nil, err
	}

	doc, err := xmlquery.Parse(bytes.NewReader(xml))
	if err != nil {
		return nil, err
	}

	nodes := xmlquery.QuerySelectorAll(doc, expr)
	results := make([]Result, len(nodes))
	for i, node := range nodes {
		res, err := buildResult(node, r, fset)
		if err != nil {
			return nil, err
		}
		results[i] = res
	}
	return results, nil
}

func astToXML(f *ast.File) ([]byte, error) {
	xml, err := astxml.Marshal(f)
	if err != nil {
		return nil, err
	}

	return append([]byte(`<?xml version="1.0" encoding="UTF-8"?>`), xml...), nil
}

func positionsFromNode(node *xmlquery.Node) (start, end int, err error) {
	start, err = strconv.Atoi(node.SelectAttr("pos-start"))
	if err != nil {
		return 0, 0, fmt.Errorf("no source start position for node: %w", err)
	}
	end, err = strconv.Atoi(node.SelectAttr("pos-end"))
	if err != nil {
		return 0, 0, fmt.Errorf("no source end position for node: %w", err)
	}
	return
}

func buildResult(node *xmlquery.Node, r io.ReaderAt, fset *token.FileSet) (Result, error) {
	res := Result{
		XML:      node.OutputXML(true),
		XMLInner: node.OutputXML(false),
	}

	start, end, err := positionsFromNode(node)
	if err == nil {
		pos := fset.Position(token.Pos(start))
		res.Position = pos

		line, err := readutil.ReadLine(r, int64(start))
		if err != nil {
			return Result{}, err
		}
		res.LineValue = string(line)

		val := make([]byte, end-start)
		if _, err := r.ReadAt(val, int64(start-1)); err != nil {
			return Result{}, err
		}
		res.Source = string(val)
	}

	return res, nil
}
