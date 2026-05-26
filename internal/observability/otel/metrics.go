// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package otel

import (
	"context"
	"time"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

const shutdownTimeout = 5 * time.Second

func InitMetrics(ctx context.Context, env Env) (func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createOTLPMetricsExporter(ctx context.Context, env Env) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func dropHighCardinalityLabels() sdkmetric.View {
	_ = "STUB: not implemented"
	return *new(sdkmetric.View)
}
