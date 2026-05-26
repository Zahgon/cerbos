// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"fmt"
	"reflect"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
)

const (
	spiffeIDFn                     = "spiffeID"
	spiffeIDTypeName               = "cerbos.lib.spiffeID"
	spiffeMatcherTypeName          = "cerbos.lib.spiffeMatcher"
	spiffeMatchAnyFn               = "spiffeMatchAny"
	spiffeMatchExactFn             = "spiffeMatchExact"
	spiffeMatchOneOfFn             = "spiffeMatchOneOf"
	spiffeMatchTrustDomainFn       = "spiffeMatchTrustDomain"
	spiffeTrustDomainFn            = "spiffeTrustDomain"
	spiffeTrustDomainTypeName      = "cerbos.lib.spiffeTrustDomain"
	overloadSpiffeIDIsMemberOf     = "isMemberOf"
	overloadSpiffeIDPath           = "path"
	overloadSpiffeIDTrustDomain    = "trustDomain"
	overloadSpiffeMatcherMatchesID = "matchesID"
	overloadSpiffeTrustDomainID    = "id"
	overloadSpiffeTrustDomainName  = "name"
)

var (
	SPIFFEIDType          = cel.ObjectType(spiffeIDTypeName, traits.ReceiverType, traits.ComparerType)
	SPIFFETrustDomainType = cel.ObjectType(spiffeTrustDomainTypeName, traits.ReceiverType, traits.ComparerType)
	SPIFFEMatcherType     = cel.ObjectType(spiffeMatcherTypeName, traits.ReceiverType)

	SPIFFEIDFunc = cel.Function(spiffeIDFn,
		cel.Overload(
			fmt.Sprintf("%s_string", spiffeIDFn),
			[]*cel.Type{cel.StringType},
			SPIFFEIDType,
			cel.UnaryBinding(unarySPIFFEIDFnImpl),
		),
	)

	SPIFFETrustDomainFunc = cel.Function(spiffeTrustDomainFn,
		cel.Overload(
			fmt.Sprintf("%s_string", spiffeTrustDomainFn),
			[]*cel.Type{cel.StringType},
			SPIFFETrustDomainType,
			cel.UnaryBinding(unarySPIFFETrustDomainFnImpl),
		),
	)

	SPIFFEMatchAnyFunc = cel.Function(spiffeMatchAnyFn,
		cel.Overload(
			spiffeMatchAnyFn,
			nil,
			SPIFFEMatcherType,
			cel.FunctionBinding(unarySPIFFEMatchAnyFnImpl),
		),
	)

	SPIFFEMatchExactFunc = cel.Function(spiffeMatchExactFn,
		cel.Overload(
			fmt.Sprintf("%s_spiffeID", spiffeMatchExactFn),
			[]*cel.Type{SPIFFEIDType},
			SPIFFEMatcherType,
			cel.UnaryBinding(unarySPIFFEMatchExactFnImpl),
		),

		cel.Overload(
			fmt.Sprintf("%s_string", spiffeMatchExactFn),
			[]*cel.Type{cel.StringType},
			SPIFFEMatcherType,
			cel.UnaryBinding(unarySPIFFEMatchExactFnImpl),
		),
	)

	SPIFFEMatchOneOfFunc = cel.Function(spiffeMatchOneOfFn,
		cel.Overload(
			fmt.Sprintf("%s_spiffeIDList", spiffeMatchOneOfFn),
			[]*cel.Type{cel.ListType(SPIFFEIDType)},
			SPIFFEMatcherType,
			cel.UnaryBinding(unarySPIFFEMatchOneOfFnImpl),
		),

		cel.Overload(
			fmt.Sprintf("%s_stringList", spiffeMatchOneOfFn),
			[]*cel.Type{cel.ListType(cel.StringType)},
			SPIFFEMatcherType,
			cel.UnaryBinding(unarySPIFFEMatchOneOfFnImpl),
		),
	)

	SPIFFEMatchTrustDomainFunc = cel.Function(spiffeMatchTrustDomainFn,
		cel.Overload(
			fmt.Sprintf("%s_spiffeTrustDomain", spiffeMatchTrustDomainFn),
			[]*cel.Type{SPIFFETrustDomainType},
			SPIFFEMatcherType,
			cel.UnaryBinding(unarySPIFFEMatchTrustDomainFnImpl),
		),

		cel.Overload(
			fmt.Sprintf("%s_string", spiffeMatchTrustDomainFn),
			[]*cel.Type{cel.StringType},
			SPIFFEMatcherType,
			cel.UnaryBinding(unarySPIFFEMatchTrustDomainFnImpl),
		),
	)

	spiffeIDTypeExpr          = types.NewObjectType(spiffeIDTypeName)
	spiffeMatcherTypeExpr     = types.NewObjectType(spiffeMatcherTypeName)
	spiffeTrustDomainTypeExpr = types.NewObjectType(spiffeTrustDomainTypeName)

	SPIFFEDeclrations = []*decls.FunctionDecl{
		newFunction(overloadSpiffeIDIsMemberOf,
			decls.MemberOverload(overloadSpiffeIDIsMemberOf,
				[]*types.Type{spiffeIDTypeExpr, spiffeTrustDomainTypeExpr},
				types.BoolType,
			),
		),

		newFunction(overloadSpiffeIDPath,
			decls.MemberOverload(overloadSpiffeIDPath,
				[]*types.Type{spiffeIDTypeExpr},
				types.StringType,
			),
		),

		newFunction(overloadSpiffeIDTrustDomain,
			decls.MemberOverload(overloadSpiffeIDTrustDomain,
				[]*types.Type{spiffeIDTypeExpr},
				spiffeTrustDomainTypeExpr,
			),
		),

		newFunction(overloadSpiffeTrustDomainID,
			decls.MemberOverload(overloadSpiffeTrustDomainID,
				[]*types.Type{spiffeTrustDomainTypeExpr},
				types.StringType,
			),
		),

		newFunction(overloadSpiffeTrustDomainName,
			decls.MemberOverload(overloadSpiffeTrustDomainName,
				[]*types.Type{spiffeTrustDomainTypeExpr},
				types.StringType,
			),
		),

		newFunction(overloadSpiffeMatcherMatchesID,
			decls.MemberOverload(fmt.Sprintf("%s_spiffeID", overloadSpiffeMatcherMatchesID),
				[]*types.Type{spiffeMatcherTypeExpr, spiffeIDTypeExpr},
				types.BoolType,
			),

			decls.MemberOverload(fmt.Sprintf("%s_string", overloadSpiffeMatcherMatchesID),
				[]*types.Type{spiffeMatcherTypeExpr, types.StringType},
				types.BoolType,
			),
		),
	}
)

func unarySPIFFEIDFnImpl(v ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

func unarySPIFFETrustDomainFnImpl(v ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func unarySPIFFEMatchAnyFnImpl(args ...ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func unarySPIFFEMatchExactFnImpl(v ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func unarySPIFFEMatchOneOfFnImpl(v ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func convertSPIFFEIDListToMatcher(l traits.Lister) (spiffeid.Matcher, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.Matcher), nil
}

func convertStringListToMatcher(l traits.Lister) (spiffeid.Matcher, error) {
	_ = "STUB: not implemented"
	return *new(spiffeid.Matcher), nil
}

func unarySPIFFEMatchTrustDomainFnImpl(v ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

type SPIFFEID struct {
	id spiffeid.ID
}

// ConvertToNative implements ref.Val.ConvertToNative.
func (sid SPIFFEID) ConvertToNative(typeDesc reflect.Type) (any, error) {
	_ = "STUB: not implemented"
	//nolint:exhaustive
	return *new(any), nil
}

// ConvertToType implements ref.Val.ConvertToType.
func (sid SPIFFEID) ConvertToType(typeVal ref.Type) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Type implements ref.Val.Type.
func (sid SPIFFEID) Type() ref.Type {
	_ = "STUB: not implemented"
	return *

	// Value implements ref.Val.Value.
	new(ref.Type)
}

func (sid SPIFFEID) Value() any {
	_ = "STUB: not implemented"

	// Equal implements ref.Val.Equal.
	return *new(any)
}

func (sid SPIFFEID) Equal(other ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

// Receive implements traits.Receiver.Receive.
func (sid SPIFFEID) Receive(function, _ string, args []ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

func spiffeIDIsMemberOf(s SPIFFEID, arg ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

type SPIFFETrustDomain struct {
	td spiffeid.TrustDomain
}

// ConvertToNative implements ref.Val.ConvertToNative.
func (std SPIFFETrustDomain) ConvertToNative(typeDesc reflect.Type) (any, error) {
	_ = "STUB: not implemented"
	//nolint:exhaustive
	return *new(any), nil
}

// ConvertToType implements ref.Val.ConvertToType.
func (std SPIFFETrustDomain) ConvertToType(typeVal ref.Type) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Type implements ref.Val.Type.
func (std SPIFFETrustDomain) Type() ref.Type { _ = "STUB: not implemented"; return *new(ref.Type) }

// Value implements ref.Val.Value.
func (std SPIFFETrustDomain) Value() any {
	_ = "STUB: not implemented"

	// Equal implements ref.Val.Equal.
	return *new(any)
}

func (std SPIFFETrustDomain) Equal(other ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Compare implements traits.Comparer.
func (std SPIFFETrustDomain) Compare(other ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Receive implements traits.Receiver.Receive.
func (std SPIFFETrustDomain) Receive(function, _ string, args []ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

type SPIFFEMatcher struct {
	matcher spiffeid.Matcher
}

// ConvertToNative implements ref.Val.ConvertToNative.
func (sm SPIFFEMatcher) ConvertToNative(typeDesc reflect.Type) (any, error) {
	_ = "STUB: not implemented"
	//nolint:exhaustive
	return *new(any), nil
}

// ConvertToType implements ref.Val.ConvertToType.
func (sm SPIFFEMatcher) ConvertToType(typeVal ref.Type) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// Type implements ref.Val.Type.
func (sm SPIFFEMatcher) Type() ref.Type {
	_ = "STUB: not implemented"
	return *

	// Value implements ref.Val.Value.
	new(ref.Type)
}

func (sm SPIFFEMatcher) Value() any {
	_ = "STUB: not implemented"

	// Equal implements ref.Val.Equal.
	return *new(any)
}

func (sm SPIFFEMatcher) Equal(_ ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *

	// Receive implements traits.Receiver.Receive.
	new(ref.Val)
}

func (sm SPIFFEMatcher) Receive(function, _ string, args []ref.Val) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}
