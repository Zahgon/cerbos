// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package util

import (
	"google.golang.org/protobuf/types/known/structpb"
)

func ToStructPB(v any) (*structpb.Value, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO (cell) Recurse
