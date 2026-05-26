// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import (
	"github.com/cerbos/cerbos/internal/policy"
)

//nolint:govet
type Filters struct { //betteralign:ignore
	Name            []string `help:"Filter policies by name"`
	NameRegexp      string   `help:"Filter policies by name, using regular expression"`
	Version         []string `help:"Filter policies by version"`
	VersionRegexp   string   `help:"Filter policies by version, using regular expression"`
	PolicyIDs       []string `help:"List of policy ids" arg:"" name:"id" optional:"" `
	ScopeRegexp     string   `help:"Filter policies by scope, using regular expression"`
	IncludeDisabled bool     `help:"Include disabled policies"`
}

func (f Filters) Validate(kind policy.Kind, listing bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:exhaustive
