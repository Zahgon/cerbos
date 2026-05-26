// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package compile

import (
	"context"

	"go.uber.org/zap"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
)

type Manager struct {
	store storage.SourceStore
	log   *zap.SugaredLogger
}

func NewManager(ctx context.Context, store storage.SourceStore) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Manager) compile(unit *policy.CompilationUnit) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Manager) GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Manager) GetAll(ctx context.Context) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Manager) GetAllMatching(ctx context.Context, modIDs []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Manager) GetPolicySet(ctx context.Context, modID namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Manager) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

type PolicyCompilationErr struct {
	underlying error
}

func (pce PolicyCompilationErr) Error() string { _ = "STUB: not implemented"; return "" }

func (pce PolicyCompilationErr) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (pce PolicyCompilationErr) Is(target error) bool { _ = "STUB: not implemented"; return false }
