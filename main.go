package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"

	"github.com/nick-jones/astpath/pkg/query"
)

func main() {
	app := &cli.App{
		Name: "astpath",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "template",
				Value: "{{.Filename}}:{{.Line}}:{{.Column}} > {{.Source}}",
				Usage: "text/template format",
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	xpath, root := c.Args().Get(0), c.Args().Get(1)
	if xpath == "" {
		return fmt.Errorf("xpath must be provided")
	}
	if root == "" {
		root = "."
	}

	expr, err := query.Compile(xpath)
	if err != nil {
		return fmt.Errorf("invalid query: %w", err)
	}

	tmpl, err := template.New("format").Parse(c.String("template"))
	if err != nil {
		return fmt.Errorf("failed to parse format flag: %w", err)
	}

	paths, err := findFiles(root)
	if err != nil {
		return err
	}

	results := make([]query.Result, 0)
	for _, path := range paths {
		res, err := queryFile(path, expr)
		if err != nil {
			return err
		}
		results = append(results, res...)
	}

	for _, res := range results {
		if err := tmpl.Execute(os.Stdout, res); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		fmt.Println()
	}

	return nil
}

func findFiles(root string) (paths []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		paths = append(paths, path)

		return nil
	})
	return
}

func queryFile(path string, expr *query.Expr) ([]query.Result, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return query.FindAll(f, path, expr)
}