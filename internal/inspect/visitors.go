// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package inspect

import (
	"github.com/google/cel-go/common/ast"

	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
)

func attributeVisitor(attrs map[string]*responsev1.InspectPoliciesResponse_Attribute) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func constantVisitor(consts map[string]*responsev1.InspectPoliciesResponse_Constant) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

func variableVisitor(vars map[string]*responsev1.InspectPoliciesResponse_Variable) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}
