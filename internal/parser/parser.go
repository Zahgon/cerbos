// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"errors"
	"fmt"
	"io"
	"regexp"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/go/protovalidate"
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/printer"
	"github.com/goccy/go-yaml/token"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	sourcev1 "github.com/cerbos/cerbos/api/genpb/cerbos/source/v1"
)

const (
	base10    = 10
	bitSize32 = 32
	bitSize64 = 64
)

var ErrNotFound = errors.New("not found")

type PanicError struct {
	Cause   any
	Context []byte
}

func (pe PanicError) Error() string { _ = "STUB: not implemented"; return "" }

var protoErrPrefix = regexp.MustCompile(`proto:(\x{00a0}|\x{0020})+\(line\s+\d+:\d+\):\s*`)

// Find a single document from the multi-document stream.
// TODO(cell): Optimize!
// For our use case, this could be optimized by storing the offset of each document and directly seeking to that offset.
// However, there are a couple of problems with that:
//  1. The offsets reported by the parser are not always reliable (I am yet to figure out why)
//  2. If YAML anchors have been used, we need to resolve those first by reading through the entire file anyway
//     However, this is a relatively niche case and we can handle that case lazily (seek first, read, and resolve anchors only if they exist in the doc)
//
// In the interest of time, I am leaving those optimizations for later.
func Find[T proto.Message](r io.Reader, match func(T) bool, out T, opts ...UnmarshalOpt) (SourceCtx, error) {
	_ = "STUB: not implemented"
	return *new(SourceCtx), nil
}

// Ignore documents not structured as policies

func Unmarshal[T proto.Message](r io.Reader, factory func() T, opts ...UnmarshalOpt) ([]T, []SourceCtx, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func UnmarshalBytes[T proto.Message](contents []byte, factory func() T, opts ...UnmarshalOpt) (_ []T, _ []SourceCtx, outErr error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ignore commented out documents

// If given an invalid file with multiple lines of text, the parser generates a "doc" for each line.
// Ignore a consecutive run of such docs.

func parse(contents []byte, detectProblems bool) (_ *ast.File, outErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nestif

func detectStringStartingWithQuote(tokens token.Tokens) (outErrs []*sourcev1.Error) {
	_ = "STUB: not implemented"
	return nil
}

// Check whether we are inside a flow block (e.g. foo: {"x": "y"})

type unmarshalOpts struct {
	validator           protovalidate.Validator
	ignoreUnknownFields bool
}

type UnmarshalOpt func(*unmarshalOpts)

// WithIgnoreUnknownFields ignores unknown fields not defined in the protobuf schema.
func WithIgnoreUnknownFields() UnmarshalOpt { _ = "STUB: not implemented"; return *new(UnmarshalOpt) }

// WithValidate validates the unmarshaled message using protovalidate.
func WithValidator(validator protovalidate.Validator) UnmarshalOpt {
	_ = "STUB: not implemented"
	return *new(UnmarshalOpt)
}

type unmarshaler[T proto.Message] struct {
	anchors map[string]ast.Node
	unmarshalOpts
}

func (u *unmarshaler[T]) unmarshalMapping(uctx *unmarshalCtx, v ast.MapNode, out protoreflect.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// Find all the merge keys first and populate the message because they need to overwritten by new values
// Basically, ensuring that "bar" is set to "baz" in the following case regardless of what value "bar" has
// in the "anchor" map.
// x:
//  bar: baz
//  <<: *anchor

// already handled above

func (u *unmarshaler[T]) resolveMerge(uctx *unmarshalCtx, n ast.Node) (ast.MapNode, error) {
	_ = "STUB: not implemented"
	return *new(ast.MapNode), nil
}

func (u *unmarshaler[T]) resolveNode(uctx *unmarshalCtx, n ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

// adapted from https://github.com/goccy/go-yaml/blob/31fe1baacec127337140701face2e64a356075fd/decode.go#L355
func (u *unmarshaler[T]) resolveAlias(uctx *unmarshalCtx, n ast.Node) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

//nolint:nestif

//nolint:mnd

func (u *unmarshaler[T]) unmarshalList(uctx *unmarshalCtx, n ast.Node, fd protoreflect.FieldDescriptor, list protoreflect.List) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaler[T]) unmarshalMap(uctx *unmarshalCtx, n ast.MapNode, fd protoreflect.FieldDescriptor, mmap protoreflect.Map, overwriteKeys bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaler[T]) unmarshalSingular(uctx *unmarshalCtx, n ast.Node, fd protoreflect.FieldDescriptor, out protoreflect.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaler[T]) unmarshalScalar(uctx *unmarshalCtx, n ast.Node, fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

func (u *unmarshaler[T]) unmarshalBool(uctx *unmarshalCtx, n ast.Node) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

func (u *unmarshaler[T]) unmarshalEnum(uctx *unmarshalCtx, n ast.Node, fd protoreflect.FieldDescriptor) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

func (u *unmarshaler[T]) unmarshalString(uctx *unmarshalCtx, n ast.Node) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

//nolint:dupl
func (u *unmarshaler[T]) unmarshalInt(uctx *unmarshalCtx, n ast.Node, bitSize int) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

//nolint:dupl
func (u *unmarshaler[T]) unmarshalUint(uctx *unmarshalCtx, n ast.Node, bitSize int) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

func (u *unmarshaler[T]) unmarshalFloat(uctx *unmarshalCtx, n ast.Node, bitSize int) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

func (u *unmarshaler[T]) unmarshalBytes(uctx *unmarshalCtx, n ast.Node) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

func (u *unmarshaler[T]) unmarshalMessage(uctx *unmarshalCtx, n ast.Node, out protoreflect.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *unmarshaler[T]) unmarshalWKT(uctx *unmarshalCtx, n ast.Node, out protoreflect.Message) error {
	_ = "STUB: not implemented"
	// Google's well-known type handling is hidden inside an internal package and can't be used directly.
	// The code is complicated and copying it here is probably not a good idea.
	// Cerbos policies don't use any WKTs. Only the test suite definitions and fixtures use them so
	// resorting to this slightly inefficient hack is OK for now.
	return nil
}

func (u *unmarshaler[T]) unmarshalMapKey(uctx *unmarshalCtx, n ast.Node, fd protoreflect.FieldDescriptor) (protoreflect.MapKey, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.MapKey), nil
}

func (u *unmarshaler[T]) validate(uctx *unmarshalCtx, msg T) (outErr error) {
	_ = "STUB: not implemented"
	return nil
}

// Taken from https://github.com/bufbuild/protovalidate-go/blob/46121307d89af5b7ae07e27a58d1c2ac26845388/internal/errors/utils.go#L133
func fieldPathString(path []*validate.FieldPathElement) string {
	_ = "STUB: not implemented"
	return ""
}

func pos(n ast.Node) string { _ = "STUB: not implemented"; return "" }

type unmarshalCtx struct {
	errPrinter printer.Printer
	srcCtx     *sourcev1.SourceContext
	doc        *ast.DocumentNode
	protoPath  string
}

func newUnmarshalCtx(doc *ast.DocumentNode) *unmarshalCtx { _ = "STUB: not implemented"; return nil }

func (uc *unmarshalCtx) forPath(path string) *unmarshalCtx { _ = "STUB: not implemented"; return nil }

func (uc *unmarshalCtx) forField(fd protoreflect.FieldDescriptor, n ast.Node) *unmarshalCtx {
	_ = "STUB: not implemented"
	return nil
}

func (uc *unmarshalCtx) forListItem(i int, n ast.Node) *unmarshalCtx {
	_ = "STUB: not implemented"
	return nil
}

func (uc *unmarshalCtx) forMapItem(key string, keyNode ast.MapKeyNode, valueNode ast.Node) *unmarshalCtx {
	_ = "STUB: not implemented"
	return nil
}

func (uc *unmarshalCtx) recordMapKeyPosition(path string, n ast.Node) {
	_ = "STUB: not implemented"
	return
}

func (uc *unmarshalCtx) recordFieldPosition(path string, n ast.Node) {
	_ = "STUB: not implemented"
	return
}

func nodePosition(n ast.Node) *sourcev1.Position { _ = "STUB: not implemented"; return nil }

func (uc *unmarshalCtx) addError(err *sourcev1.Error) { _ = "STUB: not implemented"; return }

func (uc *unmarshalCtx) perrorf(n ast.Node, msg string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (uc *unmarshalCtx) verrorf(path, msg string) error { _ = "STUB: not implemented"; return nil }

func (uc *unmarshalCtx) buildErrContext(path string) string { _ = "STUB: not implemented"; return "" }

func (uc *unmarshalCtx) toSourceCtx() SourceCtx { _ = "STUB: not implemented"; return *new(SourceCtx) }

type UnmarshalError struct {
	Err *sourcev1.Error
}

func NewUnmarshalError(err *sourcev1.Error) UnmarshalError {
	_ = "STUB: not implemented"
	return *new(UnmarshalError)
}

func (ue UnmarshalError) Error() string { _ = "STUB: not implemented"; return "" }

func (ue UnmarshalError) StringWithoutContext() string { _ = "STUB: not implemented"; return "" }

func (ue UnmarshalError) Format(state fmt.State, verb rune) { _ = "STUB: not implemented"; return }
