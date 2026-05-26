// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package compile

import (
	"errors"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/types/known/structpb"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

var errInvalidConstantName = errors.New("invalid constant name")

func compilePolicyConstants(modCtx *moduleCtx, constants *policyv1.Constants) {
	_ = "STUB: not implemented"
	return
}

func compileExportConstants(modCtx *moduleCtx) { _ = "STUB: not implemented"; return }

type constantDefinitions struct {
	modCtx  *moduleCtx
	values  map[string]*structpb.Value
	sources map[string][]*constantCtx
	used    map[string]struct{}
}

func newConstantDefinitions(modCtx *moduleCtx) *constantDefinitions {
	_ = "STUB: not implemented"
	return nil
}

func (cd *constantDefinitions) Compile(definitions map[string]*structpb.Value, path, source string) {
	_ = "STUB: not implemented"
	return
}

func (cd *constantDefinitions) Import(from *moduleCtx, source string) {
	_ = "STUB: not implemented"
	return
}

func (cd *constantDefinitions) Add(name string, value *structpb.Value, constCtx *constantCtx) {
	_ = "STUB: not implemented"
	return
}

func (cd *constantDefinitions) Resolve() { _ = "STUB: not implemented"; return }

func (cd *constantDefinitions) IsDefined(name string) bool { _ = "STUB: not implemented"; return false }

func (cd *constantDefinitions) reportRedefinedConstants() { _ = "STUB: not implemented"; return }

//nolint:mnd

func constantDefinitionPlaces(contexts []*constantCtx) []string {
	_ = "STUB: not implemented"
	return nil
}

func (cd *constantDefinitions) references(path string, expr *expr.CheckedExpr) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (cd *constantDefinitions) ResetUsage() { _ = "STUB: not implemented"; return }

func (cd *constantDefinitions) Use(path string, expr *expr.CheckedExpr) {
	_ = "STUB: not implemented"
	return
}

func (cd *constantDefinitions) Used() map[string]*structpb.Value {
	_ = "STUB: not implemented"
	return nil
}
