// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package files

import (
	"io"
	"io/fs"
)

type Found interface {
	Open() (io.Reader, error)
	ID() string
	Path() string
}

type foundInFs struct {
	fsys fs.FS
	id   string
	path string
}

func (f foundInFs) Open() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (f foundInFs) ID() string { _ = "STUB: not implemented"; return "" }

func (f foundInFs) Path() string { _ = "STUB: not implemented"; return "" }
