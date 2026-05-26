// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package tracer

import (
	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

// Singleton components for stateless trace components to reduce allocations.
var (
	componentConditionAll  = &enginev1.Trace_Component{Kind: enginev1.Trace_Component_KIND_CONDITION_ALL}
	componentConditionAny  = &enginev1.Trace_Component{Kind: enginev1.Trace_Component_KIND_CONDITION_ANY}
	componentConditionNone = &enginev1.Trace_Component{Kind: enginev1.Trace_Component_KIND_CONDITION_NONE}
	componentCondition     = &enginev1.Trace_Component{Kind: enginev1.Trace_Component_KIND_CONDITION}
	componentVariables     = &enginev1.Trace_Component{Kind: enginev1.Trace_Component_KIND_VARIABLES}
)

func Start(sink Sink) Context { _ = "STUB: not implemented"; return *new(Context) }

type context struct {
	sink       Sink
	components []*enginev1.Trace_Component
}

func (c *context) StartAction(action string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) StartConditionAll() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartConditionAny() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartConditionNone() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartCondition() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartDerivedRole(name string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) StartExpr(expr string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartNthCondition(index int) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) StartPolicy(name string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartRolePolicyScope(scope string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) StartResource(kind string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) StartRole(role string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartRule(name string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartScope(scope string) Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartVariable(name, expr string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) StartVariables() Context { _ = "STUB: not implemented"; return *new(Context) }

func (c *context) StartOutput(ruleName string) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) start(component *enginev1.Trace_Component) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (c *context) Activated() { _ = "STUB: not implemented"; return }

func (c *context) AppliedEffect(effect effectv1.Effect, message string) {
	_ = "STUB: not implemented"
	return
}

func (c *context) ComputedBoolResult(result bool, err error, message string) {
	_ = "STUB: not implemented"
	return
}

func (c *context) ComputedOutput(output *enginev1.OutputEntry) { _ = "STUB: not implemented"; return }

func (c *context) ComputedResult(result any) { _ = "STUB: not implemented"; return }

func protobufValue(goValue any) *structpb.Value {
	_ = "STUB: not implemented"
	// Fast path for common primitive types to avoid JSON marshal/unmarshal overhead.
	return nil
}

// Fall back to JSON for complex types (structs, slices, maps, proto messages).

func (c *context) Failed(err error, message string) { _ = "STUB: not implemented"; return }

func (c *context) Skipped(err error, message string) { _ = "STUB: not implemented"; return }

func (c *context) addTrace(event *enginev1.Trace_Event) { _ = "STUB: not implemented"; return }

func errorString(err error) string { _ = "STUB: not implemented"; return "" }
