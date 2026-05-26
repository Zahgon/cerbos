// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package util

func logError(globExpr string, err error) { _ = "STUB: not implemented"; return }
