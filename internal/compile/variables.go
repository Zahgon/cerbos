// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package compile

import (
	"errors"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
)

var errInvalidVariableName = errors.New("invalid variable name")

func compilePolicyVariables(modCtx *moduleCtx, variables *policyv1.Variables) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck

func compileExportVariables(modCtx *moduleCtx) { _ = "STUB: not implemented"; return }

func sortCompiledVariables(fqn string, variables map[string]*runtimev1.Expr) ([]*runtimev1.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type variableNode struct {
	*runtimev1.Variable
	varCtx *variableCtx
	id     int64
}

func (v *variableNode) ID() int64 { _ = "STUB: not implemented"; return 0 }

type variableDefinitions struct {
	modCtx  *moduleCtx
	graph   *simple.DirectedGraph
	ids     map[string]int64
	sources map[string][]*variableCtx
	used    map[string]struct{}
	nextID  int64
}

func newVariableDefinitions(modCtx *moduleCtx) *variableDefinitions {
	_ = "STUB: not implemented"
	return nil
}

func (vd *variableDefinitions) Compile(definitions map[string]string, path, source string) {
	_ = "STUB: not implemented"
	return
}

func (vd *variableDefinitions) Import(from *moduleCtx, source string) {
	_ = "STUB: not implemented"
	return
}

//nolint:forcetypeassert

func (vd *variableDefinitions) Add(variable *runtimev1.Variable, varCtx *variableCtx) {
	_ = "STUB: not implemented"
	return
}

func (vd *variableDefinitions) Resolve() { _ = "STUB: not implemented"; return }

func (vd *variableDefinitions) reportRedefinedVariables() { _ = "STUB: not implemented"; return }

//nolint:mnd

func variableDefinitionPlaces(contexts []*variableCtx) []string {
	_ = "STUB: not implemented"
	return nil
}

func (vd *variableDefinitions) resolveReferences() { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

func (vd *variableDefinitions) references(path string, expr *expr.CheckedExpr) (constants, variables map[string]struct{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vd *variableDefinitions) ResetUsage() { _ = "STUB: not implemented"; return }

func (vd *variableDefinitions) Use(path string, expr *expr.CheckedExpr) {
	_ = "STUB: not implemented"
	return
}

func (vd *variableDefinitions) use(id int64, name string) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

//nolint:forcetypeassert

func (vd *variableDefinitions) Used() ([]*runtimev1.Variable, map[string]*runtimev1.Expr) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vd *variableDefinitions) All() ([]*runtimev1.Variable, map[string]*runtimev1.Expr) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vd *variableDefinitions) list(includeUnused bool) ([]*runtimev1.Variable, map[string]*runtimev1.Expr) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forcetypeassert

func (vd *variableDefinitions) reportCyclicalVariables(cycles [][]graph.Node) {
	_ = "STUB: not implemented"
	return
}

//nolint:mnd

//nolint:forcetypeassert

//nolint:forcetypeassert
//nolint:govet

//nolint:govet

func sortVariablesByName(nodes []graph.Node) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert
