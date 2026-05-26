// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package evaluator

import (
	"context"
	"time"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/conditions"
	"github.com/cerbos/cerbos/internal/engine/tracer"
)

type Evaluator interface {
	Check(context.Context, []*enginev1.CheckInput, ...CheckOpt) ([]*enginev1.CheckOutput, error)
	Plan(context.Context, *enginev1.PlanResourcesInput, ...CheckOpt) (*enginev1.PlanResourcesOutput, error)
}

// CheckOpt defines options for engine Check calls.
type CheckOpt func(*CheckOptions)

func WithTraceSink(tracerSink tracer.Sink) CheckOpt {
	_ = "STUB: not implemented"
	return *new(CheckOpt)
}

// WithNowFunc sets the function for determining `now` during condition evaluation.
// The function should return the same timestamp every time it is invoked.
func WithNowFunc(nowFunc func() time.Time) CheckOpt {
	_ = "STUB: not implemented"
	return *new(CheckOpt)
}

// WithLenientScopeSearch enables lenient scope search.
func WithLenientScopeSearch() CheckOpt { _ = "STUB: not implemented"; return *new(CheckOpt) }

// WithGlobals sets the global variables for the engine.
func WithGlobals(globals map[string]any) CheckOpt { _ = "STUB: not implemented"; return *new(CheckOpt) }

// WithDefaultPolicyVersion sets the default policy version for the engine.
func WithDefaultPolicyVersion(defaultPolicyVersion string) CheckOpt {
	_ = "STUB: not implemented"
	return *new(CheckOpt)
}

// WithDefaultScope sets the default scope for the engine.
func WithDefaultScope(defaultScope string) CheckOpt {
	_ = "STUB: not implemented"
	return *new(CheckOpt)
}

type CheckOptions struct {
	TracerSink tracer.Sink
	EvalParams EvalParams
}

func (co *CheckOptions) NowFunc() func() time.Time { _ = "STUB: not implemented"; return nil }

func (co *CheckOptions) DefaultPolicyVersion() string { _ = "STUB: not implemented"; return "" }

func (co *CheckOptions) DefaultScope() string { _ = "STUB: not implemented"; return "" }

func (co *CheckOptions) LenientScopeSearch() bool { _ = "STUB: not implemented"; return false }

func (co *CheckOptions) Globals() map[string]any { _ = "STUB: not implemented"; return nil }

type EvalParams struct {
	Globals              map[string]any
	NowFunc              conditions.NowFunc
	DefaultPolicyVersion string
	DefaultScope         string
	LenientScopeSearch   bool
}

func PolicyVersion(version string, params EvalParams) string { _ = "STUB: not implemented"; return "" }

func Scope(scope string, params EvalParams) string { _ = "STUB: not implemented"; return "" }
