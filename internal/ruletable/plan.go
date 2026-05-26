// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package ruletable

import (
	"context"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	"github.com/cerbos/cerbos/internal/conditions"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/ruletable/planner"
	"github.com/cerbos/cerbos/internal/schema"
)

func (rt *RuleTable) Plan(ctx context.Context, conf *evaluator.Conf, schemaMgr schema.Manager, input *enginev1.PlanResourcesInput, opts ...evaluator.CheckOpt) (*enginev1.PlanResourcesOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (rt *RuleTable) planWithAuditTrail(
	ctx context.Context,
	schemaMgr schema.Manager,
	input *enginev1.PlanResourcesInput,
	principalScope, principalVersion, resourceScope, resourceVersion string,
	nowFunc conditions.NowFunc, globals map[string]any, lenientScopeSearch bool,
) (*enginev1.PlanResourcesOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// evaluate resource policies before principal policies

// Principal rules are role agnostic (they treat the rows as having a `*` role). Therefore we can
// break out of the loop after the first iteration as it covers all potential principal rows.

// Once a child OVERRIDE_PARENT scope has matched an unconditional ALLOW, no
// parent-scope rule can change the outcome, so we skip

//nolint:nestif

// principal ID is only passed for principal policies; for resource
// policies an empty string means "match all principals".

//nolint:nestif

//nolint:exhaustive

// ignore constant false DENY nodes

//nolint:nestif

// only an ALLOW from a scope with ScopePermissions_SCOPE_PERMISSIONS_REQUIRE_PARENTAL_CONSENT_FOR_ALLOWS exists with no
// matching rules in the parent scopes, therefore null the node

// Const DENY overrides any ALLOW in the same role. Check both deny types.

// Roles are evaluated independently, therefore an ALLOW for one role needs to override a DENY for another.
// If we pass the role level `DENY==true`, we end up overriding the result for all roles with an `AND(..., NOT(true))`
// due to the policyTypeDenyNode inversion below. Inverting and resolving in the allow node ensures the role is OR'd
// against others, e.g. `OR(false, roleAllow1, roleAllow2, ...)`).

// Break out of the roles loop entirely

// If there is a role policy restriction for this specific role, we must apply it here.
// An ALLOW from this role is valid ONLY IF it is NOT denied by this role's role policy.

// PolicyType denies need to reside at the top level of their PolicyType sub trees (e.g. a conditional
// DENY in a principal policy needs to be a top level `(NOT deny condition) AND nested ALLOW`), so we
// invert and AND them as we go.

// reset a conditional DENY to an unconditional one

// for each action

func noMatchPlanOutput(input *enginev1.PlanResourcesInput, validationErrors []*schemav1.ValidationError) *enginev1.PlanResourcesOutput {
	_ = "STUB: not implemented"
	return nil
}

func addNode(curr, next *planner.QpN, combine func([]*planner.QpN) *planner.QpN) *planner.QpN {
	_ = "STUB: not implemented"
	return nil
}

// gateByChildOverrideAllow narrows a parent-scope DENY to only fire when no child OVERRIDE_PARENT
// scope has already produced an ALLOW.
func gateByChildOverrideAllow(childOverrideAllow, deny *planner.QpN) *planner.QpN {
	_ = "STUB: not implemented"
	return nil
}
