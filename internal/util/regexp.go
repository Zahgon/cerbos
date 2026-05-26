// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package util

import (
	"regexp"
	"sync"
)

type RegexpCache struct {
	cache map[string]*regexp.Regexp
	mu    sync.RWMutex
}

func NewRegexpCache() *RegexpCache { _ = "STUB: not implemented"; return nil }

// GetCompiledExpr lazily compiles (and stores) regexp.
func (c *RegexpCache) GetCompiledExpr(re string) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
