// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package jsonschema

import (
	"errors"
	"io/fs"
	"log"

	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/cerbos/cerbos/schema"
)

var (
	ErrEmptyFile = errors.New("empty file")

	testSchema              *jsonschema.Schema
	principalFixturesSchema *jsonschema.Schema
	resourceFixturesSchema  *jsonschema.Schema
	auxDataFixturesSchema   *jsonschema.Schema
)

func init() {
	var err error

	if testSchema, err = jsonschema.CompileString("TestSuite.schema.json", schema.TestSuiteJSONSchema); err != nil {
		log.Fatalf("failed to compile test schema: %v", err)
	}

	if principalFixturesSchema, err = jsonschema.CompileString("Principals.schema.json", schema.PrincipalFixturesJSONSchema); err != nil {
		log.Fatalf("failed to compile principal fixtures schema: %v", err)
	}

	if resourceFixturesSchema, err = jsonschema.CompileString("Resources.schema.json", schema.ResourceFixturesJSONSchema); err != nil {
		log.Fatalf("failed to compile resource fixtures schema: %v", err)
	}

	if auxDataFixturesSchema, err = jsonschema.CompileString("AuxData.schema.json", schema.AuxDataFixturesJSONSchema); err != nil {
		log.Fatalf("failed to compile aux data fixtures schema: %v", err)
	}
}

// ValidateTest validates the test in the fsys with the JSON schema.
func ValidateTest(fsys fs.FS, path string) error { _ = "STUB: not implemented"; return nil }

// ValidatePrincipalFixtures validates the principal fixtures file in the fsys with the JSON schema.
func ValidatePrincipalFixtures(fsys fs.FS, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateResourceFixtures validates the resource fixtures file in the fsys with the JSON schema.
func ValidateResourceFixtures(fsys fs.FS, path string) error { _ = "STUB: not implemented"; return nil }

// ValidatePrincipalFixtures validates the aux data fixtures file in the fsys with the JSON schema.
func ValidateAuxDataFixtures(fsys fs.FS, path string) error { _ = "STUB: not implemented"; return nil }

func validate(s *jsonschema.Schema, fsys fs.FS, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func newValidationError(ve *jsonschema.ValidationError) error {
	_ = "STUB: not implemented"
	return nil
}

func sortIssues(m map[string]struct{}) string { _ = "STUB: not implemented"; return "" }
