// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/google/cel-go/common/types"
)

var (
	RuntimeType     = MessageType[*enginev1.Runtime]()
	runtimeTypeName = RuntimeType.TypeName()
)

type Runtime interface {
	GetEffectiveDerivedRoles() []string
}

var _ Runtime = (*enginev1.Runtime)(nil)

func runtimeFieldType(fieldName string) (*types.FieldType, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
