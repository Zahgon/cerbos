// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

func ResourcePolicyRuleProtoPath(idx int) string { _ = "STUB: not implemented"; return "" }

func ResourcePolicyRuleReferencedDerivedRoleProtoPath(ruleIdx, roleIdx int) string {
	_ = "STUB: not implemented"
	return ""
}

func ResourcePolicyImportDerivedRolesProtoPath(idx int) string {
	_ = "STUB: not implemented"
	return ""
}

func ResourcePolicyPrincipalSchemaProtoPath() string { _ = "STUB: not implemented"; return "" }

func ResourcePolicyResourceSchemaProtoPath() string { _ = "STUB: not implemented"; return "" }

func PrincipalPolicyRuleProtoPath(idx int) string { _ = "STUB: not implemented"; return "" }

func PrincipalPolicyActionRuleProtoPath(parentIdx, idx int) string {
	_ = "STUB: not implemented"
	return ""
}

func RolePolicyRuleProtoPath(idx int) string { _ = "STUB: not implemented"; return "" }

func RolePolicyConditionProtoPath(idx int) string { _ = "STUB: not implemented"; return "" }

func DerivedRoleConditionProtoPath(idx int) string { _ = "STUB: not implemented"; return "" }

func DerivedRoleRuleProtoPath(idx int) string { _ = "STUB: not implemented"; return "" }

func ExportConstantsConstantProtoPath() string { _ = "STUB: not implemented"; return "" }

func ConstantsImportProtoPath(p *policyv1.Policy, idx int) string {
	_ = "STUB: not implemented"
	return ""
}

func ConstantsLocalProtoPath(p *policyv1.Policy) string { _ = "STUB: not implemented"; return "" }

func ExportVariablesVariableProtoPath() string { _ = "STUB: not implemented"; return "" }

func VariablesImportProtoPath(p *policyv1.Policy, idx int) string {
	_ = "STUB: not implemented"
	return ""
}

func VariablesLocalProtoPath(p *policyv1.Policy) string { _ = "STUB: not implemented"; return "" }

func policyKind(p *policyv1.Policy) string { _ = "STUB: not implemented"; return "" }
