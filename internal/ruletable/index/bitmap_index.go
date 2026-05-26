// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package index

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

// bitmapIndex is the core in-memory bitmap index. Each binding gets a uint32 ID,
// and each (dimension, value) pair has a bitmap tracking which binding IDs
// have that value. Queries are flat bitmap AND operations.
type bitmapIndex struct {
	action             *globDimension
	coresBySum         map[uint64]*FunctionalCore // dedup behavioural part
	version            dimension[string]
	scope              dimension[string]
	role               *globDimension
	policyKind         dimension[policyv1.Kind]
	resource           *globDimension
	fqnBindings        fqnDimension
	principal          dimension[string]
	universe           *Bitmap
	allowActionsBitmap *Bitmap
	freeIDs            []uint32
	bindings           []*Binding
}

func newBitmapIndex() *bitmapIndex { _ = "STUB: not implemented"; return nil }

// allocID returns a uint32 to use as a binding's bit position across all bitmaps.
// It reuses IDs freed by freeID so that repeated add/remove cycles don't grow the
// bindings slice or make the bitmaps increasingly sparse.
func (idx *bitmapIndex) allocID() uint32 { _ = "STUB: not implemented"; return 0 }

func (idx *bitmapIndex) freeID(id uint32) { _ = "STUB: not implemented"; return }

func (idx *bitmapIndex) addBinding(b *Binding) { _ = "STUB: not implemented"; return }

// Scope "" is a valid literal (root scope), always indexed. Other dimensions
// skip "" to avoid leaking empties from policies that don't participate in
// them (e.g. principal-policy noop rows have no role/resource).

// removeBinding removes the binding from the slice and all dimension bitmaps,
// and returns the ID to the free list.
// It does NOT touch fqnBindings — that is managed by DeletePolicy, which needs
// to inspect fqnBindings across origins before deciding whether to remove the binding.
func (idx *bitmapIndex) removeBinding(b *Binding) { _ = "STUB: not implemented"; return }

func (idx *bitmapIndex) getBinding(id uint32) *Binding { _ = "STUB: not implemented"; return nil }

// dimension is a thin wrapper around map[T]*Bitmap for exact-match dimensions.
type dimension[T comparable] struct {
	m map[T]*Bitmap
}

func newDimension[T comparable]() dimension[T] { _ = "STUB: not implemented"; return nil }

func (d dimension[T]) Add(key T, id uint32) { _ = "STUB: not implemented"; return }

func (d dimension[T]) Remove(key T, id uint32) { _ = "STUB: not implemented"; return }

func (d dimension[T]) Get(key T) (*Bitmap, bool) { _ = "STUB: not implemented"; return nil, false }

func (d dimension[T]) Delete(key T) { _ = "STUB: not implemented"; return }

func (d dimension[T]) Keys() []T { _ = "STUB: not implemented"; return nil }

// Query returns OR(d[k] for k in keys).
// The returned bitmap may alias a stored bitmap; callers must not mutate it.
func (d dimension[T]) Query(arena *bitmapArena, keys []T) *Bitmap {
	_ = "STUB: not implemented"
	return nil
}

// fqnDimension maps policy FQNs to the binding IDs that originated from them.
type fqnDimension struct {
	m map[string][]uint32
}

const fqnSliceInitCap = 4

func newFqnDimension() fqnDimension { _ = "STUB: not implemented"; return *new(fqnDimension) }

func (d fqnDimension) Add(fqn string, id uint32) { _ = "STUB: not implemented"; return }

func (d fqnDimension) Get(fqn string) ([]uint32, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (d fqnDimension) Delete(fqn string) { _ = "STUB: not implemented"; return }
