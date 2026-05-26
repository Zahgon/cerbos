// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package util

import (
	"io"
	"io/fs"

	"google.golang.org/protobuf/proto"
)

// LoadFromJSONOrYAML reads a JSON or YAML encoded protobuf from the given path.
func LoadFromJSONOrYAML(fsys fs.FS, path string, dest proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// OpenDirectoryFS attempts to open a directory FS at the given location. It'll initially check if the target file is an archive,
// and if so, will return the appropriate type which implements the fs.FS interface.
func OpenDirectoryFS(path string) (fs.FS, error) {
	_ = "STUB: not implemented"
	// We don't use `switch filepath.Ext(path)` here because it only suffixes from the final `.`, so `.tar.gz` won't be
	// correctly handled
	return *new(fs.FS), nil
}

func getFsFromTar(r io.Reader, closers ...io.Closer) (fs.FS, error) {
	_ = "STUB: not implemented"
	return *new(fs.FS), nil
}
