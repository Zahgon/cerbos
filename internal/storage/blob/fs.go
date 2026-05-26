// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package blob

import (
	"io"
	"io/fs"
)

const (
	perm775 = 0o775
)

var _ FS = blobFS{}

// FS represents file system interface that used by the Cloner and Store.
type FS interface {
	fs.StatFS
	Remove(name string) error
	RemoveAll(name string) error
	Rename(name, newName string) error
	Create(name string) (io.WriteCloser, error)
	MkdirAll(path string, perm fs.FileMode) error
}

func newBlobFS(dir string) FS { _ = "STUB: not implemented"; return *new(FS) }

type blobFS struct {
	fsys fs.FS
	dir  string
}

func (s blobFS) Create(name string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (s blobFS) MkdirAll(path string, perm fs.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

func (s blobFS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

func (s blobFS) Remove(name string) error { _ = "STUB: not implemented"; return nil }

func (s blobFS) RemoveAll(name string) error { _ = "STUB: not implemented"; return nil }

func (s blobFS) Rename(name, newName string) error { _ = "STUB: not implemented"; return nil }

func (s blobFS) Stat(name string) (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

func createOrValidateDir(dir string) error { _ = "STUB: not implemented"; return nil }

//nolint:mnd
