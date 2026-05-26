// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package evaluator

import (
	"context"
)

func NewCheckOptions(ctx context.Context, conf *Conf, opts ...CheckOpt) *CheckOptions {
	_ = "STUB: not implemented"
	return nil
}
