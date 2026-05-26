// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package ruletable

import (
	"context"

	"github.com/google/cel-go/common/types/ref"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/types/known/structpb"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	"github.com/cerbos/cerbos/internal/engine/tracer"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/ruletable/index"
	"github.com/cerbos/cerbos/internal/ruletable/internal"
	"github.com/cerbos/cerbos/internal/schema"
)

func (rt *RuleTable) Check(ctx context.Context, conf *evaluator.Conf, schemaMgr schema.Manager, inputs []*enginev1.CheckInput, opts ...evaluator.CheckOpt) ([]*enginev1.CheckOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Primary use for this Evaluator interface is the ePDP, so we run the checks synchronously (for now)

func (rt *RuleTable) checkWithAuditTrail(ctx context.Context, tctx tracer.Context, schemaMgr schema.Manager, evalParams evaluator.EvalParams, input *enginev1.CheckInput) (*enginev1.CheckOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// update the output

func (rt *RuleTable) check(ctx context.Context, tctx tracer.Context, schemaMgr schema.Manager, evalParams evaluator.EvalParams, input *enginev1.CheckInput) (*policyEvalResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate the input

// We can cache evaluated conditions for combinations of parameters and conditions.
// We use a compound key comprising the parameter origin and the rule FQN.

// Reset `actionEffectInfo` for this policy type with the correct policy key.
// This ensures we use the right policy name if no rules match

// Principal rules are role agnostic (they treat the rows as having a `*` role). Therefore we can
// break out of the loop after the first iteration as it covers all potential principal rows.

// a "policy" exists, regardless of potentially matching rules, so we update the policyKey

// This is for backwards compatibility with effectiveDerivedRoles.
// If we reach this point, we can assert that the given {origin policy + scope} combination has been evaluated
// and therefore we build the effectiveDerivedRoles from those referenced in the policy.
//nolint:nestif
//nolint:nestif

// we don't use `conditionCache` as we don't do any evaluations scoped solely to derived role conditions

// principal ID is only passed for principal policies; for resource
// policies an empty string means "match all principals".

//nolint:nestif

// We evaluate the derived role condition (if any) first, as this leads to a more sane engine trace output.

// Derived role engine trace logs are handled above. Because derived role conditions are baked into the rule table rows, we don't want to
// confuse matters by adding condition trace logs if a rule is referencing a derived role, so we pass a no-op context here.
// TODO(saml) we could probably pre-compile the condition also

// terminate early if the derived role condition isn't satisfied, which is consistent with the pre-rule table implementation

//nolint:nestif

// Implicit DENY generated as a result of no matching role policy action
// needs to be attributed to said role policy

//nolint:exhaustive

// Match the first result

// Finalise and return the first independent ALLOW

// Override `noMatchScopePermissions` DENYs with explicit ones for clarity

// Skip to next action if this action already has a definitive result from principal policies

type EffectInfo struct {
	Policy string
	Scope  string
	Effect effectv1.Effect
}

type policyEvalResult struct {
	effects               map[string]EffectInfo
	effectiveDerivedRoles map[string]struct{}
	toResolve             map[string]struct{}
	auditTrail            *auditv1.AuditTrail
	validationErrors      []*schemav1.ValidationError
	outputs               []*enginev1.OutputEntry
}

func newEvalResult(actions []string, auditTrail *auditv1.AuditTrail) *policyEvalResult {
	_ = "STUB: not implemented"
	return nil
}

func (er *policyEvalResult) unresolvedActions() []string { _ = "STUB: not implemented"; return nil }

// setEffect sets the effect for an action. DENY always takes precedence.
func (er *policyEvalResult) setEffect(action string, effect EffectInfo) {
	_ = "STUB: not implemented"
	return
}

func newAuditTrail(srcAttr map[string]*policyv1.SourceAttributes) *auditv1.AuditTrail {
	_ = "STUB: not implemented"
	return nil
}

func checkInputToRequest(input *enginev1.CheckInput) *enginev1.Request {
	_ = "STUB: not implemented"
	return nil
}

type EvalContext struct {
	request               *enginev1.Request
	runtime               *enginev1.Runtime
	effectiveDerivedRoles internal.StringSet
	programCache          *ProgramCache
	evaluator.EvalParams
}

func NewEvalContext(ep evaluator.EvalParams, request *enginev1.Request, programCache *ProgramCache) *EvalContext {
	_ = "STUB: not implemented"
	return nil
}

func (ec *EvalContext) withEffectiveDerivedRoles(effectiveDerivedRoles internal.StringSet) *EvalContext {
	_ = "STUB: not implemented"
	return nil
}

func (ec *EvalContext) lazyRuntime() any {
	_ = "STUB: not implemented" // We have to return `any` rather than `*enginev1.Runtime` here to be able to use this function as a lazy binding in the CEL evaluator.
	return *new(any)
}

func (ec *EvalContext) evaluateVariables(ctx context.Context, tctx tracer.Context, constants map[string]any, variables []*runtimev1.Variable) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ec *EvalContext) buildEvalVars(constants, variables map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (ec *EvalContext) evaluatePrograms(tctx tracer.Context, constants map[string]any, celPrograms []*index.CelProgram) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore errors for expressions that evaluate to an error value (e.g., missing keys).
// This matches the behavior of evaluateCELExprToRaw which returns nil for such cases.

func (ec *EvalContext) SatisfiesCondition(ctx context.Context, tctx tracer.Context, cond *runtimev1.Condition, constants, variables map[string]any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ec *EvalContext) evaluateBoolCELExpr(ctx context.Context, expr *exprpb.CheckedExpr, constants, variables map[string]any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ec *EvalContext) evaluateOutput(ctx context.Context, tctx tracer.Context, name, src, action string, expr *exprpb.CheckedExpr, constants, variables map[string]any) *enginev1.OutputEntry {
	_ = "STUB: not implemented"
	return nil
}

func (ec *EvalContext) evaluateOutputExpr(ctx context.Context, expr *exprpb.CheckedExpr, constants, variables map[string]any) (*structpb.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Something is broken in `ConvertToNative`

func (ec *EvalContext) evaluateCELExpr(ctx context.Context, expr *exprpb.CheckedExpr, constants, variables map[string]any) (ref.Val, error) {
	_ = "STUB: not implemented"
	return *new(ref.Val), nil
}

func (ec *EvalContext) evaluateCELExprToRaw(ctx context.Context, expr *exprpb.CheckedExpr, constants, variables map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
