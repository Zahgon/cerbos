// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"

	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
)

type ErrSource string

const (
	ErrSourcePrincipal ErrSource = "P.attr"
	ErrSourceResource  ErrSource = "R.attr"
)

func (e ErrSource) toProto() schemav1.ValidationError_Source {
	_ = "STUB: not implemented"
	return *new(schemav1.ValidationError_Source)
}

func newValidationError(err *jsonschema.ValidationError, source ErrSource) ValidationError {
	_ = "STUB: not implemented"
	return *new(ValidationError)
}

type validationErrorFilter func(*jsonschema.ValidationError) bool

func newValidationErrorList(validationErr *jsonschema.ValidationError, source ErrSource, filter validationErrorFilter) ValidationErrorList {
	_ = "STUB: not implemented"
	return *new(ValidationErrorList)
}

func NewLoadErr(source ErrSource, schema string, err error) ValidationErrorList {
	_ = "STUB: not implemented"
	return *new(ValidationErrorList)
}

func newLoadErr(source ErrSource, message string) ValidationErrorList {
	_ = "STUB: not implemented"
	return *new(ValidationErrorList)
}

type ValidationError struct {
	Path    string
	Message string
	Source  ErrSource
}

func (e ValidationError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ValidationError) toProto() *schemav1.ValidationError { _ = "STUB: not implemented"; return nil }

type ValidationErrorList []ValidationError

func (e ValidationErrorList) ErrOrNil() error { _ = "STUB: not implemented"; return nil }

func (e ValidationErrorList) Error() string { _ = "STUB: not implemented"; return "" }

func (e ValidationErrorList) ErrorMessages() []string { _ = "STUB: not implemented"; return nil }

func (e ValidationErrorList) SchemaErrors() []*schemav1.ValidationError {
	_ = "STUB: not implemented"
	return nil
}
