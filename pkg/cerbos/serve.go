// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package cerbos

import (
	"context"

	"go.uber.org/zap/zapcore"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type serveOptions struct {
	zapCore            zapcore.Core
	configOverrides    map[string]any
	configFilePath     string
	debugListenAddr    string
	logLevel           LogLevel
	metricsDisabled    bool
	tracesDisabled     bool
	goMaxProcsDisabled bool
}

// ServeOption defines options for [Serve].
type ServeOption func(*serveOptions)

// WithConfigFile sets the path to the [server configuration file].
//
// [server configuration file]: https://docs.cerbos.dev/cerbos/latest/configuration/
func WithConfigFile(path string) ServeOption { _ = "STUB: not implemented"; return *new(ServeOption) }

// WithConfig sets [server configuration], overriding any values present in the configuration file.
//
// [server configuration]: https://docs.cerbos.dev/cerbos/latest/configuration/
func WithConfig(overrides map[string]any) ServeOption {
	_ = "STUB: not implemented"
	return *new(ServeOption)
}

// WithDebug enables the [gops] agent listening on the given host:port.
//
// [gops]: https://github.com/google/gops
func WithDebug(addr string) ServeOption { _ = "STUB: not implemented"; return *new(ServeOption) }

// WithLogLevel sets the minimum level at which logs will be emitted.
func WithLogLevel(level LogLevel) ServeOption { _ = "STUB: not implemented"; return *new(ServeOption) }

// WithMetricsDisabled disables registering OpenTelemetry metrics publishers.
func WithMetricsDisabled() ServeOption { _ = "STUB: not implemented"; return *new(ServeOption) }

// WithTracesDisabled disables registering OpenTelemetry trace publishers.
func WithTracesDisabled() ServeOption { _ = "STUB: not implemented"; return *new(ServeOption) }

// WithGoMaxProcsDisabled disables automatically setting GOMAXPROCS.
func WithGoMaxProcsDisabled() ServeOption { _ = "STUB: not implemented"; return *new(ServeOption) }

// WithZapCore sets the Zap core used by the logging system.
func WithZapCore(zapCore zapcore.Core) ServeOption {
	_ = "STUB: not implemented"
	return *new(ServeOption)
}

// Serve runs the Cerbos policy decision point server, stopping it when the context is canceled.
func Serve(ctx context.Context, options ...ServeOption) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

// initialize metrics

// load configuration

// initialize tracing

func startDebugListener(listenAddr string) { _ = "STUB: not implemented"; return }
