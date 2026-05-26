// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package conditions

import (
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var exprMessageName = (&runtimev1.Expr{}).ProtoReflect().Descriptor().FullName()

// WalkExprs visits all `cerbos.runtime.v1.Expr` messages within a compiled artefact.
func WalkExprs(message proto.Message, f func(*runtimev1.Expr)) { _ = "STUB: not implemented"; return }

func walkExprs(message protoreflect.Message, f func(*runtimev1.Expr)) {
	_ = "STUB: not implemented"
	return
}

//nolint:forcetypeassert

// nothing to do

// MigrateVariablesType rewrites the `checked` field, changing the type of all references
// to constants, globals, and variables from `map<string, dyn>` to `cerbos.Variables`.
func MigrateVariablesType(expr *runtimev1.Expr) { _ = "STUB: not implemented"; return }
