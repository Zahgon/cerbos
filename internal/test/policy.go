// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build tests

package test

import (
	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type NameMod func(string) string

type PolicyBuilder interface {
	Build() *policyv1.Policy
}

// ResourceRuleBuilder is a builder for resource rules.
type ResourceRuleBuilder struct {
	rule *policyv1.ResourceRule
}

func NewResourceRule(actions ...string) *ResourceRuleBuilder { _ = "STUB: not implemented"; return nil }

func (rrb *ResourceRuleBuilder) WithRoles(roles ...string) *ResourceRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rrb *ResourceRuleBuilder) WithDerivedRoles(roles ...string) *ResourceRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rrb *ResourceRuleBuilder) WithMatchExpr(expr ...string) *ResourceRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rrb *ResourceRuleBuilder) WithScript(script string) *ResourceRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rrb *ResourceRuleBuilder) WithEffect(effect effectv1.Effect) *ResourceRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rrb *ResourceRuleBuilder) Build() *policyv1.ResourceRule {
	_ = "STUB: not implemented"

	// ResourcePolicyBuilder is a builder for resource policies.
	return nil
}

type ResourcePolicyBuilder struct {
	rp *policyv1.ResourcePolicy
}

func NewResourcePolicyBuilder(resource, version string) *ResourcePolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rpb *ResourcePolicyBuilder) WithDerivedRolesImports(imp ...string) *ResourcePolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rpb *ResourcePolicyBuilder) WithLocalConstant(name string, value *structpb.Value) *ResourcePolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rpb *ResourcePolicyBuilder) WithLocalVariable(name, value string) *ResourcePolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rpb *ResourcePolicyBuilder) WithRules(rules ...*policyv1.ResourceRule) *ResourcePolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rpb *ResourcePolicyBuilder) Build() *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func buildAndCondition(expr ...string) *policyv1.Condition { _ = "STUB: not implemented"; return nil }

func GenDisabledResourcePolicy(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func GenDisabledRolePolicy(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func GenScopedResourcePolicy(scope string, mod NameMod) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}

// GenResourcePolicy generates a sample resource policy with some names modified by the NameMod.
func GenResourcePolicy(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

// GenRolePolicy generates a sample role policy with some names modified by the NameMod.
func GenRolePolicy(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

// PrincipalRuleBuilder is a builder for principal rules.
type PrincipalRuleBuilder struct {
	rule *policyv1.PrincipalRule
}

func NewPrincipalRuleBuilder(resource string) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) AllowAction(action string) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) DenyAction(action string) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) AllowActionWhenMatch(action string, expr ...string) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) DenyActionWhenMatch(action string, expr ...string) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) AllowActionWhenScript(action, script string) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) DenyActionWhenScript(action, script string) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) addAction(action string, effect effectv1.Effect, comp *policyv1.Condition) *PrincipalRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (prb *PrincipalRuleBuilder) Build() *policyv1.PrincipalRule {
	_ = "STUB: not implemented"

	// PrincipalPolicyBuilder is a builder for principal policies.
	return nil
}

type PrincipalPolicyBuilder struct {
	pp *policyv1.PrincipalPolicy
}

func NewPrincipalPolicyBuilder(principal, version string) *PrincipalPolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (ppb *PrincipalPolicyBuilder) WithLocalConstant(name string, value *structpb.Value) *PrincipalPolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (ppb *PrincipalPolicyBuilder) WithLocalVariable(name, value string) *PrincipalPolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (ppb *PrincipalPolicyBuilder) WithRules(rules ...*policyv1.PrincipalRule) *PrincipalPolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (ppb *PrincipalPolicyBuilder) Build() *policyv1.Policy { _ = "STUB: not implemented"; return nil }

// RoleRuleBuilder is a builder for resource rules.
type RoleRuleBuilder struct {
	rule *policyv1.RoleRule
}

func NewRoleRule(resource string, actions ...string) *RoleRuleBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rrb *RoleRuleBuilder) Build() *policyv1.RoleRule {
	_ = "STUB: not implemented"

	// RolePolicyBuilder is a builder for role policies.
	return nil
}

type RolePolicyBuilder struct {
	rp *policyv1.RolePolicy
}

func NewRolePolicyBuilder(role string) *RolePolicyBuilder { _ = "STUB: not implemented"; return nil }

func (rpb *RolePolicyBuilder) WithRules(rules ...*policyv1.RoleRule) *RolePolicyBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (rpb *RolePolicyBuilder) Build() *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func GenDisabledPrincipalPolicy(mod NameMod) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}

func GenPrincipalPolicy(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

type DerivedRolesBuilder struct {
	dr *policyv1.DerivedRoles
}

func NewDerivedRolesBuilder(name string) *DerivedRolesBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (drb *DerivedRolesBuilder) AddRole(name string, parentRoles ...string) *DerivedRolesBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (drb *DerivedRolesBuilder) AddRoleWithMatch(name string, parentRoles []string, expr ...string) *DerivedRolesBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (drb *DerivedRolesBuilder) addRoleDef(name string, parentRoles []string, comp *policyv1.Condition) *DerivedRolesBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (drb *DerivedRolesBuilder) Build() *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func GenDisabledDerivedRoles(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func GenDerivedRoles(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func GenDisabledExportConstants(mod NameMod) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}

func GenExportConstants(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func GenDisabledExportVariables(mod NameMod) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}

func GenExportVariables(mod NameMod) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func PrefixAndSuffix(prefix, suffix string) NameMod {
	_ = "STUB: not implemented"
	return *new(NameMod)
}

func Suffix(suffix string) NameMod { _ = "STUB: not implemented"; return *new(NameMod) }

func NoMod() NameMod { _ = "STUB: not implemented"; return *new(NameMod) }
