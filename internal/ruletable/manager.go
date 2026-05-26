// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package ruletable

import (
	"context"
	"sync"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/conditions"
	"github.com/cerbos/cerbos/internal/engine/policyloader"
	"github.com/cerbos/cerbos/internal/engine/tracer"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/observability/logging"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage"
)

type Manager struct {
	*RuleTable
	conf         *evaluator.Conf
	policyLoader policyloader.PolicyLoader
	schemaMgr    schema.Manager
	log          *logging.Logger
	mu           sync.RWMutex
}

func NewRuleTableManager(ruleTable *RuleTable, policyLoader policyloader.PolicyLoader, schemaMgr schema.Manager) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mgr *Manager) Check(ctx context.Context, tctx tracer.Context, evalParams evaluator.EvalParams, input *enginev1.CheckInput) (*enginev1.CheckOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (mgr *Manager) Plan(ctx context.Context, input *enginev1.PlanResourcesInput, principalScope, principalVersion, resourceScope, resourceVersion string, nowFunc conditions.NowFunc, globals map[string]any, lenientScopeSearch bool) (*enginev1.PlanResourcesOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (mgr *Manager) SubscriberID() string { _ = "STUB: not implemented"; return "" }

func (mgr *Manager) OnStorageEvent(events ...storage.Event) { _ = "STUB: not implemented"; return }

func (mgr *Manager) reload() error { _ = "STUB: not implemented"; return nil }

// If compilation fails, maintain the last valid rule table state.
// Set isStale to false to prevent repeated recompilation attempts until new events arrive.

func (mgr *Manager) processPolicyEvent(evt storage.Event) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:exhaustive

// Only delete if we successfully retrieved the policy above (e.g. no compilation errors occurred)

// handle reloading dependents atomically

// we leave ruletable state static until we're sure all dependents are valid, and then update

func (mgr *Manager) addPolicy(rps *runtimev1.RunnablePolicySet) error {
	_ = "STUB: not implemented"
	return nil
}

func (mgr *Manager) deletePolicy(moduleID namer.ModuleID) error {
	_ = "STUB: not implemented"
	return nil
}

// doDeletePolicy implements the delete logic. The caller must obtain a lock first.
func (mgr *Manager) doDeletePolicy(moduleID namer.ModuleID) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(saml) many of these ephemeral caches on the RuleTable should probably now reside inside the Index layer.
