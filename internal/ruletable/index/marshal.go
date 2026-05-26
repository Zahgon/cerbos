// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package index

import (
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
)

// Marshal produces an immutable artefact: indexes reconstructed via Unmarshal are read-only snapshots and must not be passed to any mutating method on Index.
func (m *Index) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func marshalCore(c *FunctionalCore) (*runtimev1.BitmapIndex_FunctionalCore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalRowParams(rp *RowParams) (*runtimev1.RuleTable_RuleRow_Params, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalBinding(b *Binding, coreIndex map[*FunctionalCore]uint32) *runtimev1.BitmapIndex_Binding {
	_ = "STUB: not implemented"
	return nil
}

func marshalGlobDimension(gd *globDimension) (*runtimev1.BitmapIndex_GlobDimension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalEntries(m map[string]*Bitmap) ([]*runtimev1.BitmapIndex_Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalParentRoles(parentRoles map[string]map[string][]string) map[string]*runtimev1.BitmapIndex_RoleParents {
	_ = "STUB: not implemented"
	return nil
}

func Unmarshal(data []byte) (*Index, error) { _ = "STUB: not implemented"; return nil, nil }

func unmarshalCores(pbCores []*runtimev1.BitmapIndex_FunctionalCore) ([]*FunctionalCore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalBindings(pbBindings []*runtimev1.BitmapIndex_Binding, cores []*FunctionalCore) []*Binding {
	_ = "STUB: not implemented"
	return nil
}

// find the max ID to size the bindings slice.

func unmarshalEntries(entries []*runtimev1.BitmapIndex_Entry) (dimension[string], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalGlobDimension(pb *runtimev1.BitmapIndex_GlobDimension) (*globDimension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// can't use `unmarshalEntries` because each glob entry also
// needs its pattern compiled into gd.compiled.

func unmarshalParentRoles(pb map[string]*runtimev1.BitmapIndex_RoleParents) map[string]map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func bitmapFromBytes(data []byte) (*Bitmap, error) { _ = "STUB: not implemented"; return nil, nil }
