// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package compile

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	sourcev1 "github.com/cerbos/cerbos/api/genpb/cerbos/source/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/parser"
	"github.com/cerbos/cerbos/internal/policy"
)

type unitCtx struct {
	unit   *policy.CompilationUnit
	errors *ErrorSet
}

func newUnitCtx(unit *policy.CompilationUnit) *unitCtx { _ = "STUB: not implemented"; return nil }

func (uc *unitCtx) error() error { _ = "STUB: not implemented"; return nil }

func (uc *unitCtx) moduleCtx(id namer.ModuleID) *moduleCtx { _ = "STUB: not implemented"; return nil }

type moduleCtx struct {
	*unitCtx
	def        *policyv1.Policy
	srcCtx     parser.SourceCtx
	constants  *constantDefinitions
	variables  *variableDefinitions
	fqn        string
	sourceFile string
}

func (mc *moduleCtx) error() error { _ = "STUB: not implemented"; return nil }

func (mc *moduleCtx) addErrWithDesc(err error, description string, params ...any) {
	_ = "STUB: not implemented"
	return
}

func (mc *moduleCtx) addErrForMapKeyAtProtoPath(path string, err error, description string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (mc *moduleCtx) addErrForValueAtProtoPath(path string, err error, description string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (mc *moduleCtx) addErrWithPositionAndContext(pos *sourcev1.Position, context string, err error, description string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (mc *moduleCtx) constantCtx(source, path string) *constantCtx {
	_ = "STUB: not implemented"
	return nil
}

type constantCtx struct {
	*moduleCtx
	path   string
	source string
}

func (cc *constantCtx) withSource(source string) *constantCtx {
	_ = "STUB: not implemented"
	return nil
}

func (mc *moduleCtx) variableCtx(source, path string) *variableCtx {
	_ = "STUB: not implemented"
	return nil
}

type variableCtx struct {
	*moduleCtx
	path   string
	source string
}

func (vc *variableCtx) withSource(source string) *variableCtx {
	_ = "STUB: not implemented"
	return nil
}
