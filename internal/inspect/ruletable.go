// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package inspect

import (
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/ruletable/index"
	"github.com/cerbos/cerbos/internal/util"
)

// BindingSource provides access to index bindings for inspection.
// It exists here to prevent a circular dependency between `inspect`, `ruletable` and `storage`.
type BindingSource interface {
	GetAllRows() []*index.Binding
}

func RuleTables(src BindingSource) *RuleTable { _ = "STUB: not implemented"; return nil }

// RuleTable is not safe for concurrent use: the internal caches are unsynchronized.
type RuleTable struct {
	src            BindingSource
	results        map[string]*responsev1.InspectPoliciesResponse_Result
	constantsCache map[*index.FunctionalCore][]*responsev1.InspectPoliciesResponse_Constant
	variablesCache map[*index.FunctionalCore][]*responsev1.InspectPoliciesResponse_Variable
}

// Inspect inspects the given rule table and caches the inspection related information internally.
func (rt *RuleTable) Inspect() error { _ = "STUB: not implemented"; return nil }

func (rt *RuleTable) inspectBinding(
	b *index.Binding,
	actionSet util.StringSet,
	attrSet util.StringSet,
	constantSet util.StringSet,
	derivedRoleSet util.StringSet,
	variableSet util.StringSet,
	result *responsev1.InspectPoliciesResponse_Result,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Results returns the final inspection results.
func (rt *RuleTable) Results() map[string]*responsev1.InspectPoliciesResponse_Result {
	_ = "STUB: not implemented"
	return nil
}

func getOrInitStringSet(m map[string]util.StringSet, key string) util.StringSet {
	_ = "STUB: not implemented"
	return *new(util.StringSet)
}

func listActions(b *index.Binding) []string { _ = "STUB: not implemented"; return nil }

// Results are cached per-Core because many bindings share the same Core pointer
// (they differ only in routing dimensions) and structpb conversion is non-trivial.
func (rt *RuleTable) listConstants(b *index.Binding) ([]*responsev1.InspectPoliciesResponse_Constant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDerivedRole(b *index.Binding) *responsev1.InspectPoliciesResponse_DerivedRole {
	_ = "STUB: not implemented"
	return nil
}

// Results are cached per-Core because many bindings share the same Core pointer
// (they differ only in routing dimensions).
func (rt *RuleTable) listVariables(b *index.Binding) []*responsev1.InspectPoliciesResponse_Variable {
	_ = "STUB: not implemented"
	return nil
}
