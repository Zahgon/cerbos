// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package tracer

import (
	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
)

type Context interface {
	StartAction(action string) Context
	StartConditionAll() Context
	StartConditionAny() Context
	StartConditionNone() Context
	StartCondition() Context
	StartDerivedRole(name string) Context
	StartExpr(expr string) Context
	StartNthCondition(index int) Context
	StartPolicy(name string) Context
	StartResource(kind string) Context
	StartRole(role string) Context
	StartRule(name string) Context
	StartScope(scope string) Context
	StartVariable(name, expr string) Context
	StartVariables() Context
	StartOutput(ruleName string) Context
	StartRolePolicyScope(scope string) Context
	Activated()
	AppliedEffect(effect effectv1.Effect, message string)
	ComputedBoolResult(result bool, err error, message string)
	ComputedOutput(output *enginev1.OutputEntry)
	ComputedResult(result any)
	Failed(err error, message string)
	Skipped(err error, message string)
}

type noopContext struct{}

func (c noopContext) StartAction(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartConditionAll() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartConditionAny() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartConditionNone() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartCondition() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartDerivedRole(string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c noopContext) StartExpr(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartNthCondition(int) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c noopContext) StartPolicy(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartRolePolicyScope(string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c noopContext) StartResource(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartRole(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartRule(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartScope(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartVariable(string, string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c noopContext) StartVariables() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c noopContext) StartOutput(string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (noopContext) Activated() { _ = "STUB: not implemented"; return }

func (noopContext) AppliedEffect(effectv1.Effect, string) { _ = "STUB: not implemented"; return }

func (noopContext) ComputedBoolResult(bool, error, string) { _ = "STUB: not implemented"; return }

func (noopContext) ComputedOutput(*enginev1.OutputEntry) { _ = "STUB: not implemented"; return }

func (noopContext) ComputedResult(any) { _ = "STUB: not implemented"; return }

func (noopContext) Failed(error, string) { _ = "STUB: not implemented"; return }

func (noopContext) Skipped(error, string) { _ = "STUB: not implemented"; return }
