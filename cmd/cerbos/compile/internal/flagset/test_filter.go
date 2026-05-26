// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import (
	"github.com/cerbos/cerbos/internal/verify"
)

type TestFilter []string

func (tf TestFilter) ToFilterConfig() (*verify.FilterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
