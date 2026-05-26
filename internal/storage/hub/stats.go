// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package hub

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/ruletable/index"
	"github.com/cerbos/cerbos/internal/storage"
)

const numPolicyKinds = 4

type statsCollector struct {
	policies          map[uint64]*stats
	uniqueRules       map[uint64]struct{}
	uniqueResources   map[uint64]struct{}
	uniqueActions     map[uint64]struct{}
	schemaRefs        map[uint64]struct{}
	hasOutput         bool
	hasScopedPolicies bool
}

type stats struct {
	kind           policy.Kind
	ruleCount      uint32
	conditionCount uint32
}

type policyStats struct {
	stats
	hasOutput bool
	scoped    bool
}

func newStatsCollector() *statsCollector { _ = "STUB: not implemented"; return nil }

func (s *statsCollector) collate() storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (s *statsCollector) addRow(row *index.Binding) { _ = "STUB: not implemented"; return }

func (s *statsCollector) addRunnablePolicySet(rps *runtimev1.RunnablePolicySet) {
	_ = "STUB: not implemented"
	return
}

func (s *statsCollector) procRunnablePrincipalPolicySet(rpps *runtimev1.RunnablePrincipalPolicySet) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) procRunnableResourcePolicySet(rrps *runtimev1.RunnableResourcePolicySet) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) procRunnableRolePolicySet(rrps *runtimev1.RunnableRolePolicySet) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) addSchemas(schemas *policyv1.Schemas) { _ = "STUB: not implemented"; return }
