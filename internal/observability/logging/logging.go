// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package logging

import (
	"context"
	"sync"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultTmpLogLevelDuration = 10 * time.Minute

var initOnce sync.Once

type ctxLog struct{}

var ctxLogKey = &ctxLog{}

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(name string) *Logger { _ = "STUB: not implemented"; return nil }

// Implement missing functions here.
var (
	String  = zap.String
	Strings = zap.Strings
	Error   = zap.Error
	Uint32  = zap.Uint32
)

// InitLogging initializes the global logger.
func InitLogging(ctx context.Context, level string, zapCore zapcore.Core) {
	_ = "STUB: not implemented"
	return
}

func doInitLogging(ctx context.Context, level string, zapCore zapcore.Core) {
	_ = "STUB: not implemented"
	return
}

// setLogLevelForDuration temporarily sets the global log level to the given level for a period of time.
func setLogLevelForDuration(ctx context.Context, doneChan chan<- struct{}, extendChan <-chan struct{}, originalLevel zapcore.Level, atomicLevel *zap.AtomicLevel) {
	_ = "STUB: not implemented"
	return
}

// FromContext returns the logger from the context if one exists. Otherwise it returns a new logger.
func FromContext(ctx context.Context) *zap.Logger { _ = "STUB: not implemented"; return nil }

// ToContext adds a logger to the context.
func ToContext(ctx context.Context, log *zap.Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ReqScopeLog returns a request-scoped logger with request fields populated.
func ReqScopeLog(ctx context.Context) *zap.Logger { _ = "STUB: not implemented"; return nil }

func GRPCLogFieldsToZap(fields logging.Fields) []zap.Field { _ = "STUB: not implemented"; return nil }

//nolint:mnd
