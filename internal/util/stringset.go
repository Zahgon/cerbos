// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package util

func ToStringSet(values []string) StringSet { _ = "STUB: not implemented"; return *new(StringSet) }

type StringSet map[string]struct{}

func (ss StringSet) Contains(value string) bool { _ = "STUB: not implemented"; return false }
