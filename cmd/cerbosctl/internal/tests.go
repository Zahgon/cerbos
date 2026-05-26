// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build tests

package internal

import (
	"testing"
	"time"

	"github.com/cerbos/cerbos-sdk-go/testutil"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/flagset"
)

const (
	adminUsername = "cerbos"
	adminPassword = "cerbosAdmin"
	readyTimeout  = 30 * time.Second
)

func StartTestServer(t *testing.T) *testutil.CerbosServerInstance {
	_ = "STUB: not implemented"
	return nil
}

func CreateGlobalsFlagset(t *testing.T, address string) *flagset.Globals {
	_ = "STUB: not implemented"
	return nil
}
