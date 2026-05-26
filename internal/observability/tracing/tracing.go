// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package tracing

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/trace"

	"github.com/cerbos/cerbos/internal/engine/tracer"
)

func HTTPHandler(handler http.Handler, path string) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func StartTracer(sink tracer.Sink) tracer.Context {
	_ = "STUB: not implemented"
	return *new(tracer.Context)
}

func MarkFailed(span trace.Span, code int, err error) { _ = "STUB: not implemented"; return }

func RecordSpan(ctx context.Context, name string, fn func(context.Context, trace.Span)) {
	_ = "STUB: not implemented"
	return
}

func RecordSpan1(ctx context.Context, name string, fn func(context.Context, trace.Span) error) error {
	_ = "STUB: not implemented"
	return nil
}

func RecordSpan2[T any](ctx context.Context, name string, fn func(context.Context, trace.Span) (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
