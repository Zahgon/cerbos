// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package index

import "sync"

var bitmapPool = sync.Pool{
	New: func() any {
		b := NewBitmap()
		b.words = make([]uint64, 0, 4) //nolint:mnd
		b.meta = make([]uint64, 0, 1)  //nolint:mnd
		return b
	},
}

// emptyBitmap is shared. Callers must not mutate it.
var emptyBitmap = NewBitmap()

// bitmapArena tracks pooled bitmaps acquired during a single query so they can
// be released in bulk when the query completes.
type bitmapArena struct {
	used []*Bitmap
}

func newBitmapArena() *bitmapArena { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func (a *bitmapArena) release() { _ = "STUB: not implemented"; return }

// get returns a cleared bitmap from the pool and tracks it for later release.
func (a *bitmapArena) get() *Bitmap { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert

// orInto ORs all parts into a pooled bitmap using in-place Or, avoiding
// intermediate bitmap allocations. Starts with the largest bitmap so that
// ensure allocates once to the final size.
func (a *bitmapArena) orInto(parts []*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// and2 ANDs exactly two bitmaps into a fresh pooled bitmap. Copies the
// shorter bitmap first to minimise intermediate work.
func (a *bitmapArena) and2(x, y *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// andInto ANDs bitmaps into a fresh pooled bitmap. Copies the shortest
// bitmap first to minimise intermediate work.
func (a *bitmapArena) andInto(bitmaps []*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }
