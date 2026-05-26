// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package inspect

import (
	"context"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/policy"
)

type loadPolicyFn func(ctx context.Context, policyKey ...string) ([]*policy.Wrapper, error)

func Policies() *Policy { _ = "STUB: not implemented"; return nil }

type Policy struct {
	derivedRolesImports   map[string][]string
	derivedRolesToResolve map[string]map[string]bool
	constantImports       map[string][]string
	constantsToResolve    map[string]map[string]bool
	variableImports       map[string][]string
	variablesToResolve    map[string]map[string]bool
	results               map[string]*responsev1.InspectPoliciesResponse_Result
}

// Inspect inspects the given policy and caches the inspection related information internally.
func (pol *Policy) Inspect(p *policyv1.Policy) error { _ = "STUB: not implemented"; return nil }

//nolint:nestif

// sort constants if there is nothing to resolve since we are not going to modify constants in the future.

// sort variables if there is nothing to resolve since we are not going to modify variables in the future.

// Results returns the final inspection results.
func (pol *Policy) Results(ctx context.Context, loadPolicy loadPolicyFn) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pol *Policy) resolveDerivedRoles(ctx context.Context, loadPolicy loadPolicyFn) {
	_ = "STUB: not implemented"
	return
}

func (pol *Policy) resolveConstants(ctx context.Context, loadPolicy loadPolicyFn) {
	_ = "STUB: not implemented"
	return
}

//nolint:nestif

func (pol *Policy) resolveVariables(ctx context.Context, loadPolicy loadPolicyFn) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:nestif

// listReferencedAttributes lists the attributes referenced from the conditions and variables in the given policy.
func (pol *Policy) listReferencedAttributes(p *policyv1.Policy) ([]*responsev1.InspectPoliciesResponse_Attribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listReferencedDerivedRoles lists the referenced derived roles in the given resource policy.
func (pol *Policy) listReferencedDerivedRoles(rp *policyv1.ResourcePolicy) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// listReferencedConstants lists the constants referenced from the conditions in the given policy.
func (pol *Policy) listReferencedConstants(p *policyv1.Policy) (map[string]*responsev1.InspectPoliciesResponse_Constant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listReferencedVariables lists the variables referenced from the conditions in the given policy.
func (pol *Policy) listReferencedVariables(p *policyv1.Policy) (map[string]*responsev1.InspectPoliciesResponse_Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// listImports lists the derived roles and export constants/variables imported by the given policy.
func (pol *Policy) listImports(p *policyv1.Policy) (derivedRoleImports, constantImports, variableImports []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func policyKeys(names []string, fqn func(string) string) []string {
	_ = "STUB: not implemented"
	return nil
}
