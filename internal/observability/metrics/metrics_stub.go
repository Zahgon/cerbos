// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package metrics

import (
	"context"
)

func KindKey(string) string   { _ = "STUB: not implemented"; return "" }
func DriverKey(string) string { _ = "STUB: not implemented"; return "" }

func CompileDuration() any { _ = "STUB: not implemented"; return *new(any) }

func IndexCRUDCount() any             { _ = "STUB: not implemented"; return *new(any) }
func IndexEntryCount() any            { _ = "STUB: not implemented"; return *new(any) }
func StorePollCount() any             { _ = "STUB: not implemented"; return *new(any) }
func StoreSyncErrorCount() any        { _ = "STUB: not implemented"; return *new(any) }
func StoreLastSuccessfulRefresh() any { _ = "STUB: not implemented"; return *new(any) }

func Record(context.Context, ...any) { _ = "STUB: not implemented"; return }
func RecordDuration2[T any](any, func() (T, error)) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func Inc(context.Context, ...any) { _ = "STUB: not implemented"; return }
func Add(context.Context, ...any) { _ = "STUB: not implemented"; return }
