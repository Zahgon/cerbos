// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package logging

import "context"

type Logger struct{}

func NewLogger(string) *Logger { _ = "STUB: not implemented"; return nil }

func FromContext(context.Context) *Logger {
	_ = "STUB: not implemented"

	// Add any missing required stub methods here
	return nil
}

func (*Logger) Debug(...any)          { _ = "STUB: not implemented"; return }
func (*Logger) Debugw(string, ...any) { _ = "STUB: not implemented"; return }
func (*Logger) Debugf(string, ...any) { _ = "STUB: not implemented"; return }
func (*Logger) Warn(string, ...any)   { _ = "STUB: not implemented"; return }
func (*Logger) Warnw(string, ...any)  { _ = "STUB: not implemented"; return }
func (*Logger) Info(...any)           { _ = "STUB: not implemented"; return }

func String(string, string) any { _ = "STUB: not implemented"; return *new(any) }

func Strings(string, []string) any { _ = "STUB: not implemented"; return *new(any) }

func Error(error) any { _ = "STUB: not implemented"; return *new(any) }

func Uint32(string, uint32) any { _ = "STUB: not implemented"; return *new(any) }
