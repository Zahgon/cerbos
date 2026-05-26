// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"time"

	"google.golang.org/protobuf/types/known/structpb"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	sourcev1 "github.com/cerbos/cerbos/api/genpb/cerbos/source/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/parser"
)

// Kind defines the type of policy (resource, principal, derived_roles etc.).
type Kind policyv1.Kind

const (
	DerivedRolesKind    Kind = Kind(policyv1.Kind_KIND_DERIVED_ROLES)
	ExportConstantsKind Kind = Kind(policyv1.Kind_KIND_EXPORT_CONSTANTS)
	ExportVariablesKind Kind = Kind(policyv1.Kind_KIND_EXPORT_VARIABLES)
	PrincipalKind       Kind = Kind(policyv1.Kind_KIND_PRINCIPAL)
	ResourceKind        Kind = Kind(policyv1.Kind_KIND_RESOURCE)
	RolePolicyKind      Kind = Kind(policyv1.Kind_KIND_ROLE_POLICY)
)

const (
	ResourceKindStr        = "RESOURCE"
	PrincipalKindStr       = "PRINCIPAL"
	DerivedRolesKindStr    = "DERIVED_ROLES"
	ExportConstantsKindStr = "EXPORT_CONSTANTS"
	ExportVariablesKindStr = "EXPORT_VARIABLES"
	RolePolicyKindStr      = "ROLE"
)

var IgnoreHashFields = map[string]struct{}{
	"cerbos.policy.v1.Policy.description": {},
	"cerbos.policy.v1.Policy.disabled":    {},
	"cerbos.policy.v1.Policy.json_schema": {},
	"cerbos.policy.v1.Policy.metadata":    {},
}

func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

// GetKind returns the kind of the given policy.
func GetKind(p *policyv1.Policy) Kind { _ = "STUB: not implemented"; return *new(Kind) }

// KindFromFQN returns the kind of policy referred to by the given fully-qualified name.
func KindFromFQN(fqn string) Kind { _ = "STUB: not implemented"; return *new(Kind) }

// Dependencies returns the module names of dependencies of the policy.
func Dependencies(p *policyv1.Policy) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ancestors returns the module IDs of the ancestors of this policy from most recent to oldest.
func Ancestors(p *policyv1.Policy) []namer.ModuleID { _ = "STUB: not implemented"; return nil }

// first element is the policy itself so we ignore that

// RequiredAncestors returns the moduleID to FQN mapping of required ancestors of the policy.
func RequiredAncestors(p *policyv1.Policy) map[namer.ModuleID]string {
	_ = "STUB: not implemented"
	return nil
}

// first element is the policy itself so we ignore that

// SchemaReferences returns references to the schemas found in the policy.
func SchemaReferences(p *policyv1.Policy) []string { _ = "STUB: not implemented"; return nil }

func GetScope(p *policyv1.Policy) string { _ = "STUB: not implemented"; return "" }

// SourceAttribute holds structured information about the policy from its source.
type SourceAttribute struct {
	Value *structpb.Value
	Key   string
}

// SourceDriver creates a source attribute for the storage driver.
func SourceDriver(driver string) SourceAttribute {
	_ = "STUB: not implemented"
	return *new(SourceAttribute)
}

// SourceFile creates a source attribute describing the file name of the policy.
func SourceFile(source string) SourceAttribute {
	_ = "STUB: not implemented"
	return *new(SourceAttribute)
}

// SourceUpdateTS creates a source attribute describing the time a policy was updated in a mutable store.
func SourceUpdateTS(timestamp time.Time) SourceAttribute {
	_ = "STUB: not implemented"
	return *new(SourceAttribute)
}

// SourceUpdateTSNow creates a source attribute setting the update time to now.
func SourceUpdateTSNow() SourceAttribute { _ = "STUB: not implemented"; return *new(SourceAttribute) }

// WithSourceAttributes adds given source attributes to the policy.
func WithSourceAttributes(p *policyv1.Policy, attrs ...SourceAttribute) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}

// WithMetadata adds metadata to the policy.
func WithMetadata(p *policyv1.Policy, source string, annotations map[string]string, storeIdentifier string, sourceAttr ...SourceAttribute) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}

func mergeAnnotations(a, b map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// WithStoreIdentifier adds the store identifier to the metadata.
func WithStoreIdentifier(p *policyv1.Policy, storeIdentifier string) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck
// TODO: Remove this after deprecated StoreIdentifer no longer exists

// WithHash calculates the hash for the policy and adds it to metadata.
func WithHash(p *policyv1.Policy) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

// GetHash returns the hash of the policy.
func GetHash(p *policyv1.Policy) uint64 { _ = "STUB: not implemented"; return 0 }

// GetSourceFile gets the source file name from metadata if it exists.
func GetSourceFile(p *policyv1.Policy) string { _ = "STUB: not implemented"; return "" }

// Wrapper is a convenience layer over the policy definition.
type Wrapper struct {
	*policyv1.Policy
	FQN     string
	Name    string
	Version string
	Scope   string
	ID      namer.ModuleID
	Kind    Kind
}

// ListActions returns unique list of actions in a policy.
func ListActions(p *policyv1.Policy) []string { _ = "STUB: not implemented"; return nil }

// ListExportedDerivedRoles returns exported derived roles defined in the given derived roles policy.
func ListExportedDerivedRoles(drp *policyv1.DerivedRoles) []*responsev1.InspectPoliciesResponse_DerivedRole {
	_ = "STUB: not implemented"
	return nil
}

// ListConstants returns local and exported constants (not imported ones) defined in a policy.
func ListConstants(p *policyv1.Policy) map[string]*responsev1.InspectPoliciesResponse_Constant {
	_ = "STUB: not implemented" //nolint:dupl
	return nil
}

// ListVariables returns local and exported variables (not imported ones) defined in a policy.
func ListVariables(p *policyv1.Policy) map[string]*responsev1.InspectPoliciesResponse_Variable {
	_ = "STUB: not implemented" //nolint:dupl
	return nil
}

// ListPolicySetActions returns unique list of actions in a policy set.
func ListPolicySetActions(ps *runtimev1.RunnablePolicySet) []string {
	_ = "STUB: not implemented"
	return nil
}

// ListPolicySetDerivedRoles returns imported and used derived roles defined in a policy set.
func ListPolicySetDerivedRoles(ps *runtimev1.RunnablePolicySet) []*responsev1.InspectPoliciesResponse_DerivedRole {
	_ = "STUB: not implemented"
	return nil
}

// ListPolicySetConstants returns local and exported variables defined in a policy set.
func ListPolicySetConstants(ps *runtimev1.RunnablePolicySet) []*responsev1.InspectPoliciesResponse_Constant {
	_ = "STUB: not implemented"
	return nil
}

// ListPolicySetVariables returns local and exported variables defined in a policy set.
func ListPolicySetVariables(ps *runtimev1.RunnablePolicySet) []*responsev1.InspectPoliciesResponse_Variable {
	_ = "STUB: not implemented"
	return nil
}

// Wrap augments a policy with useful information about itself.
func Wrap(p *policyv1.Policy) Wrapper { _ = "STUB: not implemented"; return *new(Wrapper) }

func (w Wrapper) Dependencies() []namer.ModuleID { _ = "STUB: not implemented"; return nil }

func (w Wrapper) ToProto() *sourcev1.PolicyWrapper { _ = "STUB: not implemented"; return nil }

// CompilationUnit is the set of policies that need to be compiled together.
// For example, if a resource policy named R imports derived roles named D, the compilation unit will contain
// both R and D with the ModID field pointing to R because it is the main policy.
type CompilationUnit struct {
	Definitions    map[namer.ModuleID]*policyv1.Policy
	SourceContexts map[namer.ModuleID]parser.SourceCtx
	ModID          namer.ModuleID
}

func (cu *CompilationUnit) AddDefinition(id namer.ModuleID, p *policyv1.Policy, sc parser.SourceCtx) {
	_ = "STUB: not implemented"
	return
}

func (cu *CompilationUnit) MainSourceFile() string { _ = "STUB: not implemented"; return "" }

func (cu *CompilationUnit) MainPolicy() *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func (cu *CompilationUnit) Ancestors() []namer.ModuleID { _ = "STUB: not implemented"; return nil }

// Key returns the human readable identifier for the main module.
func (cu *CompilationUnit) Key() string { _ = "STUB: not implemented"; return "" }
