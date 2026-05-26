// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build tests

package internal

import (
	"context"
	"testing"
	"time"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
)

const timeout = 2 * time.Second

//nolint:mnd
func TestSuite(store DBStorage) func(*testing.T) {
	_ = "STUB: not implemented"
	//nolint:thelper
	return nil
}

func withScope(p *policyv1.Policy, scope string) policy.Wrapper {
	_ = "STUB: not implemented"
	//nolint:exhaustive
	return *new(policy.Wrapper)
}

func TestCheckSchema(ctx context.Context, t *testing.T, s storage.Verifiable) {
	_ = "STUB: not implemented"
	return
}

func requireCompilationUnits(t *testing.T, want map[policy.Wrapper][]policy.Wrapper, have map[namer.ModuleID]*policy.CompilationUnit) {
	_ = "STUB: not implemented"
	return
}

func requireCompilationUnit(t *testing.T, wantModID namer.ModuleID, wantDefinitions []policy.Wrapper, have *policy.CompilationUnit) {
	_ = "STUB: not implemented"
	return
}
