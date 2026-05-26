// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package index

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/google/cel-go/cel"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// functionalRuleRowFields lists proto field names that affect evaluation outcome.
// Routing fields (scope, version, resource, role, action, principal) are excluded
// because they are handled by the bitmap index dimensions.
var functionalRuleRowFields = map[protoreflect.Name]struct{}{
	"condition": {}, "derived_role_condition": {},
	"effect": {}, "scope_permissions": {},
	"emit_output": {}, "params": {},
	"derived_role_params": {}, "policy_kind": {},
	"from_role_policy": {},
}

var nonFunctionalChecksumFields = func() map[string]struct{} {
	res := make(map[string]struct{})
	desc := (&runtimev1.RuleTable_RuleRow{}).ProtoReflect().Descriptor()
	fields := desc.Fields()
	for i := range fields.Len() {
		f := fields.Get(i)
		if _, ok := functionalRuleRowFields[f.Name()]; !ok {
			res[string(f.FullName())] = struct{}{}
		}
	}
	res["cerbos.runtime.v1.RuleTableMetadata.source_attributes"] = struct{}{}
	return res
}()

type CelProgram struct {
	Prog cel.Program
	Name string
	Expr string
}

type Option func(*Index)

type Index struct {
	bi          *bitmapIndex
	parentRoles map[string]map[string][]string
}

func New(opts ...Option) *Index { _ = "STUB: not implemented"; return nil }

func (m *Index) IndexRules(rules []*runtimev1.RuleTable_RuleRow) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:exhaustive

// hashpb does not include field tags, so Condition=X/DerivedRoleCondition=nil
// hashes identically to Condition=nil/DerivedRoleCondition=X when X is the
// same Condition content. Feed a discriminator byte into the hasher to break
// the collision.
// TODO(saml): addressing upstream in protoc-gen-go-hashpb, remove this when that lands.

func (m *Index) GetAllRows() []*Binding { _ = "STUB: not implemented"; return nil }

// Query returns bindings matching the given dimensions. Nil or zero-values mean
// "match all" for that dimension. Synthetic DENYs from role policy AllowActions
// are prepended when querying for a specific action with KIND_RESOURCE.
func (m *Index) Query(version, resource, scope, action string, roles []string, policyKind policyv1.Kind, principalID string, buf []*Binding) []*Binding {
	_ = "STUB: not implemented"
	return nil
}

// scope is always filtered because "" is a valid literal scope (root scope).

//nolint:mnd

// baseBM = AND of all non-action dimensions.

//nolint:mnd

// resultBM = AND(baseBM, actionBM) for regular (non-AllowActions) bindings.

// Role policy synthetic DENYs are prepended so the evaluator sees them
// before regular ALLOWs, which is required for scope permission semantics.

// Regular bindings.

// appendRolePolicyDenies appends synthetic DENY bindings for each (role,
// resource, action) triple where the role has a role policy matching
// versionBM ∩ scopeBM ∩ roleBM but doesn't explicitly allow that action
// for that resource.
//
// When roles is empty, iterates every role whose policy matches the
// non-resource dims (in binding-discovery order). Returns res unchanged if
// resources is empty.
//
// When targetActions is empty, the action set for synthesis is derived
// per-resource from the bindings (excluding principal-policy bindings) for
// that resource intersected with versionBM ∩ scopeBM. The role filter is
// deliberately ignored when deriving actions, so the action set reflects the
// resource's full surface, not just actions referenced by filtered roles.
func (m *Index) appendRolePolicyDenies(
	arena *bitmapArena, bi *bitmapIndex,
	resources, roles, targetActions []string,
	versionBM, scopeBM, roleBM *Bitmap,
	res []*Binding,
) []*Binding {
	_ = "STUB: not implemented"
	// candidateBM deliberately omits resource so we can spot roles whose
	// policy exists but doesn't cover the requested resource.
	return nil
}

//nolint:mnd

// Retrieve one sample binding per role for version/scope

// Relevant in some (external) downstream consumers

// role policy exists, but no resource bindings present

// role policy exists with resource bindings, but action not specified

// Pure ACL allow: fall through. Role-policy bindings are
// otherwise dropped here, so emit any output via a no-effect
// binding.

// Synthetic DENY for the negated condition. Outputs are swapped
// because synthetic-activated == user-condition-not-met.

// collectResourceActions returns the action names referenced by bindings on
// the given resource (intersected with version/scope). Principal-policy
// actions are excluded — they're principal-specific and shouldn't widen the
// resource's action set.
func collectResourceActions(arena *bitmapArena, bi *bitmapIndex, resBM, versionBM, scopeBM *Bitmap) []string {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

func newNoMatchRolePolicyDeny(role, version, scope, resource, action string) *Binding {
	_ = "STUB: not implemented"
	return nil
}

// QueryMulti returns bindings matching across multiple values per dimension
// (OR within each dimension, AND across dimensions). When withRolePolicyDenies
// is true, synthetic DENY bindings are appended for each (role, resource,
// action) triple where the role has a matching role policy that doesn't
// explicitly allow the action. An empty actions slice expands to every action
// in the index for synthesis.
func (m *Index) QueryMulti(versions, resources, scopes, roles, actions []string, withRolePolicyDenies bool) []*Binding {
	_ = "STUB: not implemented"
	return nil
}

// An empty resourceBM doesn't short-circuit: role-policy synthesis
// still emits NoMatch denies when a role has a policy in the other
// dimensions but no rows for the requested resource.

//nolint:mnd

func (m *Index) applyActionFilter(arena *bitmapArena, baseBM *Bitmap, actions []string) *Bitmap {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

// AddParentRoles returns the given roles plus the union of all their parent roles across the provided scopes.
// When multiple scopes define parents for the same role, the results are merged rather than overwritten.
func (m *Index) AddParentRoles(scopes, roles []string) []string {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

func (m *Index) IndexParentRoles(scopeParentRoles map[string]*runtimev1.RuleTable_RoleParentRoles) error {
	_ = "STUB: not implemented"
	return nil
}

func compileParentRoleAncestors(scopeParentRoles map[string]*runtimev1.RuleTable_RoleParentRoles) map[string]map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func collectParentRoles(scopeParentRoles map[string]*runtimev1.RuleTable_RoleParentRoles, scope, role string, parentRoleSet, visited map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func (m *Index) DeletePolicy(fqn string) error { _ = "STUB: not implemented"; return nil }

func (m *Index) GetScopes() []string { _ = "STUB: not implemented"; return nil }

func (m *Index) GetRoles() []string { _ = "STUB: not implemented"; return nil }

func (m *Index) GetVersions() []string { _ = "STUB: not implemented"; return nil }

func (m *Index) GetActions() []string { _ = "STUB: not implemented"; return nil }

func (m *Index) GetResources() []string { _ = "STUB: not implemented"; return nil }

// Roles returns role keys matching the filters. An empty filter means
// match-all on that dimension. Glob keys (e.g. "manager:*") appear verbatim.
// Principal-policy rows are excluded.
func (m *Index) Roles(versions, scopes []string) []string { _ = "STUB: not implemented"; return nil }

// Resources returns resource keys matching the filters.
// Matches `Roles` semantics above.
func (m *Index) Resources(versions, scopes []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// ActionsForResource returns distinct actions on resource, filtered by
// versions and scopes. Empty filter = match-all; empty resource = nil.
//
// resource is matched as at evaluation: literal lookup plus any matching
// glob patterns. Principal-policy rows are excluded.
func (m *Index) ActionsForResource(resource string, versions, scopes []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// nonPrincipalDimensionKeys returns keys of gd matching the filters,
// excluding principal-policy bindings. Returns nil if no KIND_RESOURCE
// bindings are indexed.
func (bi *bitmapIndex) nonPrincipalDimensionKeys(gd *globDimension, versions, scopes []string) []string {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

// versionScopeFilters builds the filter bitmaps. ok=false means a non-empty
// input matched nothing. A nil bitmap means "no filter" for that dimension.
func (bi *bitmapIndex) versionScopeFilters(arena *bitmapArena, versions, scopes []string) (versionBM, scopeBM *Bitmap, ok bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

func (m *Index) ScopedResourceExists(version, resource string, scopes []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Index) ScopedPrincipalExists(version string, scopes []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Index) Reset() { _ = "STUB: not implemented"; return }

// getOrGenerateParams returns cached RowParams for the given proto content hash,
// compiling CEL programs on miss. Rows with identical params share the same pointer and Key.
func getOrGenerateParams(cache map[uint64]*RowParams, proto *runtimev1.RuleTable_RuleRow_Params) (*RowParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCelProgramsFromExpressions(vars []*runtimev1.Variable) ([]*CelProgram, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
