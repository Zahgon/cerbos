// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package ruletable

import (
	"context"

	epdpv2 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/epdp/v2"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/schema"
)

type Evaluator interface {
	Check(context.Context, []*enginev1.CheckInput, ...evaluator.CheckOpt) ([]*enginev1.CheckOutput, *auditv1.AuditTrail, error)
	Plan(context.Context, *enginev1.PlanResourcesInput, ...evaluator.CheckOpt) (*enginev1.PlanResourcesOutput, *auditv1.AuditTrail, error)
}

func NewRuleTableFromProto(protoRT *runtimev1.RuleTable, conf *epdpv2.Config) (Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(Evaluator), nil
}

func evaluatorConfFromProto(confProto *epdpv2.Config_Evaluator) (*evaluator.Conf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func schemaConfFromProto(confProto *epdpv2.Config_Schema) (*schema.Conf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use default
