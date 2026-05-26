// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package compile

import (
	"context"
	"io/fs"

	"google.golang.org/protobuf/types/known/structpb"

	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage/index"
)

type Index = index.Index

type Artefact struct {
	Error      error
	PolicySet  *runtimev1.RunnablePolicySet
	SourceFile string
}

type PanicError struct {
	Cause   any
	Context []byte
}

func (pe PanicError) Error() string { _ = "STUB: not implemented"; return "" }

type Errors struct {
	*runtimev1.Errors
}

func (e *Errors) Error() string { _ = "STUB: not implemented"; return "" }

type SourceAttribute struct {
	Value *structpb.Value
	Key   string
}

type SchemaLoader = schema.Loader

type SchemaResolver = schema.Resolver

type SchemaResolverMaker func(SchemaLoader) SchemaResolver

func BuildIndex(ctx context.Context, fsys fs.FS, attrs ...SourceAttribute) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func Files(ctx context.Context, fsys fs.FS, schemaResolverMaker SchemaResolverMaker, attrs ...SourceAttribute) (Index, <-chan Artefact, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil, nil
}
