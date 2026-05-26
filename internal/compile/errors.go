// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package compile

import (
	"errors"

	"github.com/google/cel-go/cel"

	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
)

var (
	errAmbiguousDerivedRole   = errors.New("ambiguous derived role")
	errConstantRedefined      = errors.New("constant redefined")
	errCyclicalVariables      = errors.New("cyclical variable definitions")
	errEmptyOutput            = errors.New("empty output")
	errImportNotFound         = errors.New("import not found")
	errInvalidCompilationUnit = errors.New("invalid compilation unit")
	errInvalidResourceRule    = errors.New("invalid resource rule")
	errInvalidSchema          = errors.New("invalid schema")
	errMissingDefinition      = errors.New("missing policy definition")
	errScriptsUnsupported     = errors.New("scripts in conditions are no longer supported")
	errUndefinedConstant      = errors.New("undefined constant")
	errUndefinedVariable      = errors.New("undefined variable")
	errUnexpectedErr          = errors.New("unexpected error")
	errUnknownDerivedRole     = errors.New("unknown derived role")
	errVariableRedefined      = errors.New("variable redefined")
)

type ErrorSet struct {
	CompileErrors map[uint64]*runtimev1.CompileErrors_Err
}

func newErrorSet() *ErrorSet { _ = "STUB: not implemented"; return nil }

func (e *ErrorSet) Errors() *runtimev1.CompileErrors { _ = "STUB: not implemented"; return nil }

func (e *ErrorSet) ErrOrNil() error { _ = "STUB: not implemented"; return nil }

func (e *ErrorSet) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorSet) Add(err error) { _ = "STUB: not implemented"; return }

// Error describes an error encountered during compilation.
type Error struct {
	*runtimev1.CompileErrors_Err
}

func (e *Error) Display() string { _ = "STUB: not implemented"; return "" }

func errorDisplay(err *runtimev1.CompileErrors_Err) string { _ = "STUB: not implemented"; return "" }

func (e *Error) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func errorString(err *runtimev1.CompileErrors_Err) string { _ = "STUB: not implemented"; return "" }

func newError(file, desc string, err error) *Error { _ = "STUB: not implemented"; return nil }

// CELCompileError holds CEL compilation errors.
type CELCompileError struct {
	issues *cel.Issues
	expr   string
}

func newCELCompileError(expr string, issues *cel.Issues) *CELCompileError {
	_ = "STUB: not implemented"
	return nil
}

func (cce *CELCompileError) Error() string { _ = "STUB: not implemented"; return "" }

func (cce *CELCompileError) Unwrap() error { _ = "STUB: not implemented"; return nil }
