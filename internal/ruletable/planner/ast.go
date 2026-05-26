// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package planner

import (
	"errors"
	"strings"

	celast "github.com/google/cel-go/common/ast"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/types/known/structpb"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
)

const (
	Or                 = "or"
	And                = "and"
	Not                = "not"
	Equals             = "eq"
	NotEquals          = "ne"
	GreaterThan        = "gt"
	GreaterThanOrEqual = "ge"
	LessThan           = "lt"
	LessThanOrEqual    = "le"
	In                 = "in"
	List               = "list"
	Struct             = "struct"
	Add                = "add"
	Sub                = "sub"
	Mult               = "mult"
	Div                = "div"
	Mod                = "mod"
	SetField           = "set-field"
	GetField           = "get-field"
	Index              = "index"
	All                = "all"
	Filter             = "filter"
	TransformMap       = "transformMap"
	TransformMapEntry  = "transformMapEntry"
	TransformList      = "transformList"
	Exists             = "exists"
	ExistsOne          = "exists_one"
	Map                = "map"
	Lambda             = "lambda"
	If                 = "if"
)

var ErrUnknownOperator = errors.New("unknown operator")

func opFromCLE(fn string) (string, error) { _ = "STUB: not implemented"; return "", nil }

type replaceVarsFunc func(e celast.Expr) (output celast.Expr, matched bool, err error)

func replaceVarsGen(e celast.Expr, f replaceVarsFunc) (output celast.Expr, err error) {
	_ = "STUB: not implemented"
	return *new(celast.Expr), nil
}

// This functions wraps references to known resource attributes in an `id` function call, which simply returns its argument.
// E.g. Replace R.attr.field1 with id(R.attr.field1) iif R.attr.field1 is passed in the request to the Query Planner API.
// This trick is necessary to evaluate expression like `P.attr.struct1[R.attr.field1]`, otherwise CEL tries to use `R.attr.field1`
// as a qualifier for `P.attr.struct1` and produces the error https://github.com/cerbos/cerbos/issues/1340
func replaceResourceVals(e celast.Expr, vals map[string]*structpb.Value) (output celast.Expr, err error) {
	_ = "STUB: not implemented"
	return *new(celast.Expr), nil
}

// match R.attr.<field>

// match request.resource.attr.<field>

func replaceVars(e celast.Expr, vars map[string]celast.Expr) (output celast.Expr, err error) {
	_ = "STUB: not implemented"
	return *new(celast.Expr), nil
}

func convert(expr *enginev1.PlanResourcesAst_Node, acc *enginev1.PlanResourcesFilter_Expression_Operand) error {
	_ = "STUB: not implemented"
	return nil
}

func mkConstStringExpr(s string) *exprpb.Expr { _ = "STUB: not implemented"; return nil }

func structKeys(x *exprpb.Expr_CreateStruct) []*exprpb.Expr { _ = "STUB: not implemented"; return nil }

// sort by expression, ignoring AST node id

func mkExprOpExpr(op string, args ...*enginev1.PlanResourcesFilter_Expression_Operand) *exprOpExpr {
	_ = "STUB: not implemented"
	return nil
}

func buildExpr(expr *exprpb.Expr, acc *enginev1.PlanResourcesFilter_Expression_Operand) error {
	_ = "STUB: not implemented"
	return nil
}

type (
	exprOp      = enginev1.PlanResourcesFilter_Expression_Operand
	exprOpExpr  = enginev1.PlanResourcesFilter_Expression_Operand_Expression
	exprOpValue = enginev1.PlanResourcesFilter_Expression_Operand_Value
	exprOpVar   = enginev1.PlanResourcesFilter_Expression_Operand_Variable
)

func buildExprImpl(cur *exprpb.Expr, acc *enginev1.PlanResourcesFilter_Expression_Operand, parent *exprpb.Expr) error {
	_ = "STUB: not implemented"
	return nil
}

// This is a compound "a.b.c" variable

// only values in list, so acc.Node is a list of values

//nolint:forcetypeassert

// list of expressions

// convert struct to a list, if it is a key membership test operation

// right-hand side arg index

func (lambdaAst *lambdaAST) mkNode(cur *exprpb.Expr) (*exprOpExpr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lambdaAst *lambdaAST) buildIterRangeOp(cur *exprpb.Expr) (*exprOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildLambdaExprOp(expr, cur *exprpb.Expr) (*exprOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func visitConst(c *exprpb.Constant) (*structpb.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToFilter(plan *enginev1.PlanResourcesAst_Node) (*enginev1.PlanResourcesFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normaliseFilter(filter *enginev1.PlanResourcesFilter) *enginev1.PlanResourcesFilter {
	_ = "STUB: not implemented"
	return nil
}

func normaliseFilterExprOp(cond *enginev1.PlanResourcesFilter_Expression_Operand) *enginev1.PlanResourcesFilter_Expression_Operand {
	_ = "STUB: not implemented"
	return nil
}

var (
	falseExprOpValue = &enginev1.PlanResourcesFilter_Expression_Operand{
		Node: &enginev1.PlanResourcesFilter_Expression_Operand_Value{
			Value: structpb.NewBoolValue(false),
		},
	}
	trueExprOpValue = &enginev1.PlanResourcesFilter_Expression_Operand{
		Node: &enginev1.PlanResourcesFilter_Expression_Operand_Value{
			Value: structpb.NewBoolValue(true),
		},
	}
)

func normaliseFilterExprOpExpr(expr *enginev1.PlanResourcesFilter_Expression_Operand_Expression) *enginev1.PlanResourcesFilter_Expression_Operand {
	_ = "STUB: not implemented"
	return nil
}

// Ignore literal true values because they don't matter

// Ignore literal false values because they don't matter

/* && !boolVal*/
// A literal false makes the whole AND expression return false

// A literal true makes the whole OR expression return true

// Ignore repeated values in and/or

// AND or OR of a single value is the value itself
//nolint:nestif

// because all true operands were removed, the result simplifies to true (true && true == true)

// because all false operands were removed, the result simplifies to false (false || false == false)

// AND or OR of a single value is the value itself

// NOT of a single bool value

// normaliseInExpr normalises an IN expression in place.
// If the return value is nil, then the expression can be simplified further by other normalisers.
func normaliseInExpr(expr *enginev1.PlanResourcesFilter_Expression_Operand_Expression) *enginev1.PlanResourcesFilter_Expression_Operand {
	_ = "STUB: not implemented"
	return nil
}

// operand is not a list or map so convert it to a EQ

func normaliseFilterExprOpVar(v *enginev1.PlanResourcesFilter_Expression_Operand_Variable) *enginev1.PlanResourcesFilter_Expression_Operand {
	_ = "STUB: not implemented"
	return nil
}

func asBoolValue(op *enginev1.PlanResourcesFilter_Expression_Operand) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func FilterToString(filter *enginev1.PlanResourcesFilter) string {
	_ = "STUB: not implemented"
	return ""
}

func filterExprOpToString(b *strings.Builder, cond *enginev1.PlanResourcesFilter_Expression_Operand) {
	_ = "STUB: not implemented"
	return
}

func filterExprOpExprToString(b *strings.Builder, expr *enginev1.PlanResourcesFilter_Expression) {
	_ = "STUB: not implemented"
	return
}

func canOperateOnStruct(op string) bool { _ = "STUB: not implemented"; return false }
