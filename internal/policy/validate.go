// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package policy

import (
	"errors"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	sourcev1 "github.com/cerbos/cerbos/api/genpb/cerbos/source/v1"
	"github.com/cerbos/cerbos/internal/parser"
)

var errEmptyPolicy = errors.New("policy is empty")

type ValidationError struct {
	Err *sourcev1.Error
}

func newValidationError(msg string, pos *sourcev1.Position, context string) ValidationError {
	_ = "STUB: not implemented"
	return *new(ValidationError)
}

func (ve ValidationError) Error() string { _ = "STUB: not implemented"; return "" }

func Validate(p *policyv1.Policy, sc parser.SourceCtx) error { _ = "STUB: not implemented"; return nil }

func validateResourcePolicy(rp *policyv1.ResourcePolicy, sc parser.SourceCtx) (outErr error) {
	_ = "STUB: not implemented"
	return nil
}

// check for rule without any roles or derived roles defined

// check for name clashes

func validatePrincipalPolicy(rp *policyv1.PrincipalPolicy, sc parser.SourceCtx) (outErr error) {
	_ = "STUB: not implemented"
	return nil
}

func validateRolePolicy(rp *policyv1.RolePolicy, sc parser.SourceCtx) error {
	_ = "STUB: not implemented"
	return nil
}

func validateDerivedRoles(dr *policyv1.DerivedRoles, sc parser.SourceCtx) (outErr error) {
	_ = "STUB: not implemented"
	return nil
}

// Check for name clashes

func validateExportConstants(p *policyv1.Policy, sc parser.SourceCtx) error {
	_ = "STUB: not implemented"
	return nil
	//nolint:staticcheck
}

func validateExportVariables(p *policyv1.Policy, sc parser.SourceCtx) error {
	_ = "STUB: not implemented"
	return nil
	//nolint:staticcheck
}
