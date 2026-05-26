// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build tests

package internal

import (
	"testing"

	"github.com/cerbos/cerbos/internal/storage"
)

// MutateStoreFn points to a function which mutates the store (ex: add, delete a policy).
type MutateStoreFn func() error

func TestSuiteReloadable(store storage.Store, initFn, addFn, deleteFn MutateStoreFn) func(*testing.T) {
	_ = "STUB: not implemented"
	//nolint:thelper
	return nil
}
