// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package metrics

import (
	"context"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	UnitBytes         = "By"
	UnitDimensionless = "1"
	UnitMilliseconds  = "ms"
)

var (
	DriverKey      = attribute.Key("driver").String
	KindKey        = attribute.Key("kind").String
	OpKey          = attribute.Key("op").String
	RemoteEventKey = attribute.Key("remote_event").String
	ResultKey      = attribute.Key("result").String
	SourceKey      = attribute.Key("source").String
	StatusKey      = attribute.Key("status").String
)

var defaultHistBoundsMS = []float64{0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

var Meter = sync.OnceValue(func() metric.Meter {
	return otel.GetMeterProvider().Meter("cerbos.dev/cerbos")
})

var (
	AuditErrorCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_audit_error_count",
			metric.WithDescription("Number of errors encountered while writing an audit log entry"),
		)
	})

	AuditOversizedEntryCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_audit_oversized_entry_count",
			metric.WithDescription("Number of audit log entries skipped due to exceeding maximum size"),
		)
	})

	BundleFetchErrorsCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_store_bundle_fetch_errors_count",
			metric.WithDescription("Count of errors encountered during bundle downloads"),
		)
	})

	BundleNotFoundErrorsCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_store_bundle_not_found_errors_count",
			metric.WithDescription("Count of bundle not found errors"),
		)
	})

	BundleStoreLatency = once(func() (metric.Float64Histogram, error) {
		return Meter().Float64Histogram(
			"cerbos_dev_store_bundle_op_latency",
			metric.WithDescription("Time to do an operation with the bundle store"),
			metric.WithUnit(UnitMilliseconds),
			metric.WithExplicitBucketBoundaries(defaultHistBoundsMS...),
		)
	})

	BundleStoreRemoteEventsCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_store_bundle_remote_events_count",
			metric.WithDescription("Count of remote server events received by the bundle store"),
		)
	})

	BundleStoreUpdatesCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_store_bundle_updates_count",
			metric.WithDescription("Count of bundle updates from remote source"),
		)
	})

	CacheAccessCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_cache_access_count",
			metric.WithDescription("Counter of cache access"),
		)
	})

	CacheLiveObjGauge = once(func() (metric.Int64UpDownCounter, error) {
		return Meter().Int64UpDownCounter(
			"cerbos_dev_cache_live_objects",
			metric.WithDescription("Number of live objects in the cache"),
		)
	})

	CacheMaxSize = once(func() (metric.Int64UpDownCounter, error) {
		return Meter().Int64UpDownCounter(
			"cerbos_dev_cache_max_size",
			metric.WithDescription("Maximum capacity of the cache"),
		)
	})

	CompileDuration = once(func() (metric.Float64Histogram, error) {
		return Meter().Float64Histogram(
			"cerbos_dev_compiler_compile_duration",
			metric.WithDescription("Time to compile a set of policies"),
			metric.WithUnit(UnitMilliseconds),
			metric.WithExplicitBucketBoundaries(defaultHistBoundsMS...),
		)
	})

	EngineCheckLatency = once(func() (metric.Float64Histogram, error) {
		return Meter().Float64Histogram(
			"cerbos_dev_engine_check_latency",
			metric.WithDescription("Time to match a request against a policy and provide a decision"),
			metric.WithUnit(UnitMilliseconds),
			metric.WithExplicitBucketBoundaries(defaultHistBoundsMS...),
		)
	})

	EngineCheckBatchSize = once(func() (metric.Int64Histogram, error) {
		return Meter().Int64Histogram(
			"cerbos_dev_engine_check_batch_size",
			metric.WithDescription("Batch size distribution of check requests"),
			metric.WithUnit(UnitDimensionless),
			metric.WithExplicitBucketBoundaries(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 14, 16, 18, 20, 25, 30, 35, 40, 45, 50), //nolint:mnd
		)
	})

	EnginePlanLatency = once(func() (metric.Float64Histogram, error) {
		return Meter().Float64Histogram(
			"cerbos_dev_engine_plan_latency",
			metric.WithDescription("Time to produce a query plan"),
			metric.WithUnit(UnitMilliseconds),
			metric.WithExplicitBucketBoundaries(defaultHistBoundsMS...),
		)
	})

	HubConnected = once(func() (metric.Int64UpDownCounter, error) {
		return Meter().Int64UpDownCounter(
			"cerbos_dev_hub_connected",
			metric.WithDescription("Is the instance connected to Cerbos Hub"),
		)
	})

	IndexCRUDCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_index_crud_count",
			metric.WithDescription("Number of create_update_delete operations"),
		)
	})

	IndexEntryCount = once(func() (metric.Int64UpDownCounter, error) {
		return Meter().Int64UpDownCounter(
			"cerbos_dev_index_entry_count",
			metric.WithDescription("Number of entries in the index"),
		)
	})

	StorePollCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_store_poll_count",
			metric.WithDescription("Number of times the remote store was polled for updates"),
		)
	})

	StoreSyncErrorCount = once(func() (metric.Int64Counter, error) {
		return Meter().Int64Counter(
			"cerbos_dev_store_sync_error_count",
			metric.WithDescription("Number of errors encountered while syncing updates from the remote store"),
		)
	})

	StoreLastSuccessfulRefresh = once(func() (metric.Int64Gauge, error) {
		return Meter().Int64Gauge(
			"cerbos_dev_store_last_successful_refresh",
			metric.WithDescription("The last time the store was successfully refreshed manually or automatically"),
		)
	})
)

func NewHandler() (http.Handler, error) { _ = "STUB: not implemented"; return *new(http.Handler), nil }

func once[T any](fn func() (T, error)) func() T { _ = "STUB: not implemented"; return nil }

func must[T any](retVal T, err error) T { _ = "STUB: not implemented"; return *new(T) }

func RecordDuration2[T any](hist metric.Float64Histogram, fn func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func RecordDuration3[A, B any](hist metric.Float64Histogram, fn func() (A, B, error)) (A, B, error) {
	_ = "STUB: not implemented"
	return *new(A), *new(B), nil
}

func TotalTimeMS(startTime time.Time) float64 { _ = "STUB: not implemented"; return 0 }

type counter interface {
	Add(context.Context, int64, ...metric.AddOption)
}

func Inc[T counter](ctx context.Context, m T, attr ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

func Add[T counter](ctx context.Context, m T, v int64, attr ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}

type gauge interface {
	Record(context.Context, int64, ...metric.RecordOption)
}

func Record[T gauge](ctx context.Context, m T, v int64, attr ...attribute.KeyValue) {
	_ = "STUB: not implemented"
	return
}
