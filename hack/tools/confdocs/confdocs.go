// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build confdocs
// +build confdocs

package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"go.uber.org/zap"
	"golang.org/x/tools/go/packages"
)

const (
	interfacePackage   = "github.com/cerbos/cerbos/internal/config"
	interfaceName      = "Section"
	defaultLogLevel    = "ERROR"
	keyRequired        = "required"
	keyOptional        = "optional"
	optionExampleValue = "example"
	optionIgnore       = "ignore"
)

//go:embed generator.go.tmpl
var templateText string

var (
	errInterfaceNotFound = errors.New("interface not found")
	errTagNotExists      = errors.New("yaml tag does not exist")
	descRegex            = regexp.MustCompile(`\+desc=(.+)`)
)

var (
	rootDir    = flag.String("rootDir", ".", "Root directory to scan")
	outputFile = flag.String("outFile", "docs/modules/configuration/partials/fullconfiguration.adoc", "Path to output the content")

	logger      *zap.SugaredLogger
	excludeObjs = map[string]struct{}{"CompilationUnit": {}, "Section": {}}
)

type StructInfo struct {
	Pkg           string
	Name          string
	Documentation string
	Fields        []FieldInfo
}

type FieldInfo struct {
	Name          string
	Documentation string
	Tag           string
	Fields        []FieldInfo
	Array         bool
}

type TagInfo struct {
	DefaultValue string
	Name         string
	Ignore       bool
	Required     bool
}

type Output struct {
	Imports  map[string]string
	Sections map[string]string
	File     string
}

func init() {
	if envLevel := os.Getenv("CONFDOCS_LOG_LEVEL"); envLevel != "" {
		doInitLogging(envLevel)
		return
	}
	doInitLogging(defaultLogLevel)
}

func main() {
	flag.Parse()
	absRootDir, err := filepath.Abs(*rootDir)
	if err != nil {
		logger.Fatalf("Failed to find absolute path to %q: %v", *rootDir, err)
	}

	absOutputFile, err := filepath.Abs(*outputFile)
	if err != nil {
		logger.Fatalf("Failed to find absolute path to %q: %v", *outputFile, err)
	}

	pkgs, err := loadPackages(absRootDir)
	if err != nil {
		logger.Fatalf("failed to load packages: %v", err)
	}

	iface, err := findInterfaceDef(pkgs)
	if err != nil {
		logger.Fatalf("failed to find %s.%s: %v", interfacePackage, interfaceName, err)
	}

	structs := findIfaceImplementors(iface, pkgs)

	output := Output{
		File:     absOutputFile,
		Imports:  make(map[string]string),
		Sections: make(map[string]string),
	}
	for i, s := range structs {
		docs := genDocs(s)
		// Skip empty sections (structs with only ignored fields)
		if strings.TrimSpace(docs) == "" {
			continue
		}
		imp := fmt.Sprintf("c%d", i)
		output.Imports[imp] = s.Pkg
		output.Sections[fmt.Sprintf("%s.%s", imp, s.Name)] = docs
	}

	tmpl, err := template.New("generator.go").Parse(templateText)
	if err != nil {
		logger.Fatalf("failed to parse template: %v", err)
	}

	if err := tmpl.Execute(os.Stdout, output); err != nil {
		logger.Fatalf("failed to render template: %v", err)
	}
}

func loadPackages(pkgDir string) ([]*packages.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findInterfaceDef(pkgs []*packages.Package) (*types.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findIfaceImplementors(iface *types.Interface, pkgs []*packages.Package) []*StructInfo {
	_ = "STUB: not implemented"
	return nil

	// `traversedObjMap` is used to store fields against the object FQN as we traverse the AST.
	// This allows us to assign fields for non-local embedded structs on the fly by constructing
	// the FQN and retrieving the fields from the map.
	// We don't have to wait until traversal completion as embedded structs should already be available
	// in the map given the depth-first traversal order of the AST.
}

func implementsIface(iface *types.Interface, obj types.Object) bool {
	_ = "STUB: not implemented"
	return false
}

func inspect(pkg *packages.Package, obj types.Object, traversedObjMap map[string][]FieldInfo) *StructInfo {
	_ = "STUB: not implemented"
	return nil
}

func find(files []*ast.File, objName string) (*ast.TypeSpec, *ast.CommentGroup) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inspectStruct(node ast.Expr, info *types.Info, traversedObjMap map[string][]FieldInfo) []FieldInfo {
	_ = "STUB: not implemented"
	return nil
}

// Handle non-local embedded structs

func genDocs(si *StructInfo) string { _ = "STUB: not implemented"; return "" }

func doGenDocs(out io.Writer, si *StructInfo, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func walkFields(out io.Writer, fields []FieldInfo, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func indentf(out io.Writer, n int, format string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTag(tag string) (*TagInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func parseDescMarker(cg *ast.CommentGroup) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

type finder struct {
	typeSpec     *ast.TypeSpec
	commentGroup *ast.CommentGroup
	objName      string
}

func (f *finder) Visit(n ast.Node) ast.Visitor { _ = "STUB: not implemented"; return *new(ast.Visitor) }

func (f finder) findReceiverObj(tpe ast.Expr) *ast.StructType {
	_ = "STUB: not implemented"
	return nil
}

func doInitLogging(level string) { _ = "STUB: not implemented"; return }
