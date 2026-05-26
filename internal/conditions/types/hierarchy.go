// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"reflect"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/decls"
	"github.com/google/cel-go/common/operators"
	"github.com/google/cel-go/common/overloads"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
)

const (
	hierarchyDelim            = "."
	hierarchyFn               = "hierarchy"
	hierarchyTypeName         = "cerbos.lib.hierarchy"
	overloadAncestorOf        = "ancestorOf"
	overloadCommonAncestors   = "commonAncestors"
	overloadDescendentOf      = "descendentOf"
	overloadImmediateChildOf  = "immediateChildOf"
	overloadImmediateParentOf = "immediateParentOf"
	overloadOverlaps          = "overlaps"
	overloadSiblingOf         = "siblingOf"
)

var (
	HierarchyType = cel.ObjectType(hierarchyTypeName,
		traits.IndexerType,
		traits.SizerType,
		traits.ReceiverType)

	HierarchyFunc = cel.Function(hierarchyFn,
		cel.Overload(
			fmt.Sprintf("%s_string", hierarchyFn),
			[]*cel.Type{cel.StringType},
			HierarchyType,
			cel.UnaryBinding(unaryHierarchyFnImpl),
		),

		cel.Overload(
			fmt.Sprintf("%s_string_string", hierarchyFn),
			[]*cel.Type{cel.StringType, cel.StringType},
			HierarchyType,
			cel.BinaryBinding(binaryHierarchyFnImpl),
		),

		cel.Overload(
			fmt.Sprintf("%s_stringarray", hierarchyFn),
			[]*cel.Type{cel.ListType(cel.StringType)},
			HierarchyType,
			cel.UnaryBinding(unaryHierarchyFnImpl),
		),
	)

	HierarchyDeclrations = []*decls.FunctionDecl{
		newFunction(overloadAncestorOf,
			decls.MemberOverload(overloadAncestorOf,
				[]*types.Type{HierarchyType, HierarchyType},
				types.BoolType,
			),
		),

		newFunction(overloadCommonAncestors,
			decls.MemberOverload(overloadCommonAncestors,
				[]*types.Type{HierarchyType, HierarchyType},
				HierarchyType,
			),
		),

		newFunction(overloadDescendentOf,
			decls.MemberOverload(overloadDescendentOf,
				[]*types.Type{HierarchyType, HierarchyType},
				types.BoolType,
			),
		),

		newFunction(overloadImmediateChildOf,
			decls.MemberOverload(overloadImmediateChildOf,
				[]*types.Type{HierarchyType, HierarchyType},
				types.BoolType,
			),
		),

		newFunction(overloadImmediateParentOf,
			decls.MemberOverload(overloadImmediateParentOf,
				[]*types.Type{HierarchyType, HierarchyType},
				types.BoolType,
			),
		),

		newFunction(overloadOverlaps,
			decls.MemberOverload(overloadOverlaps,
				[]*types.Type{HierarchyType, HierarchyType},
				types.BoolType,
			),
		),

		newFunction(overloadSiblingOf,
			decls.MemberOverload(overloadSiblingOf,
				[]*types.Type{HierarchyType, HierarchyType},
				types.BoolType,
			),
		),

		newFunction(overloads.Size,
			decls.MemberOverload(fmt.Sprintf("%s_size", hierarchyFn),
				[]*types.Type{HierarchyType},
				types.IntType,
			),
		),

		newFunction(operators.Index,
			decls.Overload(fmt.Sprintf("%s_index", hierarchyFn),
				[]*types.Type{HierarchyType, types.IntType},
				types.StringType,
			),
		),
	}

	hierarchyOneArgOverloads = map[string]func(Hierarchy, ref.Val) ref.Val{
		overloadAncestorOf:        hierarchyAncestorOf,
		overloadCommonAncestors:   hierarchyCommonAncestors,
		overloadDescendentOf:      hierarchyDescendentOf,
		overloadImmediateChildOf:  hierarchyImmediateChildOf,
		overloadImmediateParentOf: hierarchyImmediateParentOf,
		overloadOverlaps:          hierarchyOverlaps,
		overloadSiblingOf:         hierarchySiblingOf,
	}
)

func newFunction(name string, opts ...decls.FunctionOpt) *decls.FunctionDecl {
	_ = "STUB: not implemented"
	return nil
}

func unaryHierarchyFnImpl(v ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

func binaryHierarchyFnImpl(v, delim ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Hierarchy is a type that represents a dot-separated hierarchy such as a.b.c.d.
type Hierarchy []string

// ConvertToNative implements ref.Val.ConvertToNative.
func (h Hierarchy) ConvertToNative(typeDesc reflect.Type) (any, error) {
	_ = "STUB: not implemented"
	//nolint:exhaustive
	return *new(any), nil
}

// ConvertToType implements ref.Val.ConvertToType.
func (h Hierarchy) ConvertToType(typeVal ref.Type) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Type implements ref.Val.Type.
func (h Hierarchy) Type() ref.Type {
	_ = "STUB: not implemented"
	return *

	// Value implements ref.Val.Value.
	new(ref.Type)
}

func (h Hierarchy) Value() any {
	_ = "STUB: not implemented"

	// Equal implements ref.Val.Equal.
	return *new(any)
}

func (h Hierarchy) Equal(other ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

// Receive implements traits.Receiver.Receive.
func (h Hierarchy) Receive(function, _ string, args []ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Get implements traits.Indexer.Get.
func (h Hierarchy) Get(index ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

// Size implements traits.Sizer.Size.
func (h Hierarchy) Size() ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

func hierarchyAncestorOf(h Hierarchy, path ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func hierarchyCommonAncestors(h Hierarchy, path ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

//nolint:prealloc

func hierarchyDescendentOf(h Hierarchy, path ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func hierarchyImmediateChildOf(h Hierarchy, path ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func hierarchyImmediateParentOf(h Hierarchy, path ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func hierarchySiblingOf(h Hierarchy, path ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func hierarchyOverlaps(h Hierarchy, path ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func toHierarchy(v ref.Val) (Hierarchy, ref.Val) {
	_ = "STUB: not implemented"
	return *new(Hierarchy), *new(ref.Val)
}
