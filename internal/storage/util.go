// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package storage

import (
	"context"

	"github.com/cerbos/cerbos/internal/policy"
)

const MaxPoliciesInBatch = 25

func BatchLoadPolicy(
	ctx context.Context,
	maxPoliciesInBatch int,
	loadPolicyFn func(context.Context, ...string) ([]*policy.Wrapper, error),
	processPolicyFn func(*policy.Wrapper) error,
	ids ...string,
) error {
	_ = "STUB: not implemented"
	return nil
}
