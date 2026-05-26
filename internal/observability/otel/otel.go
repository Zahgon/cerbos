// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package otel

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

type EnvVar struct {
	Name string
	Alt  string
}

var (
	DisabledEV               = EnvVar{Name: "OTEL_SDK_DISABLED"}
	MetricsEndpointEV        = EnvVar{Name: "OTEL_EXPORTER_OTLP_METRICS_ENDPOINT", Alt: "OTEL_EXPORTER_OTLP_ENDPOINT"}
	MetricsExporterEV        = EnvVar{Name: "OTEL_METRICS_EXPORTER"}
	MetricsExportIntervalEV  = EnvVar{Name: "OTEL_METRIC_EXPORT_INTERVAL"}
	MetricsExportTimeoutEV   = EnvVar{Name: "OTEL_METRIC_EXPORT_TIMEOUT"}
	MetricsProtocolEV        = EnvVar{Name: "OTEL_EXPORTER_OTLP_METRICS_PROTOCOL", Alt: "OTEL_EXPORTER_OTLP_PROTOCOL"}
	ServiceNameEV            = EnvVar{Name: "OTEL_SERVICE_NAME"}
	TracesEndpointEV         = EnvVar{Name: "OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", Alt: "OTEL_EXPORTER_OTLP_ENDPOINT"}
	TracesEndpointInsecureEV = EnvVar{Name: "OTEL_EXPORTER_OTLP_TRACES_INSECURE", Alt: "OTEL_EXPORTER_OTLP_INSECURE"}
	TracesExporterEV         = EnvVar{Name: "OTEL_TRACES_EXPORTER"}
	TracesSamplerEV          = EnvVar{Name: "OTEL_TRACES_SAMPLER"}
	TracesSamplerArgEV       = EnvVar{Name: "OTEL_TRACES_SAMPLER_ARG"}
	TracesProtocolEV         = EnvVar{Name: "OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", Alt: "OTEL_EXPORTER_OTLP_PROTOCOL"}
)

const (
	NoneExporter       = "none"
	OTLPExporter       = "otlp"
	PrometheusExporter = "prometheus"

	GRPCProtocol         = "grpc"
	HTTPProtobufProtocol = "http/protobuf"

	AlwaysOffSampler               = "always_off"
	AlwaysOnSampler                = "always_on"
	JaegerRemoteSampler            = "jaeger_remote"
	ParentBasedAlwaysOffSampler    = "parentbased_always_off"
	ParentBasedAlwaysOnSampler     = "parentbased_always_on"
	ParentBasedJaegerRemoteSampler = "parentbased_jaeger_remote"
	ParentBasedTraceIDRatioSampler = "parentbased_traceidratio"
	TraceIDRatioSampler            = "traceidratio"
)

var noopCloseFn = func() error { return nil }

type Env func(string) (string, bool)

func (env Env) Get(ev EnvVar) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (env Env) GetOrDefault(ev EnvVar, defaultVal string) string {
	_ = "STUB: not implemented"
	return ""
}

func newResource(ctx context.Context, serviceName string) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
