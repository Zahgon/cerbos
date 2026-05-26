// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package namer

import (
	"database/sql/driver"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

func (m *ModuleID) Scan(src any) error { _ = "STUB: not implemented"; return nil }

// hack to work around unpredictable behaviour from the MySQL driver (it's a feature, not a bug).
// https://github.com/go-sql-driver/mysql/issues/861

func (m ModuleID) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *

	// GenModuleID generates a short ID for the module.
	new(driver.Value), nil
}

func GenModuleID(p *policyv1.Policy) ModuleID { _ = "STUB: not implemented"; return *new(ModuleID) }

// FQN returns the fully-qualified name of the policy.
func FQN(p *policyv1.Policy) string { _ = "STUB: not implemented"; return "" }

// FQNTree returns the tree of FQNs that are ancestors of the given policy (including itself) sorted by most recent to oldest.
// For example, if the policy has scope a.b.c, the returned tree will contain the FQNs in the following order:
// - a.b.c
// - a.b
// - a
// - "" (empty scope).
func FQNTree(p *policyv1.Policy) []string { _ = "STUB: not implemented"; return nil }

// role policies don't functionally have ancestors

// PolicyKey returns a human-friendly identifier that can be used to refer to the policy in logs and other outputs.
func PolicyKey(p *policyv1.Policy) string { _ = "STUB: not implemented"; return "" }

// ResourceRuleName returns the name of the given resource rule.
func ResourceRuleName(rule *policyv1.ResourceRule, idx int) string {
	_ = "STUB: not implemented"
	return ""
}

// PrincipalResourceActionRuleName returns the name for an action rule defined for a particular resource.
func PrincipalResourceActionRuleName(rule *policyv1.PrincipalRule_Action, resource string, idx int) string {
	_ = "STUB: not implemented"
	return ""
}
