// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	celast "github.com/google/cel-go/common/ast"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type IDGen struct {
	ids map[int64]int64
	c   int64
}

func NewIDGen() *IDGen { _ = "STUB: not implemented"; return nil }

func (g *IDGen) Remap(id int64) int64 { _ = "STUB: not implemented"; return 0 }

func RenumberIDs(e celast.Expr) { _ = "STUB: not implemented"; return }

func ZeroIDs(e celast.Expr) { _ = "STUB: not implemented"; return }

func UpdateIDs(e *exprpb.Expr) { _ = "STUB: not implemented"; return }

func MkListExpr(elems []celast.Expr) celast.Expr {
	_ = "STUB: not implemented"
	return *new(celast.Expr)
}

func MkListExprProto(elems []*exprpb.Expr) *exprpb.Expr { _ = "STUB: not implemented"; return nil }

func MkCallExpr(op string, args ...celast.Expr) celast.Expr {
	_ = "STUB: not implemented"
	return *new(celast.Expr)
}

func MkCallExprProto(op string, args ...*exprpb.Expr) *exprpb.Expr {
	_ = "STUB: not implemented"
	return nil
}

func MkSelectExpr(operand celast.Expr, field string) celast.Expr {
	_ = "STUB: not implemented"
	return *new(celast.Expr)
}

func MkSelectExprProto(operand *exprpb.Expr, field string) *exprpb.Expr {
	_ = "STUB: not implemented"
	return nil
}
