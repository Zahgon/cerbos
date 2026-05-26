// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"github.com/google/cel-go/common/decls"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/conditions"
)

const lastResultVar = "_"

var (
	qualifiedPrincipal = conditions.Fqn(conditions.CELPrincipalField)
	qualifiedResource  = conditions.Fqn(conditions.CELResourceField)

	specialVars = buildSpecialVarsSet()
)

func buildSpecialVarsSet() map[string]struct{} { _ = "STUB: not implemented"; return nil }

func resetVarsAndDecls() (variables, map[string]*decls.VariableDecl) {
	_ = "STUB: not implemented"
	return *new(variables), nil
}

// request and variable decls are already defined in StdEnv

func getCheckInput(vars variables) (*enginev1.CheckInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
