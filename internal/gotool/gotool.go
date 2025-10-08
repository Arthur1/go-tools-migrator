package gotool

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/natefinch/atomic"
	"golang.org/x/mod/modfile"
)

func Migrate(toolsGoFilePath, goModFilePath string, dryRun bool) (newGoModContent string, err error) {
	pkgs, err := parseToolsGoFile(toolsGoFilePath)
	if err != nil {
		return
	}

	goMod, err := readGoModFile(goModFilePath)
	if err != nil {
		return
	}
	if err = appendTools(goMod, pkgs); err != nil {
		return
	}

	newGoModBytes, err := goMod.Format()
	if err != nil {
		return
	}

	if !dryRun {
		if err = atomic.WriteFile(goModFilePath, bytes.NewReader(newGoModBytes)); err != nil {
			return
		}
		if err = os.Remove(toolsGoFilePath); err != nil {
			return
		}
	}

	newGoModContent = string(newGoModBytes)
	return
}

func parseToolsGoFile(filePath string) ([]string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %w", err)
	}
	v := new(toolsGoFileVisitor)
	ast.Walk(v, file)
	return v.packages, nil
}

type toolsGoFileVisitor struct {
	packages []string
}

func (v *toolsGoFileVisitor) Visit(node ast.Node) ast.Visitor {
	if imp, ok := node.(*ast.ImportSpec); ok {
		path := strings.Trim(imp.Path.Value, `"`)
		v.packages = append(v.packages, path)
	}
	return v
}

func readGoModFile(filePath string) (*modfile.File, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return modfile.Parse(filePath, b, nil)
}

func appendTools(goMod *modfile.File, toolPkgs []string) error {
	var merr error
	for _, toolPkg := range toolPkgs {
		if err := goMod.AddTool(toolPkg); err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	return merr
}
