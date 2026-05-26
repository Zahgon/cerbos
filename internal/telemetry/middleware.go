// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"
	"sync/atomic"
	"time"

	telemetryv1 "github.com/cerbos/cerbos/api/genpb/cerbos/telemetry/v1"
	"google.golang.org/grpc"
)

const collectorBufferSize = 64

var totalReqCount atomic.Uint64

type Interceptors interface {
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	StreamServerInterceptor() grpc.StreamServerInterceptor
}

type nopInterceptors struct{}

func (nopInterceptors) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func (nopInterceptors) StreamServerInterceptor() grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

type methodInfo struct {
	name      string
	userAgent string
}

type statsInterceptors struct {
	reporter    Reporter
	collector   chan methodInfo
	methodTally map[string]uint64
	uaTally     map[string]uint64
}

func newStatsInterceptors(reporter Reporter, interval time.Duration, shutdown <-chan struct{}) *statsInterceptors {
	_ = "STUB: not implemented"
	return nil
}

func (i *statsInterceptors) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

func (i *statsInterceptors) StreamServerInterceptor() grpc.StreamServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.StreamServerInterceptor)
}

func (i *statsInterceptors) collectStats(ctx context.Context, method string) {
	_ = "STUB: not implemented"
	return
}

func (i *statsInterceptors) doTally(interval time.Duration, shutdown <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (i *statsInterceptors) report() { _ = "STUB: not implemented"; return }

func toCountStats(m map[string]uint64) []*telemetryv1.Event_CountStat {
	_ = "STUB: not implemented"
	return nil
}
