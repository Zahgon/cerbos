// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type typeRegistry struct {
	types.Adapter
	types.Provider
}

var (
	_ types.Adapter  = (*typeRegistry)(nil)
	_ types.Provider = (*typeRegistry)(nil)
)

func Registry() cel.EnvOption { _ = "STUB: not implemented"; return *new(cel.EnvOption) }

func (r *typeRegistry) FindStructType(structType string) (*types.Type, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *typeRegistry) FindStructFieldType(structType, fieldName string) (*types.FieldType, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *typeRegistry) NativeToValue(native any) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

type object interface {
	ref.Val
	traits.FieldTester
	traits.Indexer
}

type protoMessageObject struct {
	object
	fields protoreflect.FieldDescriptors
}

func (o *protoMessageObject) IsSet(field ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func (o *protoMessageObject) Get(index ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func (o *protoMessageObject) resolveField(method func(ref.Val) ref.Val, field ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}
