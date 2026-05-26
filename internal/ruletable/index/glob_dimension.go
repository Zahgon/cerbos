// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package index

import (
	"github.com/gobwas/glob"
)

// globDimension is a bitmap-based dimension index that supports both literal
// and glob pattern keys. It mirrors the semantics of internal.GlobMap,
// returning bitmaps for queried values.
type globDimension struct {
	// Raw maps rather than dimension[string] because globs and compiled have
	// coupled lifecycles that don't fit the dimension abstraction cleanly.
	literals map[string]*Bitmap
	// `globs` and `compiled` share the same key
	globs    map[string]*Bitmap
	compiled map[string]glob.Glob
}

func newGlobDimension() *globDimension { _ = "STUB: not implemented"; return nil }

func (gd *globDimension) Set(key string, id uint32) { _ = "STUB: not implemented"; return }

//nolint:nestif

func (gd *globDimension) Remove(key string, id uint32) { _ = "STUB: not implemented"; return }

//nolint:nestif

// Query returns the OR of the literal bitmap for value and all glob bitmaps
// whose pattern matches value. The returned bitmap may alias a stored bitmap;
// callers must not mutate it.
func (gd *globDimension) Query(arena *bitmapArena, value string) *Bitmap {
	_ = "STUB: not implemented"
	return nil
}

// Collect literal + matching glob bitmaps and combine with in-place OR.

// QueryMultiple returns OR of all bitmaps matching any of the given values.
// The returned bitmap may alias a stored bitmap; callers must not mutate it.
func (gd *globDimension) QueryMultiple(arena *bitmapArena, values []string) *Bitmap {
	_ = "STUB: not implemented"
	return nil
}

func (gd *globDimension) GetAllKeys() []string { _ = "STUB: not implemented"; return nil }

// RangeBitmaps iterates (key, bitmap) pairs. Glob keys (e.g. "manager:*")
// appear verbatim. Each bitmap aliases a stored bitmap; do not mutate.
func (gd *globDimension) RangeBitmaps(fn func(key string, bm *Bitmap)) {
	_ = "STUB: not implemented"
	return
}
