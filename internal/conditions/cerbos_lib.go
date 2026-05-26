// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package conditions

import (
	"context"
	"time"

	"github.com/google/cel-go/cel"
	celast "github.com/google/cel-go/common/ast"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"github.com/google/cel-go/interpreter"
	"github.com/google/cel-go/interpreter/functions"
)

const (
	exceptFn                    = "except"
	hasIntersectionFnDeprecated = "has_intersection"
	hasIntersectionFn           = "hasIntersection"
	inIPAddrRangeFn             = "inIPAddrRange"
	intersectFn                 = "intersect"
	isSubsetFnDeprecated        = "is_subset"
	isSubsetFn                  = "isSubset"
	nowFn                       = "now"
	timeSinceFn                 = "timeSince"
	IDFn                        = "id"
	basePathFn                  = "basePath"
	dirPathFn                   = "dirPath"
	extPathFn                   = "extPath"
	joinPathFn                  = "joinPath"
	pathHasPrefixFn             = "pathHasPrefix"
	pathMatchFn                 = "pathMatch"
	pathMatchAnyOfFn            = "pathMatchAnyOf"
	relPathFn                   = "relPath"
	volumeNameFn                = "volumeName"
	CELNowFnActivationKey       = "_cerbos_now_fn"
)

// CerbosCELLib returns the custom CEL functions provided by Cerbos.
func CerbosCELLib() cel.EnvOption { _ = "STUB: not implemented"; return *new(cel.EnvOption) }

type cerbosLib struct{}

func (clib cerbosLib) CompileOptions() []cel.EnvOption { _ = "STUB: not implemented"; return nil }

// options for set operations like intersect and except

// options for set checks like isIntersection and isSubset

func (clib cerbosLib) ProgramOptions() []cel.ProgramOption { _ = "STUB: not implemented"; return nil }

type NowFunc = func() time.Time

// Now returns a NowFunc that always returns the time at which Now was called.
func Now() NowFunc { _ = "STUB: not implemented"; return *new(NowFunc) }

// ContextEval returns the result of an evaluation of the ast and environment against the input vars,
// providing time-based functions with a static definition of the current time.
//
// The given nowFunc must return the same timestamp each time it is called.
//
// See https://pkg.go.dev/github.com/google/cel-go/cel#Program.ContextEval.
func ContextEval(ctx context.Context, env *cel.Env, ast *celast.AST, vars any, nowFunc NowFunc, opts ...cel.ProgramOption) (ref.Val, *cel.EvalDetails, error) {
	_ = "STUB: not implemented"
	return *new(ref.Val), nil, nil
}

func CacheFriendlyTimeDecorator() interpreter.InterpretableDecorator {
	_ = "STUB: not implemented"
	return *new(interpreter.InterpretableDecorator)
}

func cacheFriendlyTimeDecorate(in interpreter.Interpretable) (interpreter.Interpretable, error) {
	_ = "STUB: not implemented"
	return *new(interpreter.Interpretable), nil
}

// nowInterp is a custom Interpretable that looks up NowFunc from the activation at eval time.
type nowInterp struct {
	id int64
}

func (n *nowInterp) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (n *nowInterp) Eval(activation interpreter.Activation) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// timeSinceInterp is a custom Interpretable that looks up NowFunc from the activation at eval time.
type timeSinceInterp struct {
	arg interpreter.Interpretable
	id  int64
}

func (t *timeSinceInterp) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (t *timeSinceInterp) Eval(activation interpreter.Activation) ref.Val {
	_ = "STUB: not implemented"
	return *new(ref.Val)
}

// newTimeDecorator creates a decorator that bakes in the nowFunc at compile time.
// This is used for backward compatibility with ContextEval.
func newTimeDecorator(nowFunc NowFunc) interpreter.InterpretableDecorator {
	_ = "STUB: not implemented"
	return *new(interpreter.InterpretableDecorator)
}

type timeDecorator struct {
	nowFunc NowFunc
}

func (t *timeDecorator) decorate(in interpreter.Interpretable) (interpreter.Interpretable, error) {
	_ = "STUB: not implemented"
	return *new(interpreter.Interpretable), nil
}

// hashable checks whether the type is hashable, i.e. can be used in a Go map.
func hashable(t ref.Type) bool { _ = "STUB: not implemented"; return false }

// exceptList implements difference lhs-rhs returning
// items in lhs (list) that are not members of rhs (list).
func exceptList(lhs, rhs ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

// isSubset returns true value if lhs (list) is a subset of rhs (list).
func isSubset(lhs, rhs ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

func find(i traits.Iterator, item ref.Val) bool { _ = "STUB: not implemented"; return false }

const minListLengthToConvert = 3

func convertToMap(b traits.Lister) map[ref.Val]struct{} { _ = "STUB: not implemented"; return nil }

func hasIntersection(lhs, rhs ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

//nolint:forcetypeassert

// b is the longest list

func intersect(lhs, rhs ref.Val) ref.Val { _ = "STUB: not implemented"; return *new(ref.Val) }

//nolint:forcetypeassert

// b is the longest list

func (clib cerbosLib) inIPAddrRangeFunc(ipAddrVal, cidrVal string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func pathHasPrefix(path, prefix string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func pathMatchAnyOf(path string, patterns ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func callInStringStringOutBool(fn func(string, string) (bool, error)) functions.BinaryOp {
	_ = "STUB: not implemented"
	return *new(functions.BinaryOp)
}

//nolint:govet

func callInTimestampOutDuration(fn func(time.Time) time.Duration) functions.UnaryOp {
	_ = "STUB: not implemented"
	return *new(functions.UnaryOp)
}

func callInNothingOutTimestamp(fn func() time.Time) functions.FunctionOp {
	_ = "STUB: not implemented"
	return *new(functions.FunctionOp)
}

func callInStringOutString(fn func(string) string) functions.UnaryOp {
	_ = "STUB: not implemented"
	return *new(functions.UnaryOp)
}

func callInStringOutStringErr(fn func(string) (string, error)) functions.UnaryOp {
	_ = "STUB: not implemented"
	return *new(functions.UnaryOp)
}

//nolint:govet

func callInStringStringOutStringErr(fn func(string, string) (string, error)) functions.BinaryOp {
	_ = "STUB: not implemented"
	return *new(functions.BinaryOp)
}

//nolint:govet

func callInStringSliceOutStringErr(fn func(...string) (string, error)) functions.UnaryOp {
	_ = "STUB: not implemented"
	return *new(functions.UnaryOp)
}

//nolint:govet

func callInStringStringSliceOutBoolErr(fn func(string, ...string) (bool, error)) functions.BinaryOp {
	_ = "STUB: not implemented"
	return *new(functions.BinaryOp)
}

//nolint:govet
