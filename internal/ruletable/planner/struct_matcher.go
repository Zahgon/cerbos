// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package planner

import (
	"context"

	celast "github.com/google/cel-go/common/ast"
	"github.com/google/cel-go/common/types/ref"

	"github.com/google/cel-go/common/operators"
)

type exprMatcherFunc func(e celast.Expr) (bool, []celast.Expr)

type exprMatcher struct {
	f  exprMatcherFunc
	ns []*exprMatcher // argument matchers
}

type expressionProcessor interface {
	Process(ctx context.Context, e celast.Expr) (bool, celast.Expr, error)
}

type processors []expressionProcessor

func (p processors) Process(ctx context.Context, e celast.Expr) (bool, celast.Expr, error) {
	_ = "STUB: not implemented"
	return false, *new(celast.Expr), nil
}

func (m *exprMatcher) run(e celast.Expr) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func mkConstExprMatcher(s *structMatcher) *exprMatcher { _ = "STUB: not implemented"; return nil }

func mkStructIndexerExprMatcher(s *structMatcher) *exprMatcher {
	_ = "STUB: not implemented"
	return nil
}

// expression: indexExpr <function> <const>
// indexExpr: structExpr[indexerExpr].
type structMatcher struct {
	mapExpr     celast.MapExpr
	indexerExpr celast.Expr
	constExpr   celast.Expr
	rootMatch   *exprMatcher
	function    string
	field       string // optional field. E.g. P.attr[R.id].role == "OWNER"
}

func (s *structMatcher) Process(_ context.Context, e celast.Expr) (bool, celast.Expr, error) {
	_ = "STUB: not implemented"
	return false, *new(celast.Expr), nil
}

// need to sort only to make the tests deterministic

var supportedOps = map[string]struct{}{
	operators.Equals:        {},
	operators.NotEquals:     {},
	operators.Less:          {},
	operators.LessEquals:    {},
	operators.Greater:       {},
	operators.GreaterEquals: {},
}

func newExpressionProcessor(p *partialEvaluator) expressionProcessor {
	_ = "STUB: not implemented"
	return *new(expressionProcessor)
}

func mkLogicalOr(args []celast.Expr) celast.Expr {
	_ = "STUB: not implemented"
	return *new(celast.Expr)
}

func mkLogicalAnd(args []celast.Expr) celast.Expr {
	_ = "STUB: not implemented"
	return *new(celast.Expr)
}

func mkOption(op string, key, val, expr, constExpr celast.Expr) celast.Expr {
	_ = "STUB: not implemented"
	return *new(celast.Expr)
}

type lambdaMatcher struct {
	iterRange        celast.Expr
	innerExpr        celast.Expr
	rootMatcher      *exprMatcher
	partialEvaluator *partialEvaluator
	iterVar          string
	iterVar2         string
}

func containsOnlyKnownValues(expr celast.Expr) bool { _ = "STUB: not implemented"; return false }

// For other types like IdentKind, SelectKind, CallKind, ComprehensionKind, etc.
// These typically involve variables or complex expressions

func (l *lambdaMatcher) Process(ctx context.Context, e celast.Expr) (bool, celast.Expr, error) {
	_ = "STUB: not implemented"
	return false, *new(celast.Expr), nil
}

func (l *lambdaMatcher) evaluateIterVar(ctx context.Context, iterVar celast.Expr) (ref.Val, error) {
	_ = "STUB: not implemented"
	return *new(ref.Val), nil
}

const nLambdaVars = 2

func (l *lambdaMatcher) evaluateExpr(ctx context.Context, knownVars map[string]any) (celast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(celast.Expr), nil
}

// ast = celast.NewAST(l.innerExpr, nil)
