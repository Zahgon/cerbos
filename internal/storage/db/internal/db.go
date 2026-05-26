// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package internal

import (
	"context"
	"errors"
	"io"
	"regexp"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"

	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
)

const (
	driverName  = "db"
	tableLogKey = "table"
)

var errUpsertPolicyRequired = errors.New("invalid driver configuration: upsertPolicy is required")

type DBStorage interface {
	storage.Subscribable
	storage.Instrumented
	storage.Reloadable
	storage.Verifiable
	AddOrUpdate(ctx context.Context, policies ...policy.Wrapper) error
	GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (*policy.CompilationUnit, error)
	GetAll(ctx context.Context) ([]*policy.CompilationUnit, error)
	GetAllMatching(ctx context.Context, modIDs []namer.ModuleID) ([]*policy.CompilationUnit, error)
	GetCompilationUnits(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error)
	GetDependents(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error)
	GetDescendants(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.Policy, error)
	InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error)
	ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) ([]string, error)
	ListSchemaIDs(ctx context.Context) ([]string, error)
	AddOrUpdateSchema(ctx context.Context, schemas ...*schemav1.Schema) error
	Delete(ctx context.Context, policyKey ...string) (uint32, error)
	Disable(ctx context.Context, policyKey ...string) (uint32, error)
	Enable(ctx context.Context, policyKey ...string) (uint32, error)
	DeleteSchema(ctx context.Context, ids ...string) (uint32, error)
	LoadSchema(ctx context.Context, url string) (io.ReadCloser, error)
	LoadPolicy(ctx context.Context, policyKey ...string) ([]*policy.Wrapper, error)
	ListRevisions(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID]int, error)
	PurgeRevisions(ctx context.Context, keepLast uint32) (uint32, error)
}

func NewDBStorage(ctx context.Context, db *goqu.Database, dbOpts ...DBOpt) (DBStorage, error) {
	_ = "STUB: not implemented"
	return *new(DBStorage), nil
}

type dbStorage struct {
	opts *dbOpt
	db   *goqu.Database
	*storage.SubscriptionManager
}

func (s *dbStorage) AddOrUpdateSchema(ctx context.Context, schemas ...*schemav1.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *dbStorage) DeleteSchema(ctx context.Context, ids ...string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *dbStorage) LoadPolicy(ctx context.Context, policyKey ...string) ([]*policy.Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) LoadSchema(ctx context.Context, urlVar string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (s *dbStorage) AddOrUpdate(ctx context.Context, policies ...policy.Wrapper) error {
	_ = "STUB: not implemented"
	return nil
}

// We only need to retrieve the state of Dependents before the update (there's no need
// to update a dependent policy in the rule table if it's only just been added in the
// same batch). Therefore, we can share the same transaction.

// delete the existing dependency records

// insert the new dependency records

// delete the existing ancestor records

// insert the new ancestry records

// build a deduplicated union of dependents across the whole batch and
// exclude policies updated in this batch.

func (s *dbStorage) GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) GetAll(ctx context.Context) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) GetAllMatching(ctx context.Context, modIDs []namer.ModuleID) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) GetCompilationUnits(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Rather than writing a proper recursive query (which is pretty much impossible to do in a database-agnostic way), we're
// exploiting the fact that we have a maximum of two levels of dependency (resourcePolicy -> derivedRoles -> exportConstants/Variables).

type getCompilationUnitsQueryBuilder struct {
	query *goqu.SelectDataset
	depth int
}

// newGetCompilationUnitsQueryBuilder starts a query for retrieving policies and their ancestors and dependencies.
//
// The query starts out as
//
//	FROM policy AS p0
//	WHERE p0.id IN (?) AND p0.disabled = false
//
// JOIN clauses are added for ancestors using JoinAncestors and JoinDependencies, then finally a SELECT clause is
// added using Select.
func (s *dbStorage) newGetCompilationUnitsQueryBuilder(ids []namer.ModuleID) getCompilationUnitsQueryBuilder {
	_ = "STUB: not implemented"
	return *new(getCompilationUnitsQueryBuilder)
}

// JoinAncestors appends JOIN clauses to find ancestors of the policies.
func (q getCompilationUnitsQueryBuilder) JoinAncestors() getCompilationUnitsQueryBuilder {
	_ = "STUB: not implemented"
	return *new(getCompilationUnitsQueryBuilder)
}

// JoinDependencies appends JOIN clauses to find dependencies of the policies.
// It can be chained after JoinAncestors to get dependencies of ancestors, and
// after JoinDependencies to get transitive dependencies.
func (q getCompilationUnitsQueryBuilder) JoinDependencies() getCompilationUnitsQueryBuilder {
	_ = "STUB: not implemented"
	return *new(getCompilationUnitsQueryBuilder)
}

// join appends JOIN clauses for the given join table at the current depth N, producing a new query
// with depth N+1.
//
//	JOIN policy_dependency AS jN_N+1 ON pN.id = jN_N+1.policy_id
//	JOIN policy AS pN+1 ON (pN+1.id = jN_N+1.dependency_id AND pN+1.disabled = false)
func (q getCompilationUnitsQueryBuilder) join(joinTbl, joinTblParentIDCol, joinTblChildIDCol string) getCompilationUnitsQueryBuilder {
	_ = "STUB: not implemented"
	return *new(getCompilationUnitsQueryBuilder)
}

// Select produces the finished query for the current depth, N, by appending a SELECT clause.
//
//	SELECT p0.id AS unit_id, pN.id, pN.definition
func (q getCompilationUnitsQueryBuilder) Select() *goqu.SelectDataset {
	_ = "STUB: not implemented"
	return nil
}

// p is the policy table alias at depth N: pN.
func (q getCompilationUnitsQueryBuilder) p(depth int) exp.IdentifierExpression {
	_ = "STUB: not implemented"
	return *new(exp.IdentifierExpression)
}

// j is the join table (policy ancestor or dependency) alias between depth N and N+1: jN_N+1.
func (q getCompilationUnitsQueryBuilder) j(depth int) exp.IdentifierExpression {
	_ = "STUB: not implemented"
	return *new(exp.IdentifierExpression)
}

func (s *dbStorage) GetDependents(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) getDependents(ctx context.Context, tx *goqu.TxDatabase, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error) {
	_ = "STUB: not implemented"
	// Rather than writing a proper recursive query (which is pretty much impossible to do in a database-agnostic way), we're
	// exploiting the fact that we have a maximum of two levels of dependency (resourcePolicy -> derivedRoles -> exportVariables).
	return nil, nil
}

// SELECT dependency_id AS policy_id, policy_id AS dependent_id
// FROM policy_dependency
// WHERE policy_dependency.dependency_id IN (?)

// SELECT child.dependency_id AS policy_id, parent.policy_id AS dependent_id
// FROM policy_dependency AS parent
// JOIN policy_dependency AS child ON child.policy_id = parent.dependency_id
// WHERE child.dependency_id IN (?)

func (s *dbStorage) getDependentsWithNames(ctx context.Context, tx *goqu.TxDatabase, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.Policy, error) {
	_ = "STUB: not implemented"
	// Rather than writing a proper recursive query (which is pretty much impossible to do in a database-agnostic way), we're
	// exploiting the fact that we have a maximum of two levels of dependency (resourcePolicy -> derivedRoles -> exportVariables).
	return nil, nil
}

// SELECT
//	policy_dependency.dependency_id AS policy_id,
//	policy_dependency.policy_id AS dependent_id
//  policy.kind
//  policy.name
//  policy.version
//  policy.scope
// FROM policy_dependency
// JOIN policy ON (policy.id = policy_dependency.policy_id AND policy.disabled = false)
// WHERE policy_dependency.dependency_id IN (?)

// SELECT
//  child.dependency_id AS policy_id,
//  parent.policy_id AS dependent_id
//  policy.kind
//  policy.name
//  policy.version
//  policy.scope
// FROM policy_dependency AS parent
// JOIN policy_dependency AS child ON child.policy_id = parent.dependency_id
// JOIN policy ON (policy.id = parent.policy_id)
// WHERE child.dependency_id IN (?)

func (s *dbStorage) GetDescendants(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) getDescendants(ctx context.Context, tx *goqu.TxDatabase, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.Policy, error) {
	_ = "STUB: not implemented"
	// SELECT
	//  policy_ancestor.ancestor_id AS id,
	//  policy.id AS descendant_id,
	//  policy.kind,
	//  policy.name,
	//  policy.version,
	//  COALESCE(policy.scope, '') AS scope
	// FROM policy_ancestor
	// JOIN policy ON (policy_ancestor.policy_id = policy.id AND policy.disabled = false)
	// WHERE policy_ancestor.ancestor_id IN (?)
	// ORDER BY
	//  kind ASC,
	//  name ASC,
	//  version ASC,
	//  scope ASC
	return nil, nil
}

func (s *dbStorage) Delete(ctx context.Context, policyKey ...string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Build a deduplicated union of dependents across the whole batch,
// excluding policies deleted in this batch.

func (s *dbStorage) Disable(ctx context.Context, policyKey ...string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *dbStorage) Enable(ctx context.Context, policyKey ...string) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *dbStorage) validateIntegrity(
	ctx context.Context,
	tx *goqu.TxDatabase,
	mIDPolicyKey map[namer.ModuleID]string,
	mIDs []namer.ModuleID,
	dependents map[namer.ModuleID][]namer.Policy,
) error {
	_ = "STUB: not implemented"
	return nil
}

// writeBreaksScopeChainErrors checks whether deleting or disabling the given policies breaks the scope chain and writes the relevant information to out map.
func (s *dbStorage) writeBreaksScopeChainErrors(
	ctx context.Context,
	out map[string]*responsev1.IntegrityErrors,
	tx *goqu.TxDatabase,
	mIDPolicyKey map[namer.ModuleID]string,
	mIDs []namer.ModuleID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// we are not breaking the chain if the request includes all the descendants.

// writeRequiredByOtherPoliciesErrors checks whether deleting or disabling the given policies break dependents and writes the relevant information to out map.
func (s *dbStorage) writeRequiredByOtherPoliciesErrors(
	out map[string]*responsev1.IntegrityErrors,
	mIDPolicyKey map[namer.ModuleID]string,
	dependents map[namer.ModuleID][]namer.Policy,
) error {
	_ = "STUB: not implemented"
	return nil
}

// we are not breaking the dependents if the request includes all the dependents.

// hasAllPolicies checks if mIDPolicyKey includes all the elements from the policies slice.
// Policies slice could be a list of descendants or dependents issued to be deleted or disabled.
func (s *dbStorage) hasAllPolicies(mIDPolicyKey map[namer.ModuleID]string, policies []namer.Policy) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *dbStorage) InspectPolicies(ctx context.Context, listParams storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) ListPolicyIDs(ctx context.Context, listParams storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type postRegexpFilter struct {
	re  *regexp.Regexp
	col string
}

func (s *dbStorage) whereExprAndPostFilters(listParams storage.ListPolicyIDsParams) (whereExprs []exp.Expression, postFilters []postRegexpFilter, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// updateRegexpFilters updates either `whereExprs` or `postFilters` in place, dependent on whether regexp support is enabled or not.
func (s *dbStorage) updateRegexpFilters(namePattern, col string, whereExprs *[]exp.Expression, postFilters *[]postRegexpFilter) error {
	_ = "STUB: not implemented"
	return nil
}

// We need to pass a *regexp.Regexp expression to `Like` (or `RegexpLike`, which is equivalent) in order for goqu
// to correctly parse the query. We need to pass the compiled expression in order for goqu to access (only) the
// raw string (https://github.com/doug-martin/goqu/blob/master/exp/bool.go#L148).
// We use a cache to prevent the need to recompile arbitrary strings on each request.
// In the case of the SQLite driver, to support regexp, we generate an application-defined function in which we
// use the cached compiled expressions.

func (s *dbStorage) regexpEnabled() bool { _ = "STUB: not implemented"; return false }

func checkPostFilters(pc namer.PolicyCoords, postFilters []postRegexpFilter) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *dbStorage) ListSchemaIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *dbStorage) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (s *dbStorage) Reload(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// CheckSchema verifies the tables required by cerbos are available.
func (s *dbStorage) CheckSchema(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// ListRevisions list number of revisions for policies in the store.
func (s *dbStorage) ListRevisions(ctx context.Context, ids ...namer.ModuleID) (map[namer.ModuleID]int, error) {
	_ = "STUB: not implemented"
	// SELECT
	//  id,
	//  COUNT(*) AS count
	// FROM policy_revision
	// GROUP BY id
	return nil, nil
}

// WHERE id IN (<>)

// PurgeRevisions deletes revisions from the relevant table.
// If keepLast parameter is not specified, deletes all revisions from the table.
// If specified, it leaves the last N revisions for each policy.
func (s *dbStorage) PurgeRevisions(ctx context.Context, keepLast uint32) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// DELETE
// FROM policy_revision
