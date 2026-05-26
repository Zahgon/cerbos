// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package hub

import (
	"context"
	"io"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/ruletable"
	"github.com/cerbos/cerbos/internal/storage"
)

// instrument wraps the given source to produce metrics.
func instrument(name string, source Source) Source { _ = "STUB: not implemented"; return *new(Source) }

// instrumentedSource is a source that measures the time taken by each operation.
type instrumentedSource struct {
	source Source
	name   string
}

func (instrumentedSource) Driver() string { _ = "STUB: not implemented"; return "" }

func (is instrumentedSource) GetRuleTable() (*ruletable.RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is instrumentedSource) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is instrumentedSource) ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is instrumentedSource) ListSchemaIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is instrumentedSource) LoadSchema(ctx context.Context, id string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (is instrumentedSource) GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is instrumentedSource) GetAll(ctx context.Context) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is instrumentedSource) GetAllMatching(ctx context.Context, modIDs []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (is instrumentedSource) Reload(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (is instrumentedSource) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (is instrumentedSource) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

func (is instrumentedSource) Subscribe(s storage.Subscriber) { _ = "STUB: not implemented"; return }

func (is instrumentedSource) Unsubscribe(s storage.Subscriber) { _ = "STUB: not implemented"; return }

func (is instrumentedSource) Close() error { _ = "STUB: not implemented"; return nil }

func (is instrumentedSource) SourceKind() string { _ = "STUB: not implemented"; return "" }

func measureBinaryOp[T any](ctx context.Context, source, opName string, op func(context.Context) (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func withTrace[T any](ctx context.Context, source, opName string, op func(context.Context) (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
