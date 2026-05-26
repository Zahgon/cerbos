// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import (
	"github.com/cerbos/cerbos/internal/policy"
)

type Sort struct {
	SortBy SortBy `help:"Sort policies by column" default:""`
}

func (s Sort) Validate(kind policy.Kind, listing bool) error { _ = "STUB: not implemented"; return nil }

//nolint:exhaustive

type SortBy string

const (
	SortByNone     SortBy = ""
	SortByPolicyID SortBy = "policyId"
	SortByName     SortBy = "name"
	SortByVersion  SortBy = "version"
)
