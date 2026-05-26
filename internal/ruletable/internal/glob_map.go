// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"sync"

	"github.com/gobwas/glob"
)

const wildcardAny = rune('*')

// GlobMap is a map that supports glob pattern matching for keys.
//
// Thread safety: GlobMap requires external synchronization for write operations
// (Set, Clear, DeleteLiteral). Read operations (Get, GetMerged, GetAll, etc.) may
// run concurrently with each other but not with writes. The internal cacheMu only
// protects matchCache, which can be written during read operations when populating
// the cache on a miss.
type GlobMap[T any] struct {
	literals   map[string]T
	globs      map[string]T
	compiled   map[string]glob.Glob // to avoid sync overhead store local refs to globally compiled globs.
	matchCache map[string][]string  // cache: lookup key -> matching glob patterns
	cacheMu    sync.RWMutex         // protects matchCache
}

func NewGlobMap[T any](m map[string]T) *GlobMap[T] { _ = "STUB: not implemented"; return nil }

func (gm *GlobMap[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// Clear resets the map for reuse, preserving compiled globs to avoid re-fetching from global cache.
func (gm *GlobMap[T]) Clear() { _ = "STUB: not implemented"; return }

// No lock needed: writes are externally serialized.

// Keep gm.compiled - same patterns likely to be reused

func (gm *GlobMap[T]) Set(k string, v T) { _ = "STUB: not implemented"; return }

// invalid glob pattern, skip, the error is logged by the callee

// No lock needed: writes are externally serialized.

func (gm *GlobMap[T]) Get(k string) (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

func (gm *GlobMap[T]) GetWithLiteral(k string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

func (gm *GlobMap[T]) DeleteLiteral(k string) { _ = "STUB: not implemented"; return }

// No lock needed: writes are externally serialized.

func (gm *GlobMap[T]) GetAll() map[string]T { _ = "STUB: not implemented"; return nil }

func (gm *GlobMap[T]) GetAllKeys() []string { _ = "STUB: not implemented"; return nil }

func (gm *GlobMap[T]) GetMerged(k string) map[string]T {
	_ = "STUB: not implemented"
	// Fast path: no globs, just check literal
	return nil
}

// Slow path

// getMatchingGlobs returns all glob patterns that match the given key, using cache.
func (gm *GlobMap[T]) getMatchingGlobs(k string) []string { _ = "STUB: not implemented"; return nil }
