// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package hub

import (
	"context"
	"errors"
	"fmt"
	"io"

	"go.uber.org/zap"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/ruletable"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/util"
)

const DriverName = "hub"

var _ storage.BinaryStore = (*HybridStore)(nil)

var ErrBundleNotLoaded = errors.New("bundle not loaded yet")

func init() {
	storage.RegisterDriver(DriverName, func(ctx context.Context, confW *config.Wrapper) (storage.Store, error) {
		conf, err := GetConfFromWrapper(confW)
		if err != nil {
			return nil, fmt.Errorf("failed to read hub configuration: %w", err)
		}

		return NewStore(ctx, conf)
	})

	storage.RegisterDriver("bundle", func(ctx context.Context, confW *config.Wrapper) (storage.Store, error) {
		util.DeprecationReplacedWarning(storage.ConfKey+".bundle", confKey)
		conf := new(Conf)
		if err := confW.Get(storage.ConfKey+".bundle", conf); err != nil {
			return nil, fmt.Errorf("failed to read bundle configuration: %w", err)
		}

		return NewStore(ctx, conf)
	})
}

func NewStore(ctx context.Context, conf *Conf) (storage.BinaryStore, error) {
	_ = "STUB: not implemented"
	return *new(storage.BinaryStore), nil
}

type Source interface {
	SourceKind() string
	storage.BinaryStore
	ruletable.RuleTableStore
}

type HybridStore struct {
	log             *zap.Logger
	local           Source
	remote          Source
	remoteIsHealthy func() bool
}

func (*HybridStore) Driver() string { _ = "STUB: not implemented"; return "" }

func (hs *HybridStore) GetRuleTable() (*ruletable.RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hs *HybridStore) withActiveSource() Source { _ = "STUB: not implemented"; return *new(Source) }

func (hs *HybridStore) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hs *HybridStore) ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hs *HybridStore) ListSchemaIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hs *HybridStore) LoadSchema(ctx context.Context, id string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (hs *HybridStore) GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hs *HybridStore) GetAll(ctx context.Context) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hs *HybridStore) GetAllMatching(ctx context.Context, candidates []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hs *HybridStore) Subscribe(s storage.Subscriber) { _ = "STUB: not implemented"; return }

func (hs *HybridStore) Unsubscribe(s storage.Subscriber) { _ = "STUB: not implemented"; return }

func (hs *HybridStore) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

func (hs *HybridStore) SourceKind() string { _ = "STUB: not implemented"; return "" }

func (hs *HybridStore) Close() (outErr error) { _ = "STUB: not implemented"; return nil }
