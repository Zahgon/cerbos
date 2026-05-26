// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build js && wasm

package cache

import (
	"time"

	"github.com/bluele/gcache"
)

type Cache[K, V any] struct {
	cache gcache.Cache
}

func New[K, V any](kind string, size uint, _ ...any) *Cache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cache[K, V]) Has(k K) bool { _ = "STUB: not implemented"; return false }

func (c *Cache[K, V]) Get(k K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

func (c *Cache[K, V]) Set(k K, v V) { _ = "STUB: not implemented"; return }

func (c *Cache[K, V]) SetWithExpire(k K, v V, expiry time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (c *Cache[K, V]) Remove(k K) bool { _ = "STUB: not implemented"; return false }

func (c *Cache[K, V]) Purge() { _ = "STUB: not implemented"; return }
