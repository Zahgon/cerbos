// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build tests

package test

import (
	"context"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

// AnyContext is a function that can be passed to mock.MatchedBy to match any context.
func AnyContext(context.Context) bool {
	_ = "STUB: not implemented"

	// AnyPolicy is a function that can be passed to mock.MatchedBy to match any policy.
	return false
}

func AnyPolicy(*policyv1.Policy) bool { _ = "STUB: not implemented"; return false }
