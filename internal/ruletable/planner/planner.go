// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package planner

import (
	"context"
	"time"

	"github.com/google/cel-go/cel"
	celast "github.com/google/cel-go/common/ast"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
)

var ignoreHashFields = map[string]struct{}{
	"google.api.expr.v1alpha1.CheckedExpr.reference_map":  {},
	"google.api.expr.v1alpha1.CheckedExpr.type_map":       {},
	"google.api.expr.v1alpha1.CheckedExpr.source_info":    {},
	"google.api.expr.v1alpha1.CheckedExpr.expr_version":   {},
	"google.api.expr.v1alpha1.Expr.id":                    {},
	"google.api.expr.v1alpha1.Expr.CreateStruct.Entry.id": {},
}

type (
	QpN   = enginev1.PlanResourcesAst_Node
	qpNLO = enginev1.PlanResourcesAst_Node_LogicalOperation
	qpNE  = enginev1.PlanResourcesAst_Node_Expression
	RN    = struct {
		Node func() (*QpN, error)
		Role string
	}

	NodeFilter struct {
		allowFilter []*QpN
		denyFilter  []*QpN
	}
)

func (p *NodeFilter) Add(filter *QpN, effect effectv1.Effect) { _ = "STUB: not implemented"; return }

func (p *NodeFilter) DenyIsEmpty() bool { _ = "STUB: not implemented"; return false }

func (p *NodeFilter) AllowIsEmpty() bool { _ = "STUB: not implemented"; return false }

func (p *NodeFilter) ResetToUnconditionalDeny() { _ = "STUB: not implemented"; return }

func (p *NodeFilter) ToAST() *QpN { _ = "STUB: not implemented"; return nil }

// default to DENY

func MkPlanResourcesOutput(input *enginev1.PlanResourcesInput, matchedScopes map[string]string, validationErrors []*schemav1.ValidationError) *enginev1.PlanResourcesOutput {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func IsNodeConstBool(node *enginev1.PlanResourcesAst_Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func MkNodeFromLO(lo *enginev1.PlanResourcesAst_LogicalOperation) *enginev1.PlanResourcesAst_Node {
	_ = "STUB: not implemented"
	// node AND drNode
	return nil
}

func MkOrNode(nodes []*enginev1.PlanResourcesAst_Node) *enginev1.PlanResourcesAst_Node {
	_ = "STUB: not implemented"
	return nil
}

func MkAndNode(nodes []*enginev1.PlanResourcesAst_Node) *enginev1.PlanResourcesAst_Node {
	_ = "STUB: not implemented"
	return nil
}

func dedupNodes(nodes []*enginev1.PlanResourcesAst_Node) []*enginev1.PlanResourcesAst_Node {
	_ = "STUB: not implemented"
	return nil
}

// We hit a bug where Go map iteration was causing inconsistent ordering with `Expr_CreateStruct` entries. This led to
// outputs with many OR'd duplicate sets of conditions (namely in rules with wildcard roles and with many roles being resolved
// against those rules). I've tried to be a bit thorough in covering a few cases, but I'll admit that I'm shooting from
// the hip somewhat. However, this feels low risk given the context. Worst case, we'll miss a dedup opportunity and end up
// with duplicate nodes (logical no-ops), which is the situation we were in before.
// A clone probably isn't necessary, but it keeps the function side-effect free and predictable.
//nolint:forcetypeassert

func orderStructEntriesNode(node *enginev1.PlanResourcesAst_Node) {
	_ = "STUB: not implemented"
	return
}

func orderStructEntriesExpr(expr *exprpb.Expr) { _ = "STUB: not implemented"; return }

func structEntryKey(entry *exprpb.Expr_CreateStruct_Entry) string {
	_ = "STUB: not implemented"
	return ""
}

func mkOrLogicalOperation(nodes []*enginev1.PlanResourcesAst_Node) *enginev1.PlanResourcesAst_LogicalOperation {
	_ = "STUB: not implemented"
	return nil
}

func MkAndLogicalOperation(nodes []*enginev1.PlanResourcesAst_Node) *enginev1.PlanResourcesAst_LogicalOperation {
	_ = "STUB: not implemented"
	return nil
}

func MkFalseNode() *enginev1.PlanResourcesAst_Node { _ = "STUB: not implemented"; return nil }

func mkTrueNode() *enginev1.PlanResourcesAst_Node { _ = "STUB: not implemented"; return nil }

func InvertNodeBooleanValue(node *enginev1.PlanResourcesAst_Node) *enginev1.PlanResourcesAst_Node {
	_ = "STUB: not implemented"
	return nil
}

// No point NOT'ing a NOT. Therefore strip the existing NOT operator

type EvalContext struct {
	TimeFn func() time.Time
}

func (evalCtx *EvalContext) EvaluateCondition(ctx context.Context, condition *runtimev1.Condition, request *enginev1.Request, globals, constants map[string]any, variables map[string]celast.Expr, derivedRolesList func() (*exprpb.Expr, error)) (*enginev1.PlanResourcesAst_Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (evalCtx *EvalContext) evaluateConditionExpression(ctx context.Context, expr celast.Expr, request *enginev1.Request, globals, constants map[string]any, variables map[string]celast.Expr, derivedRolesList func() (*exprpb.Expr, error)) (*exprpb.CheckedExpr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ignore expressions that are invalid

type partialEvaluator struct {
	env       *cel.Env
	knownVars map[string]any
	vars      interpreter.PartialActivation
	nowFn     func() time.Time
}

func (p *partialEvaluator) evaluateUnknown(ctx context.Context, residual celast.Expr) (_ *exprpb.CheckedExpr, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *partialEvaluator) evalPartially(ctx context.Context, e celast.Expr) (ref.Val, celast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(ref.Val), *new(celast.Expr), nil
}

func newPartialEvaluator(env *cel.Env, knownVars map[string]any, nowFn func() time.Time) (*partialEvaluator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (evalCtx *EvalContext) newEvaluator(request *enginev1.Request, globals, constants map[string]any) (p *partialEvaluator, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// qualified, unqualified name

func (p *partialEvaluator) evalComprehensionBody(ctx context.Context, e celast.Expr) (celast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(celast.Expr), nil
}

// TODO(dbuduev): is this (still) necessary?
func evalComprehensionBodyImpl(ctx context.Context, env *cel.Env, pvars interpreter.PartialActivation, nowFn func() time.Time, e celast.Expr) (celast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(celast.Expr), nil
}

func residualExpr(ast *celast.AST, details *cel.EvalDetails) celast.Expr {
	_ = "STUB: not implemented"
	return *new(celast.Expr)
}

func VariableExprs(variables []*runtimev1.Variable) (map[string]celast.Expr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PlanResourcesInputToRequest(input *enginev1.PlanResourcesInput) *enginev1.Request {
	_ = "STUB: not implemented"
	return nil
}

func replaceRuntimeEffectiveDerivedRoles(expr celast.Expr, derivedRolesList func() (celast.Expr, error)) (celast.Expr, error) {
	_ = "STUB: not implemented"
	return *new(celast.Expr), nil
}

func isRuntimeEffectiveDerivedRoles(expr celast.SelectExpr) bool {
	_ = "STUB: not implemented"
	return false
}

func MkDerivedRolesList(derivedRoles []RN) func() (*exprpb.Expr, error) {
	_ = "STUB: not implemented"
	return nil
}

func mkBinaryOperatorExpr(op string, args ...*exprpb.Expr) *exprpb.Expr {
	_ = "STUB: not implemented"
	return nil
}

func derivedRoleListElement(derivedRole RN) (*exprpb.Expr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func qpNToExpr(node *QpN) (*exprpb.Expr, error) { _ = "STUB: not implemented"; return nil, nil }

func memoize[T any](f func() (T, error)) func() (T, error) { _ = "STUB: not implemented"; return nil }
