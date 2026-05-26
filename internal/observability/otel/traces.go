// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package otel

import (
	"context"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracesWithExporter(ctx context.Context, env Env, exporter tracesdk.SpanExporter) (func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func InitTraces(ctx context.Context, env Env) (func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkOutdatedConfig() {
	_ = "STUB: not implemented"

	// if the tracing block exists, this would result in an error because it cannot be unmarshaled into a struct{}
	return
}

func doInitTraces(ctx context.Context, env Env, exporter tracesdk.SpanExporter) (func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createSampler(env Env) (sampler tracesdk.Sampler, err error) {
	_ = "STUB: not implemented"
	return *new(tracesdk.Sampler), nil
}

func decorateSampler(s tracesdk.Sampler) tracesdk.Sampler {
	_ = "STUB: not implemented"
	return *new(tracesdk.Sampler)
}

type sampler struct {
	s tracesdk.Sampler
}

func (s sampler) ShouldSample(params tracesdk.SamplingParameters) tracesdk.SamplingResult {
	_ = "STUB: not implemented"
	return *new(tracesdk.SamplingResult)
}

func (s sampler) Description() string { _ = "STUB: not implemented"; return "" }

type otelErrHandler func(err error)

func (o otelErrHandler) Handle(err error) { _ = "STUB: not implemented"; return }
