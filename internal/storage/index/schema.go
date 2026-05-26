// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package index

import (
	"context"
	"io"
	"io/fs"
)

type SchemaLoader struct {
	err  error
	fsys fs.FS
}

func NewSchemaLoader(fsys fs.FS, rootDir string) *SchemaLoader {
	_ = "STUB: not implemented"
	return nil
}

func (sl *SchemaLoader) ListIDs(_ context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sl *SchemaLoader) Load(_ context.Context, id string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
