// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/google/cel-go/common/types"
)

const variablesTypeName = "cerbos.Variables"

var VariablesType = types.NewObjectType(variablesTypeName)

type Variables interface {
	IsSet(name string) bool
	Get(name string) (any, error)
}

type VariablesMap map[string]any

var _ Variables = (VariablesMap)(nil)

func (m VariablesMap) IsSet(name string) bool { _ = "STUB: not implemented"; return false }

func (m VariablesMap) Get(name string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func variablesFieldType(fieldName string) (*types.FieldType, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
