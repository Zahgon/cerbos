// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build e2e

package e2e

import (
	"time"

	"github.com/go-cmd/cmd"
)

func Setup(ctx Ctx) error { _ = "STUB: not implemented"; return nil }

// `helmfile apply` requires `helm diff`. `helmfile init` checks for required plugins

func Teardown(ctx Ctx) error { _ = "STUB: not implemented"; return nil }

func Cmd(ctx Ctx, name string, args ...string) error { _ = "STUB: not implemented"; return nil }

func CmdWithOutput(ctx Ctx, name string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func execCmd(ctx Ctx, showOutput bool, name string, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func dumpOutput(ctx Ctx, s cmd.Status) { _ = "STUB: not implemented"; return }

func checkCerbosIsUp(ctx Ctx) func() error { _ = "STUB: not implemented"; return nil }

//nolint:gosec

func Retry(op func() error, timeout time.Duration, interval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
