// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package tracer

import (
	"sync"

	"go.uber.org/zap"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cespare/xxhash/v2"
)

type Sink interface {
	Enabled() bool
	AddTrace(trace *enginev1.Trace)
}

type zapSink struct {
	log *zap.Logger
}

func NewZapSink(log *zap.Logger) *zapSink { _ = "STUB: not implemented"; return nil }

func (zs *zapSink) Enabled() bool { _ = "STUB: not implemented"; return false }

func (zs *zapSink) AddTrace(trace *enginev1.Trace) { _ = "STUB: not implemented"; return }

func zapTrace(trace *enginev1.Trace) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }

type Collector struct {
	traces []*enginev1.Trace
	mutex  sync.RWMutex
}

func NewCollector() *Collector { _ = "STUB: not implemented"; return nil }

func (c *Collector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (c *Collector) AddTrace(trace *enginev1.Trace) { _ = "STUB: not implemented"; return }

func (c *Collector) Traces() []*enginev1.Trace { _ = "STUB: not implemented"; return nil }

const defaultCapacity = 256

var (
	defIndexPool = sync.Pool{New: func() any { return make(map[uint64]uint32, defaultCapacity) }}
	ptrHashPool  = sync.Pool{New: func() any { return make(map[*enginev1.Trace_Component]uint64, defaultCapacity) }}
	serBufPool   = sync.Pool{New: func() any { b := make([]byte, 0, defaultCapacity); return &b }}
	hasherPool   = sync.Pool{New: func() any { return xxhash.New() }}
)

func hashComponentVT(comp *enginev1.Trace_Component, hasher *xxhash.Digest, buf []byte) (uint64, []byte) {
	_ = "STUB: not implemented"
	return 0, nil
}

func TracesToBatch(traces []*enginev1.Trace) *enginev1.TraceBatch {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert
//nolint:forcetypeassert
//nolint:forcetypeassert

//nolint:forcetypeassert

func BatchToTraces(batch *enginev1.TraceBatch) []*enginev1.Trace {
	_ = "STUB: not implemented"
	return nil
}
