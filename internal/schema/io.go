// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"io"
	"io/fs"

	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
)

func ReadSchemaFromFile(fsys fs.FS, path string) (*schemav1.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadSchema reads a schema from the given reader.
func ReadSchema(src io.Reader, id string) (*schemav1.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
