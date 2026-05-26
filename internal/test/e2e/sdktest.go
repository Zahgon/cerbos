// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build tests || e2e

package e2e

import (
	"testing"
	"time"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
)

const timeout = 30 * time.Second

func TestSDKClient(addr string, opts ...cerbos.Opt) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:thelper

//nolint:mnd

func generateToken(t *testing.T, expiry time.Time) string { _ = "STUB: not implemented"; return "" }

//nolint:mnd
