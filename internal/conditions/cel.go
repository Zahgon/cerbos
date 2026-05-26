// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package conditions

import (
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/decls"
	"github.com/google/cel-go/ext"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/conditions/types"
)

const (
	CELRequestIdent      = "request"
	CELResourceAbbrev    = "R"
	CELResourceKindField = "kind"
	CELResourceField     = "resource"
	CELPrincipalAbbrev   = "P"
	CELPrincipalField    = "principal"
	CELRuntimeIdent      = "runtime"
	CELConstantsIdent    = "constants"
	CELConstantsAbbrev   = "C"
	CELVariablesIdent    = "variables"
	CELVariablesAbbrev   = "V"
	CELGlobalsIdent      = "globals"
	CELGlobalsAbbrev     = "G"
	CELAttrField         = "attr"
	CELScopeField        = "scope"
)

var (
	TrueExpr  *exprpb.CheckedExpr
	FalseExpr *exprpb.CheckedExpr

	StdEnv *cel.Env

	StdEnvDecls = []*decls.VariableDecl{
		decls.NewVariable(CELRequestIdent, types.MessageType[*enginev1.Request]()),
		decls.NewVariable(CELPrincipalAbbrev, types.MessageType[*enginev1.Request_Principal]()),
		decls.NewVariable(CELResourceAbbrev, types.MessageType[*enginev1.Request_Resource]()),
		decls.NewVariable(CELRuntimeIdent, types.RuntimeType),
		decls.NewVariable(CELConstantsIdent, types.VariablesType),
		decls.NewVariable(CELConstantsAbbrev, types.VariablesType),
		decls.NewVariable(CELVariablesIdent, types.VariablesType),
		decls.NewVariable(CELVariablesAbbrev, types.VariablesType),
		decls.NewVariable(CELGlobalsIdent, types.VariablesType),
		decls.NewVariable(CELGlobalsAbbrev, types.VariablesType),
	}

	variablesType *exprpb.Type
)

func init() {
	var err error

	StdEnv, err = cel.NewEnv(
		ext.TwoVarComprehensions(),
		cel.CrossTypeNumericComparisons(true),
		cel.Types(&enginev1.Request{}, &enginev1.Request_Principal{}, &enginev1.Request_Resource{}, &enginev1.Runtime{}),
		cel.VariableDecls(StdEnvDecls...),
		ext.Lists(),
		ext.Bindings(),
		ext.Strings(),
		ext.Encoders(),
		ext.Math(),
		CerbosCELLib(),
		types.Registry(),
	)
	if err != nil {
		panic(fmt.Errorf("failed to initialize standard CEL environment: %w", err))
	}

	FalseExpr, err = compileConstant("false")
	if err != nil {
		panic(fmt.Errorf("failed to compile constant 'false': %w", err))
	}

	TrueExpr, err = compileConstant("true")
	if err != nil {
		panic(fmt.Errorf("failed to compile constant 'true': %w", err))
	}

	variablesType, err = cel.TypeToExprType(types.VariablesType)
	if err != nil {
		panic(fmt.Errorf("failed to convert cerbos.Variables type to proto: %w", err))
	}
}

func compileConstant(value string) (*exprpb.CheckedExpr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Fqn(s string) string { _ = "STUB: not implemented"; return "" }

func ResourceFqn(s string) string { _ = "STUB: not implemented"; return "" }

func ResourceAttributeNames(s string) []string { _ = "STUB: not implemented"; return nil }

// R.attr.<s>
// request.resource.attr.<s>

func ResourceFieldNames(s string) []string { _ = "STUB: not implemented"; return nil }

// R.<s>
// request.resource.<s>

func PrincipalFieldNames(s string) []string { _ = "STUB: not implemented"; return nil }

// P.<s>
// request.principal.<s>

func ExpandAbbrev(s string) string { _ = "STUB: not implemented"; return "" }
