// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"github.com/goccy/go-yaml/ast"

	sourcev1 "github.com/cerbos/cerbos/api/genpb/cerbos/source/v1"
)

type SourceCtx struct {
	doc *ast.DocumentNode
	*sourcev1.SourceContext
}

func newSourceCtx(srcCtx *sourcev1.SourceContext, doc *ast.DocumentNode) SourceCtx {
	_ = "STUB: not implemented"
	return *new(SourceCtx)
}

func NewEmptySourceCtx() SourceCtx { _ = "STUB: not implemented"; return *new(SourceCtx) }

func (sc SourceCtx) StartPosition() *sourcev1.Position { _ = "STUB: not implemented"; return nil }

func (sc SourceCtx) PositionOfMapKeyAtProtoPath(path string) *sourcev1.Position {
	_ = "STUB: not implemented"
	return nil
}

func (sc SourceCtx) PositionOfValueAtProtoPath(path string) *sourcev1.Position {
	_ = "STUB: not implemented"
	return nil
}

func (sc SourceCtx) ContextForMapKeyAtYAMLPath(path string) string {
	_ = "STUB: not implemented"
	return ""
}

func (sc SourceCtx) ContextForValueAtYAMLPath(path string) string {
	_ = "STUB: not implemented"
	return ""
}

func (sc SourceCtx) nodeAtYAMLPath(path string) ast.Node {
	_ = "STUB: not implemented"
	return *new(ast.Node)
}

func (sc SourceCtx) contextForNode(node ast.Node) string { _ = "STUB: not implemented"; return "" }

func (sc SourceCtx) PositionAndContextForMapKeyAtProtoPath(path string) (pos *sourcev1.Position, context string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func (sc SourceCtx) PositionAndContextForValueAtProtoPath(path string) (pos *sourcev1.Position, context string) {
	_ = "STUB: not implemented"
	return nil, ""
}
