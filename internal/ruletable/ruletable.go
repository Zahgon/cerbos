// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package ruletable

import (
	"context"
	"sync"

	"github.com/google/cel-go/cel"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/engine/policyloader"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/ruletable/index"
	"github.com/cerbos/cerbos/internal/schema"
)

type compilerVersionMigration func(*runtimev1.RuleTable) error

var (
	compilerVersionMigrations = []compilerVersionMigration{
		migrateFromCompilerVersion0To1,
	}

	compilerVersion = uint32(len(compilerVersionMigrations))
)

const (
	conditionNotSatisfied   = "Condition not satisfied"
	noMatchScopePermissions = "NO_MATCH_FOR_SCOPE_PERMISSIONS"
	noPolicyMatch           = "NO_MATCH"
)

func NewProtoRuletable() *runtimev1.RuleTable { _ = "STUB: not implemented"; return nil }

func LoadPolicies(ctx context.Context, rt *runtimev1.RuleTable, pl policyloader.PolicyLoader) error {
	_ = "STUB: not implemented"
	return nil
}

func LoadSchemas(ctx context.Context, rt *runtimev1.RuleTable, sl schema.Loader) error {
	_ = "STUB: not implemented"
	return nil
}

func AddPolicy(rt *runtimev1.RuleTable, rps *runtimev1.RunnablePolicySet) []*runtimev1.RuleTable_RuleRow {
	_ = "STUB: not implemented"
	return nil
}

func addPrincipalPolicy(rt *runtimev1.RuleTable, rpps *runtimev1.RunnablePrincipalPolicySet) (res []*runtimev1.RuleTable_RuleRow) {
	_ = "STUB: not implemented"
	return nil
}

// We only process the first of principal policy sets as it's assumed parent scopes are handled in separate calls

// Cover the edge case where a policy is created with no rules (useful in the multitenant case
// where a tenant might want to "inherit" the permissions of the parent by immediately falling
// through). We do this by creating a noop row in the rule table which means we bypass the
// "policy does not exist in the scope" during evaluation.

//nolint:staticcheck

//nolint:staticcheck

// Since principal policies don't have explicit roles, we use "*" to match any role

func addResourcePolicy(rt *runtimev1.RuleTable, rrps *runtimev1.RunnableResourcePolicySet) (res []*runtimev1.RuleTable_RuleRow) {
	_ = "STUB: not implemented"
	return nil
}

// we only process the first of resource policy sets as it's assumed parent scopes are handled in separate calls

// Cover the edge case where a policy is created with no rules (useful in the multitenant case
// where a tenant might want to "inherit" the permissions of the parent by immediately falling
// through). We do this by creating a noop row in the rule table which means we bypass the
// "policy does not exist in the scope" during evaluation.

//nolint:staticcheck

//nolint:staticcheck

// merge derived roles as roles with added conditions

func addRolePolicy(rt *runtimev1.RuleTable, p *runtimev1.RunnableRolePolicySet) (res []*runtimev1.RuleTable_RuleRow) {
	_ = "STUB: not implemented"
	return nil
}

func buildRawSchemas(ctx context.Context, rt *runtimev1.RuleTable, resolver schema.Resolver) error {
	_ = "STUB: not implemented"
	return nil
}

type RuleTable struct {
	*runtimev1.RuleTable
	idx                   *index.Index
	principalScopeMap     map[string]struct{}
	resourceScopeMap      map[string]struct{}
	scopeScopePermissions map[string]policyv1.ScopePermissions
	policyDerivedRoles    map[namer.ModuleID]map[string]*WrappedRunnableDerivedRole
	programCache          *ProgramCache
}

type WrappedRunnableDerivedRole struct {
	*runtimev1.RunnableDerivedRole
	Constants   map[string]any
	VarCacheKey uint64
}

// ProgramCache caches compiled CEL programs keyed by CheckedExpr pointer to avoid repeated compilation.
// Programs are compiled with CacheFriendlyTimeDecorator which looks up NowFunc from activation at eval time.
type ProgramCache struct {
	m  map[*exprpb.CheckedExpr]cel.Program
	mu sync.RWMutex
}

func NewProgramCache() *ProgramCache { _ = "STUB: not implemented"; return nil }

func (c *ProgramCache) Clear() { _ = "STUB: not implemented"; return }

func (c *ProgramCache) GetOrCreate(expr *exprpb.CheckedExpr) (cel.Program, error) {
	_ = "STUB: not implemented"
	return *new(cel.Program), nil
}

func NewRuleTableFromLoader(ctx context.Context, policyLoader policyloader.PolicyLoader) (*RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRuleTable(protoRT *runtimev1.RuleTable) (*RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rt *RuleTable) init(protoRT *runtimev1.RuleTable) error {
	rt.RuleTable = protoRT

	if err := migrate(rt.RuleTable); err != nil {
		return err
	}

	// clear maps prior to creating new ones to reduce memory pressure in reload scenarios
	clear(rt.policyDerivedRoles)
	clear(rt.principalScopeMap)
	clear(rt.resourceScopeMap)
	clear(rt.scopeScopePermissions)
	rt.programCache.Clear()

	rt.idx.Reset()
	rt.policyDerivedRoles = make(map[namer.ModuleID]map[string]*WrappedRunnableDerivedRole)
	rt.principalScopeMap = make(map[string]struct{})
	rt.resourceScopeMap = make(map[string]struct{})
	rt.scopeScopePermissions = make(map[string]policyv1.ScopePermissions)

	if err := rt.indexRules(rt.Rules); err != nil {
		return err
	}

	// rules are now indexed, we can clear up any unnecessary transport state
	clear(rt.Rules)
	rt.Rules = []*runtimev1.RuleTable_RuleRow{} // otherwise the empty slice hangs around

	clear(rt.PolicyDerivedRoles)

	return nil
}

func (rt *RuleTable) indexRules(rules []*runtimev1.RuleTable_RuleRow) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:exhaustive

func (rt *RuleTable) GetAllRows() []*index.Binding { _ = "STUB: not implemented"; return nil }

func (rt *RuleTable) GetDerivedRoles(fqn string) map[string]*WrappedRunnableDerivedRole {
	_ = "STUB: not implemented"
	return nil
}

func (rt *RuleTable) GetAllScopes(pt policyv1.Kind, scope, name, version string, lenient bool) ([]string, string, string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

//nolint:exhaustive

type scopeNode struct {
	children map[string]*scopeNode
	scope    string
}

func (rt *RuleTable) CombineScopes(principalScopes, resourceScopes []string) []string {
	_ = "STUB: not implemented"
	// Build a map to track all unique scopes
	return nil
}

// Use DFS to traverse the tree with children first, then parents

func (rt *RuleTable) GetScopeScopePermissions(scope string) policyv1.ScopePermissions {
	_ = "STUB: not implemented"
	return *new(policyv1.ScopePermissions)
}

func (rt *RuleTable) GetSchema(fqn string) *policyv1.Schemas { _ = "STUB: not implemented"; return nil }

func (rt *RuleTable) GetMeta(fqn string) *runtimev1.RuleTableMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (rt *RuleTable) Evaluator(evalConf *evaluator.Conf, schemaConf *schema.Conf) (evaluator.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(evaluator.Evaluator), nil
}

func migrate(rt *runtimev1.RuleTable) error { _ = "STUB: not implemented"; return nil }

func migrateFromCompilerVersion0To1(rt *runtimev1.RuleTable) error {
	_ = "STUB: not implemented"
	return nil
}
