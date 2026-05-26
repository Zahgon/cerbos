// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"errors"
	"io"
	"io/fs"
	"path/filepath"
)

var supportedFileTypes = map[string]struct{}{".yaml": {}, ".yml": {}, ".json": {}}

var ErrNoMatchingFiles = errors.New("no matching files")

// SchemasDirectory is the name of the special directory containing schemas. It is defined here to avoid an import loop.
const SchemasDirectory = "_schemas"

// TestDataDirectory is the name of the special directory containing test fixtures. It is defined here to avoid an import loop.
const TestDataDirectory = "testdata"

const PathSeparator = string(filepath.Separator)

// IsSupportedTestFile return true if the given file is a supported test file name, i.e. "*_test.{yaml,yml,json}".
func IsSupportedTestFile(fileName string) bool { _ = "STUB: not implemented"; return false }

// IsSupportedFileTypeExt returns true and a file extension if the given file has a supported file extension.
func IsSupportedFileTypeExt(fileName string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// IsJSONFileTypeExt returns true if the given file has a json file extension.
func IsJSONFileTypeExt(fileName string) bool { _ = "STUB: not implemented"; return false }

// IsSupportedFileType returns true if the given file has a supported file extension.
func IsSupportedFileType(fileName string) bool { _ = "STUB: not implemented"; return false }

func IsHidden(fileName string) bool { _ = "STUB: not implemented"; return false }

func PathIsHidden(path string) bool { _ = "STUB: not implemented"; return false }

func IsZip(fileName string) bool { _ = "STUB: not implemented"; return false }

func IsTar(fileName string) bool { _ = "STUB: not implemented"; return false }

func IsGzip(fileName string) bool { _ = "STUB: not implemented"; return false }

func IsArchiveFile(fileName string) bool { _ = "STUB: not implemented"; return false }

type ClosableFS struct {
	fs.FS
	io.Closer
	closers []io.Closer
}

func (cfs ClosableFS) Close() (outErr error) { _ = "STUB: not implemented"; return nil }

// GetOneOfSupportedFileNames attempts to retrieve a fileName adding supported extensions.
func GetOneOfSupportedFileNames(fsys fs.FS, fileName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type IndexedFileType uint8

const (
	FileTypeNotIndexed IndexedFileType = iota
	FileTypePolicy
	FileTypeSchema
)

// FileType categorizes the given path according to how it will be treated by the index.
// The path must be "/"-separated and relative to the root policies directory.
func FileType(path string) IndexedFileType { _ = "STUB: not implemented"; return *new(IndexedFileType) }

// RelativeSchemaPath returns the given path within the top-level schemas directory,
// and a flag to indicate whether the path was actually contained in that directory.
// The path must be "/"-separated and relative to the root policies directory.
func RelativeSchemaPath(path string) (string, bool) { _ = "STUB: not implemented"; return "", false }
