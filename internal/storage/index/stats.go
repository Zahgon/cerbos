// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package index

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
)

const numPolicyKinds = 4

type statsCollector struct {
	policyCount           map[policy.Kind]int
	ruleCountPerKind      map[policy.Kind]int
	conditionCountPerKind map[policy.Kind]int
	maxRuleCount          map[policy.Kind]int
	maxConditionCount     map[policy.Kind]int
	schemaRefs            map[uint64]struct{}
	uniqueActions         map[uint64]struct{}
	uniqueResources       map[uint64]struct{}
	hasOutput             bool
	hasScopedPolicies     bool
}

type policyStats struct {
	ruleCount      int
	conditionCount int
	hasOutput      bool
	scoped         bool
}

func newStatsCollector() *statsCollector { _ = "STUB: not implemented"; return nil }

func (s *statsCollector) collate() storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (s *statsCollector) add(p policy.Wrapper) { _ = "STUB: not implemented"; return }

func (s *statsCollector) procDerivedRoles(dr *policyv1.DerivedRoles) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) procExportConstants(ev *policyv1.ExportConstants) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) procExportVariables(ev *policyv1.ExportVariables) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) procPrincipalPolicy(pp *policyv1.PrincipalPolicy) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) procResourcePolicy(rp *policyv1.ResourcePolicy) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

func (s *statsCollector) procRolePolicy(rp *policyv1.RolePolicy) (ps policyStats) {
	_ = "STUB: not implemented"
	return *new(policyStats)
}

// Role policies are modeled differently to resource/principal policies.
// We map a set of allowable actions for a given resource, so from a stats perspective,
// each allowable action is treated as an individual condition.
