// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package compile

import (
	"google.golang.org/protobuf/types/known/emptypb"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/schema"
)

type compilerVersionMigration func(*runtimev1.RunnablePolicySet) error

const anyRoleVal = "*"

var (
	emptyVal = &emptypb.Empty{}

	compilerVersionMigrations = []compilerVersionMigration{
		migrateFromCompilerVersion0To1,
		migrateFromCompilerVersion1To2,
	}

	compilerVersion = uint32(len(compilerVersionMigrations))
)

func BatchCompile(queue <-chan *policy.CompilationUnit, schemaMgr schema.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// Compile compiles a single policy compilation unit into a runnable policy set.
// The schemaMgr parameter is optional - pass nil to skip schema validation,
// or provide a SchemaManager instance to enable validation against schema files.
func Compile(unit *policy.CompilationUnit, schemaMgr schema.Manager) (rps *runtimev1.RunnablePolicySet, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compileRolePolicySet(modCtx *moduleCtx) *runtimev1.RunnablePolicySet {
	_ = "STUB: not implemented"
	return nil
}

func compileResourcePolicySet(modCtx *moduleCtx, schemaMgr schema.Manager) *runtimev1.RunnablePolicySet {
	_ = "STUB: not implemented"
	return nil
}

// Only schema in effect is the schema defined by the "root" policy.

// TODO(cell) Check for inconsistent schema references in the policy tree.
// Either all policies must have the same schema references or only the "root" policy must have one and others should be empty.
// This should be a compiler warning instead of an error.

func compileResourcePolicy(modCtx *moduleCtx, schemaMgr schema.Manager) (*runtimev1.RunnableResourcePolicySet_Policy, *policyv1.SourceAttributes) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

func compileImportedDerivedRoles(modCtx *moduleCtx, rp *policyv1.ResourcePolicy) (map[string]*runtimev1.RunnableDerivedRole, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// used to dedupe error messages

func compileDerivedRoles(modCtx *moduleCtx) map[string]*runtimev1.RunnableDerivedRole {
	_ = "STUB: not implemented"
	return nil
}

// TODO(cell) Because derived roles can be imported many times, cache the result to avoid repeating the work

//nolint:staticcheck

func checkReferencedSchemas(modCtx *moduleCtx, rp *policyv1.ResourcePolicy, schemaMgr schema.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func compileResourceRule(modCtx *moduleCtx, path string, rule *policyv1.ResourceRule) *runtimev1.RunnableResourcePolicySet_Policy_Rule {
	_ = "STUB: not implemented"
	return nil
}

func compileOutput(modCtx *moduleCtx, path string, out *policyv1.Output) *runtimev1.Output {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

func compilePrincipalPolicySet(modCtx *moduleCtx) *runtimev1.RunnablePolicySet {
	_ = "STUB: not implemented"
	return nil
}

func compilePrincipalPolicy(modCtx *moduleCtx) (*runtimev1.RunnablePrincipalPolicySet_Policy, *policyv1.SourceAttributes) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

func reportMissingAncestors(modCtx *moduleCtx) { _ = "STUB: not implemented"; return }

// MigrateCompiledPolicies modifies a RunnablePolicySet compiled by a previous version of Cerbos to migrate it to the latest format.
func MigrateCompiledPolicies(policies *runtimev1.RunnablePolicySet) error {
	_ = "STUB: not implemented"
	return nil
}

func migrateFromCompilerVersion0To1(policies *runtimev1.RunnablePolicySet) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

func migrateFromCompilerVersion1To2(policies *runtimev1.RunnablePolicySet) error {
	_ = "STUB: not implemented"
	return nil
}
