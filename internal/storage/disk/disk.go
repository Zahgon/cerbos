// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package disk

import (
	"context"
	"fmt"
	"io"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/storage/index"
)

const DriverName = "disk"

var (
	_ storage.SourceStore  = (*Store)(nil)
	_ storage.Reloadable   = (*Store)(nil)
	_ storage.Subscribable = (*Store)(nil)
)

func init() {
	storage.RegisterDriver(DriverName, func(ctx context.Context, confW *config.Wrapper) (storage.Store, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read disk configuration: %w", err)
		}

		return NewStore(ctx, conf)
	})
}

type Store struct {
	conf   *Conf
	idx    index.Index
	source *auditv1.PolicySource
	*storage.SubscriptionManager
}

func NewStore(ctx context.Context, conf *Conf) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFromIndex(idx index.Index) (*Store, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFromIndexWithConf(idx index.Index, conf *Conf) *Store {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) Driver() string { _ = "STUB: not implemented"; return "" }

func (s *Store) GetFirstMatch(_ context.Context, candidates []namer.ModuleID) (*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetAll(ctx context.Context) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetAllMatching(_ context.Context, modIDs []namer.ModuleID) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetCompilationUnits(_ context.Context, ids ...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetDependents(_ context.Context, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) ListSchemaIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) LoadSchema(ctx context.Context, url string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (s *Store) LoadPolicy(ctx context.Context, file ...string) ([]*policy.Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (s *Store) Reload(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }
