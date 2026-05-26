// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"io"

	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
)

func printSchema(w io.Writer, schemas []*schemav1.Schema, output flagset.OutputFormat) error {
	_ = "STUB: not implemented"
	return nil
}

func printSchemaPrettyJSON(w io.Writer, schemas []*schemav1.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

func printSchemaJSON(w io.Writer, schemas []*schemav1.Schema) error {
	_ = "STUB: not implemented"
	return nil
}
