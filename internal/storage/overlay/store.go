// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package overlay

import (
	"context"
	"fmt"
	"io"

	"github.com/sony/gobreaker/v2"
	"go.uber.org/zap"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/engine/policyloader"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/storage"
)

const DriverName = "overlay"

var (
	_ Overlay             = (*Store)(nil)
	_ storage.BinaryStore = (*Store)(nil)
	_ storage.Reloadable  = (*Store)(nil)
)

func init() {
	storage.RegisterDriver(DriverName, func(ctx context.Context, confW *config.Wrapper) (storage.Store, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read overlay configuration: %w", err)
		}

		return NewStore(ctx, conf, confW)
	})
}

func NewStore(ctx context.Context, conf *Conf, confW *config.Wrapper) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Store struct {
	log                  *zap.SugaredLogger
	conf                 *Conf
	baseStore            storage.Store
	fallbackStore        storage.Store
	basePolicyLoader     policyloader.PolicyLoader
	fallbackPolicyLoader policyloader.PolicyLoader
	circuitBreaker       *gobreaker.CircuitBreaker[any]
}

func newCircuitBreaker(conf *Conf) *gobreaker.CircuitBreaker[any] {
	_ = "STUB: not implemented"
	return nil
}

// GetOverlayPolicyLoader instantiates both the base and fallback policy loaders and then returns itself.
func (s *Store) GetOverlayPolicyLoader(ctx context.Context) (policyloader.PolicyLoader, error) {
	_ = "STUB: not implemented"
	return *new(policyloader.PolicyLoader), nil
}

func withCircuitBreaker0[T any](s *Store, baseFn, fallbackFn func() T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

//nolint:forcetypeassert

func withCircuitBreaker[T any](s *Store, baseFn, fallbackFn func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

//nolint:forcetypeassert

//
// PolicyLoader interface
//

func (s *Store) GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetAllMatching(ctx context.Context, modIDs []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//
// Store interface methods
//

func (s *Store) Driver() string { _ = "STUB: not implemented"; return "" }

func (s *Store) GetAll(ctx context.Context) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
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

func (s *Store) Reload(ctx context.Context) error {
	_ = "STUB: not implemented"
	// We attempt to reload all stores in parallel, regardless of base/fallback configuration.
	// Attempts on non-Reloadable stores will result in a noop.
	return nil
}

func (s *Store) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

func (s *Store) Subscribe(subscriber storage.Subscriber) { _ = "STUB: not implemented"; return }

func (s *Store) Unsubscribe(subscriber storage.Subscriber) { _ = "STUB: not implemented"; return }

func (s *Store) Close() (outErr error) { _ = "STUB: not implemented"; return nil }
