// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package tracer

type Sink interface {
	Enabled() bool
	AddTrace(any)
}

type noopSink struct{}

func NewZapSink(any) Sink { _ = "STUB: not implemented"; return *new(Sink) }

func (noopSink) Enabled() bool { _ = "STUB: not implemented"; return false }

func (noopSink) AddTrace(any) { _ = "STUB: not implemented"; return }
