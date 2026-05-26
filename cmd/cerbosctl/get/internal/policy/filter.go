// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"github.com/cerbos/cerbos/internal/policy"
)

type filterDef struct {
	names           map[string]struct{}
	versions        map[string]struct{}
	kind            policy.Kind
	includeDisabled bool
}

func newFilterDef(kind policy.Kind, names, versions []string, includeDisabled bool) *filterDef {
	_ = "STUB: not implemented"
	return nil
}

func (fd *filterDef) filter(p policy.Wrapper) bool { _ = "STUB: not implemented"; return false }
