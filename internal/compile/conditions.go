// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package compile

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
)

func Condition(cond *policyv1.Condition) (*runtimev1.Condition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compileCondition(modCtx *moduleCtx, path string, cond *policyv1.Condition, markReferencedConstantsAndVariablesAsUsed bool) *runtimev1.Condition {
	_ = "STUB: not implemented"
	return nil
}

func compileMatch(modCtx *moduleCtx, path string, match *policyv1.Match, markReferencedConstantsAndVariablesAsUsed bool) *runtimev1.Condition {
	_ = "STUB: not implemented"
	return nil
}

func compileCELExpr(modCtx *moduleCtx, path, expr string, markReferencedConstantsAndVariablesAsUsed bool) *runtimev1.Expr {
	_ = "STUB: not implemented"
	return nil
}

func compileMatchList(modCtx *moduleCtx, path string, matches []*policyv1.Match, markReferencedVariablesAsUsed bool) *runtimev1.Condition_ExprList {
	_ = "STUB: not implemented"
	return nil
}
