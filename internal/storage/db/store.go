// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package db

import (
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
)

type IntegrityErr struct {
	Errors map[string]*responsev1.IntegrityErrors
}

func (e *IntegrityErr) Error() string { _ = "STUB: not implemented"; return "" }
