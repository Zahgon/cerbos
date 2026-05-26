// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package hub

import (
	"context"
	"io"
	"sync"

	cloudapi "github.com/cerbos/cloud-api/bundle"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/ruletable"
	"github.com/cerbos/cerbos/internal/storage"
)

var (
	_ storage.BinaryStore = (*LocalSource)(nil)
	_ storage.Reloadable  = (*LocalSource)(nil)
)

// LocalSource loads a bundle from local disk.
type LocalSource struct {
	bundle  Bundle
	cleanup func() error
	*storage.SubscriptionManager
	params LocalParams
	mu     sync.RWMutex
}

func NewLocalSourceFromConf(ctx context.Context, conf *Conf) (*LocalSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type LocalParams struct {
	BundlePath    string
	TempDir       string
	SecretKey     string
	EncryptionKey []byte
	CacheSize     uint
	BundleVersion cloudapi.Version
}

func NewLocalSource(ctx context.Context, params LocalParams) (*LocalSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) loadBundle() error { _ = "STUB: not implemented"; return nil }

func (ls *LocalSource) activeBundleID() string { _ = "STUB: not implemented"; return "" }

func (ls *LocalSource) Driver() string { _ = "STUB: not implemented"; return "" }

func (ls *LocalSource) GetRuleTable() (*ruletable.RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) (ids []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) ListSchemaIDs(ctx context.Context) (ids []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) LoadSchema(ctx context.Context, id string) (schema io.ReadCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (ls *LocalSource) GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (ps *runtimev1.RunnablePolicySet, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) GetAll(ctx context.Context) (pss []*runtimev1.RunnablePolicySet, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) GetAllMatching(ctx context.Context, modIDs []namer.ModuleID) (pss []*runtimev1.RunnablePolicySet, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *LocalSource) Reload(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (ls *LocalSource) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (ls *LocalSource) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

func (ls *LocalSource) Close() error { _ = "STUB: not implemented"; return nil }

func (ls *LocalSource) SourceKind() string { _ = "STUB: not implemented"; return "" }
