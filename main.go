package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
	"github.com/urfave/cli/v2"

	"github.com/nick-jones/astpath/internal/readutil"
	"github.com/nick-jones/astpath/pkg/astxml"
)

func main() {
	app := &cli.App{
		Name: "astpath",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "print-mode",
				Value: "source",
				Usage: "choose from: source, source-line, xml-inner, xml-outer",
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type printMode string

const (
	printModeSource     printMode = "source"      // prints the value indicated by the original token positions
	printModeSourceLine printMode = "source-line" // prints the full line indicated by the original token start position
	printModeInnerXML   printMode = "xml-inner"   // prints the raw XML excluding the selected node
	printModeOuterXML   printMode = "xml-outer"   // prints the raw XML including the selected node
)

func run(c *cli.Context) error {
	query, root := c.Args().Get(0), c.Args().Get(1)
	if query == "" {
		return fmt.Errorf("xpath must be provided")
	}
	if root == "" {
		root = "."
	}

	mode := printMode(c.String("print-mode"))

	expr, err := xpath.Compile(query)
	if err != nil {
		return err
	}

	paths, err := findFiles(root)
	if err != nil {
		return err
	}

	for _, path := range paths {
		if err := evaluateFile(path, expr, mode); err != nil {
			return err
		}
	}

	return nil
}

func findFiles(root string) (files []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		files = append(files, path)

		return nil
	})
	return
}

func evaluateFile(path string, expr *xpath.Expr, mode printMode) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fset := token.NewFileSet()
	src, err := parser.ParseFile(fset, path, f, 0)
	if err != nil {
		return err
	}

	xml, err := astToXML(src)
	if err != nil {
		return err
	}

	doc, err := xmlquery.Parse(bytes.NewReader(xml))
	if err != nil {
		return err
	}

	nodes := xmlquery.QuerySelectorAll(doc, expr)

	return printResults(nodes, f, fset, mode)
}

func astToXML(f *ast.File) ([]byte, error) {
	xml, err := astxml.Marshal(f)
	if err != nil {
		return nil, err
	}

	return append([]byte(`<?xml version="1.0" encoding="UTF-8"?>`), xml...), nil
}

func printResults(nodes []*xmlquery.Node, r io.ReaderAt, fset *token.FileSet, mode printMode) error {
	for _, node := range nodes {
		switch mode {
		case printModeInnerXML:
			fmt.Println(node.OutputXML(false))
		case printModeOuterXML:
			if node.Parent == nil {
				return fmt.Errorf("cannot print outer of root")
			}
			fmt.Println(node.OutputXML(true))
		case printModeSourceLine:
			start, _, err := positionsFromNode(node)
			if err != nil {
				return err
			}
			pos := fset.Position(token.Pos(start))
			line, err := readutil.ReadLine(r, int64(start))
			if err != nil {
				return err
			}
			fmt.Printf("%s > %s\n", pos, string(line))
		case printModeSource:
			start, end, err := positionsFromNode(node)
			if err != nil {
				return err
			}
			pos := fset.Position(token.Pos(start))
			val := make([]byte, end-start)
			if _, err := r.ReadAt(val, int64(start-1)); err != nil {
				return err
			}
			fmt.Printf("%s > %s\n", pos, string(val))
		default:
			return fmt.Errorf("unrecognised mode: %v", mode)
		}
	}
	return nil
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
