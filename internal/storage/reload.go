// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package storage

import (
	"context"

	"golang.org/x/sync/singleflight"
)

var sfGroup singleflight.Group

func Reload(ctx context.Context, rs Reloadable) error { _ = "STUB: not implemented"; return nil }
