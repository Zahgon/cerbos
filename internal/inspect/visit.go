// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package inspect

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/ruletable/index"
	"github.com/google/cel-go/common/ast"
)

func visitBinding(b *index.Binding, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitCompiledPolicySet(policySet *runtimev1.RunnablePolicySet, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitPolicy(policy *policyv1.Policy, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func visitVariables(variables map[string]string, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitExpr(expr string, visitor ast.Visitor) error { _ = "STUB: not implemented"; return nil }

func visitCondition(condition *policyv1.Condition, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitCompiledCondition(condition *runtimev1.Condition, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitCompiledConditionExprList(exprList *runtimev1.Condition_ExprList, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitCompiledExpr(expr *runtimev1.Expr, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitOutput(output *policyv1.Output, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func visitCompiledOutput(output *runtimev1.Output, visitor ast.Visitor) error {
	_ = "STUB: not implemented"
	return nil
}
