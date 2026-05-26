// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package engine

import (
	"context"
	"sync/atomic"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/engine/policyloader"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/ruletable"
	"github.com/cerbos/cerbos/internal/schema"
)

var _ evaluator.Evaluator = (*Engine)(nil)

const (
	parallelismThreshold = 5
	workerQueueSize      = 4
	workerResetJitter    = 1 << 4
	workerResetThreshold = 1 << 16
)

func ApplyCheckOptions(opts ...evaluator.CheckOpt) *evaluator.CheckOptions {
	_ = "STUB: not implemented"
	return nil
}

type Engine struct {
	schemaMgr         schema.Manager
	auditLog          audit.Log
	policyLoader      policyloader.PolicyLoader
	ruleTableManager  *ruletable.Manager
	conf              *evaluator.Conf
	metadataExtractor audit.MetadataExtractor
	workerPool        []chan<- workIn
	workerIndex       atomic.Uint64
}

type Components struct {
	AuditLog          audit.Log
	PolicyLoader      policyloader.PolicyLoader
	RuleTableManager  *ruletable.Manager
	SchemaMgr         schema.Manager
	MetadataExtractor audit.MetadataExtractor
}

func New(ctx context.Context, components Components) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFromConf(ctx context.Context, conf *evaluator.Conf, components Components) *Engine {
	_ = "STUB: not implemented"
	return nil
}

func NewEphemeral(conf *evaluator.Conf, rtMgr *ruletable.Manager, schemaMgr schema.Manager) *Engine {
	_ = "STUB: not implemented"
	return nil
}

func newEngine(conf *evaluator.Conf, c Components) *Engine { _ = "STUB: not implemented"; return nil }

func (engine *Engine) startWorker(ctx context.Context, num int, inputChan <-chan workIn) {
	_ = "STUB: not implemented"
	// Keep each goroutine around for a period of time and then recycle them to reclaim the stack space.
	// See https://adtac.in/2021/04/23/note-on-worker-pools-in-go.html
	return
}

//nolint:gosec

// restart to clear the stack

func (engine *Engine) submitWork(ctx context.Context, work workIn) error {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) Plan(ctx context.Context, input *enginev1.PlanResourcesInput, opts ...evaluator.CheckOpt) (*enginev1.PlanResourcesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (engine *Engine) doPlan(ctx context.Context, input *enginev1.PlanResourcesInput, opts *evaluator.CheckOptions) (*enginev1.PlanResourcesOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (engine *Engine) logPlanDecision(ctx context.Context, input *enginev1.PlanResourcesInput, output *enginev1.PlanResourcesOutput, planErr error, trail *auditv1.AuditTrail) (*enginev1.PlanResourcesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (engine *Engine) Check(ctx context.Context, inputs []*enginev1.CheckInput, opts ...evaluator.CheckOpt) ([]*enginev1.CheckOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the number of inputs is less than the threshold, do a serial execution as it is usually faster.
// ditto if the worker pool is not initialized

func (engine *Engine) logCheckDecision(ctx context.Context, inputs []*enginev1.CheckInput, outputs []*enginev1.CheckOutput, checkErr error, trail *auditv1.AuditTrail) ([]*enginev1.CheckOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (engine *Engine) checkSerial(ctx context.Context, inputs []*enginev1.CheckInput, checkOpts *evaluator.CheckOptions) ([]*enginev1.CheckOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (engine *Engine) checkParallel(ctx context.Context, inputs []*enginev1.CheckInput, checkOpts *evaluator.CheckOptions) ([]*enginev1.CheckOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (engine *Engine) evaluate(ctx context.Context, input *enginev1.CheckInput, checkOpts *evaluator.CheckOptions) (*enginev1.CheckOutput, *auditv1.AuditTrail, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// exit early if the context is cancelled

type workOut struct {
	err    error
	result *enginev1.CheckOutput
	trail  *auditv1.AuditTrail
	index  int
}

type workIn struct {
	ctx       context.Context
	input     *enginev1.CheckInput
	checkOpts *evaluator.CheckOptions
	out       chan<- workOut
	index     int
}
