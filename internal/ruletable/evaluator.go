// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package ruletable

import (
	"context"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/schema"
)

func NewEvaluator(evalConf *evaluator.Conf, schemaConf *schema.Conf, ruleTable *RuleTable) (*Evaluator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Evaluator struct {
	evalConf  *evaluator.Conf
	schemaMgr schema.Manager
	ruleTable *RuleTable
}

func (e *Evaluator) Check(ctx context.Context, inputs []*enginev1.CheckInput, opts ...evaluator.CheckOpt) ([]*enginev1.CheckOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (e *Evaluator) Plan(ctx context.Context, input *enginev1.PlanResourcesInput, opts ...evaluator.CheckOpt) (*enginev1.PlanResourcesOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type withoutAuditTrail Evaluator

var _ evaluator.Evaluator = (*withoutAuditTrail)(nil)

func (w *withoutAuditTrail) Check(ctx context.Context, inputs []*enginev1.CheckInput, opts ...evaluator.CheckOpt) ([]*enginev1.CheckOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *withoutAuditTrail) Plan(ctx context.Context, input *enginev1.PlanResourcesInput, opts ...evaluator.CheckOpt) (*enginev1.PlanResourcesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
