// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package namer

import (
	"iter"
	"regexp"
)

var (
	invalidIdentifierChars = regexp.MustCompile(`[^\w.]+`)
	// Naming pattern imposed on resource and principal names before Cerbos 0.30.0.
	oldNamePattern = regexp.MustCompile(`^[[:alpha:]][[:word:]\@\.\-/]*(\:[[:alpha:]][[:word:]\@\.\-/]*)*$`)
)

const (
	DerivedRolesPrefix      = fqnPrefix + "derived_roles"
	ExportConstantsPrefix   = fqnPrefix + "export_constants"
	ExportVariablesPrefix   = fqnPrefix + "export_variables"
	PrincipalPoliciesPrefix = fqnPrefix + "principal"
	ResourcePoliciesPrefix  = fqnPrefix + "resource"
	RolePoliciesPrefix      = fqnPrefix + "role"

	DefaultVersion = "default"
	DefaultScope   = ""
	fqnPrefix      = "cerbos."
)

// ModuleID is a unique identifier for modules.
type ModuleID struct {
	hash uint64
}

func (m *ModuleID) String() string { _ = "STUB: not implemented"; return "" }

func (m *ModuleID) HexStr() string { _ = "STUB: not implemented"; return "" }

func (m ModuleID) RawValue() uint64 {
	_ = "STUB: not implemented"

	// GenModuleIDFromFQN generates a short ID for the given module name.
	return 0
}

func GenModuleIDFromFQN(name string) ModuleID { _ = "STUB: not implemented"; return *new(ModuleID) }

func buildFQNTree[T any](fqn, scope string, elementFn func(string) T) []T {
	_ = "STUB: not implemented"
	return nil
}

// add the no scope FQN as the root

func ScopeParents(scope string) iter.Seq[string] { _ = "STUB: not implemented"; return nil }

func ScopeFromFQN(fqn string) string { _ = "STUB: not implemented"; return "" }

// PolicyKeyFromFQN returns a policy key from the module name.
func PolicyKeyFromFQN(m string) string { _ = "STUB: not implemented"; return "" }

// FQNFromPolicyKey returns FQN from the policy key.
func FQNFromPolicyKey(s string) string { _ = "STUB: not implemented"; return "" }

func SanitizedResource(resource string) string { _ = "STUB: not implemented"; return "" }

// ResourcePolicyFQN returns the fully-qualified name for the resource policy with given resource, version and scope.
func ResourcePolicyFQN(resource, version, scope string) string {
	_ = "STUB: not implemented"
	return ""
}

// ResourcePolicyModuleID returns the module ID for the resource policy with given resource, version and scope.
func ResourcePolicyModuleID(resource, version, scope string) ModuleID {
	_ = "STUB: not implemented"
	return *new(ModuleID)
}

// ScopedResourcePolicyModuleIDs returns a list of module IDs for each scope segment if `genTree` is true.
// For example, if the scope is `a.b.c`, the list will contain the module IDs for scopes `a.b.c`, `a.b`, `a` and `""` in that order.
func ScopedResourcePolicyModuleIDs(resource, version, scope string, genTree bool) []ModuleID {
	_ = "STUB: not implemented"
	return nil
}

// PrincipalPolicyFQN returns the fully-qualified module name for the principal policy with given principal, version and scope.
func PrincipalPolicyFQN(principal, version, scope string) string {
	_ = "STUB: not implemented"
	return ""
}

// PrincipalPolicyModuleID returns the module ID for the principal policy with given principal and version.
func PrincipalPolicyModuleID(principal, version, scope string) ModuleID {
	_ = "STUB: not implemented"
	return *new(ModuleID)
}

// RolePolicyFQN returns the fully-qualified module name for the role policies with the given scope.
// If version is empty, it defaults to DefaultVersion.
func RolePolicyFQN(role, version, scope string) string { _ = "STUB: not implemented"; return "" }

// RolePolicyModuleID returns the module ID for the role policies with the given scope.
func RolePolicyModuleID(role, version, scope string) ModuleID {
	_ = "STUB: not implemented"
	return *new(ModuleID)
}

// ScopedPrincipalPolicyModuleIDs returns a list of module IDs for each scope segment if `strict` is false.
// For example, if the scope is `a.b.c`, the list will contain the module IDs for scopes `a.b.c`, `a.b`, `a` and `""` in that order.
func ScopedPrincipalPolicyModuleIDs(principal, version, scope string, genTree bool) []ModuleID {
	_ = "STUB: not implemented"
	return nil
}

// DerivedRolesFQN returns the fully-qualified module name for the given derived roles set.
func DerivedRolesFQN(roleSetName string) string { _ = "STUB: not implemented"; return "" }

// DerivedRolesModuleID returns the module ID for the given derived roles set.
func DerivedRolesModuleID(roleSetName string) ModuleID {
	_ = "STUB: not implemented"
	return *new(ModuleID)
}

// ExportConstantsFQN returns the fully-qualified module name for the given exported constant definitions.
func ExportConstantsFQN(constantsName string) string { _ = "STUB: not implemented"; return "" }

// ExportConstantsModuleID returns the module ID for the given exported constant definitions.
func ExportConstantsModuleID(constantsName string) ModuleID {
	_ = "STUB: not implemented"
	return *new(ModuleID)
}

// ExportVariablesFQN returns the fully-qualified module name for the given exported variable definitions.
func ExportVariablesFQN(variablesName string) string { _ = "STUB: not implemented"; return "" }

// ExportVariablesModuleID returns the module ID for the given exported variable definitions.
func ExportVariablesModuleID(variablesName string) ModuleID {
	_ = "STUB: not implemented"
	return *new(ModuleID)
}

// SimpleName extracts the simple name from a derived roles, exported constants, or exported variables FQN.
func SimpleName(fqn string) string { _ = "STUB: not implemented"; return "" }

func withScope(fqn, scope string) string { _ = "STUB: not implemented"; return "" }

// sanitize replaces special characters in the string with underscores.
// Before Cerbos 0.30 the names of resources or principals had to follow a certain pattern. We then replaced some of
// the non-word characters with underscores because earlier versions of Cerbos used to generate Rego code for policies.
// Because we used the sanitized name for computing the module ID of the policy, in order to maintain backward compatibility
// and not break database stores we still have to do the same if the name matches the pattern.
func sanitize(v string) string { _ = "STUB: not implemented"; return "" }

// RuleFQN returns the FQN for the resource rule or principal resource action rule with scope granularity.
func RuleFQN(rpsMeta any, scope, ruleName string) string { _ = "STUB: not implemented"; return "" }

type PolicyCoords struct {
	Kind    string
	Name    string
	Version string
	Scope   string
}

func (pc PolicyCoords) FQN() string { _ = "STUB: not implemented"; return "" }

func (pc PolicyCoords) PolicyKey() string { _ = "STUB: not implemented"; return "" }

func ScopeValue(scope string) string { _ = "STUB: not implemented"; return "" }

type Policy struct {
	PolicyCoords
	ID ModuleID
}
