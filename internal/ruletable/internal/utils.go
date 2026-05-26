// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"google.golang.org/protobuf/types/known/emptypb"
)

const anyRoleVal = "*"

type ProtoSet map[string]*emptypb.Empty

// Merge merges keys from `o` into the original ProtoSet.
func (p ProtoSet) Merge(o ProtoSet) { _ = "STUB: not implemented"; return }

type StringSet map[string]struct{}

func (s StringSet) Values() []string { _ = "STUB: not implemented"; return nil }

func (s StringSet) IsSubSetOf(o StringSet) bool { _ = "STUB: not implemented"; return false }

func (s StringSet) UnionWith(o StringSet) { _ = "STUB: not implemented"; return }

func ToSet(values []string) StringSet { _ = "STUB: not implemented"; return *new(StringSet) }

func SetIntersects(s1 ProtoSet, s2 StringSet) bool { _ = "STUB: not implemented"; return false }

func SubtractSets(s1, s2 StringSet) { _ = "STUB: not implemented"; return }

func GetSymmetricDifference(s1, s2 StringSet) StringSet {
	_ = "STUB: not implemented"
	return *new(StringSet)
}
