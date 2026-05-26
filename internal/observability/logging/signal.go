// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !windows && !js && !wasm

package logging

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// handleUSR1Signal temporarily sets the log level to debug when a SIGUSR1 signal is received.
func handleUSR1Signal(ctx context.Context, originalLevel zapcore.Level, atomicLevel *zap.AtomicLevel) {
	_ = "STUB: not implemented"
	return
}
