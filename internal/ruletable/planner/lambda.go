// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package planner

import (
	"errors"

	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type lambdaAST struct {
	iterRange *exprpb.Expr
	expr      *exprpb.Expr
	expr2     *exprpb.Expr
	iterVar   string
	iterVar2  string
	operator  string
}

func getTransformMapExpression(w *wrapper) (string, *exprpb.Expr, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

var ErrExpectedSortBy = errors.New("expected sortBy comprehension")

const sortByFuncName = "sortBy"

func mkSortByAST(e *exprpb.Expr_Comprehension) (*lambdaAST, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildLambdaAST(e *exprpb.Expr_Comprehension) (*lambdaAST, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type wrapper exprpb.Expr

func (w *wrapper) e() *exprpb.Expr { _ = "STUB: not implemented"; return nil }

func (w *wrapper) getArg(i int) *wrapper { _ = "STUB: not implemented"; return nil }

func (w *wrapper) getArgsLen() int { _ = "STUB: not implemented"; return 0 }

func (w *wrapper) getListElement(i int) *wrapper { _ = "STUB: not implemented"; return nil }

func (w *wrapper) Function() string { _ = "STUB: not implemented"; return "" }
