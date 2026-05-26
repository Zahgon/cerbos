// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

import (

	// signal to the compiler that files need to be embedded.
	_ "embed"
	"net/http"
)

//go:embed openapiv2/cerbos/svc/v1/svc.swagger.json
var svcSwaggerRaw []byte

//go:embed assets/ui.html
var rapidocHTML []byte

//go:embed jsonschema/cerbos/policy/v1/TestSuite.schema.json
var TestSuiteJSONSchema string

//go:embed jsonschema/cerbos/policy/v1/TestFixture/Principals.schema.json
var PrincipalFixturesJSONSchema string

//go:embed jsonschema/cerbos/policy/v1/TestFixture/Resources.schema.json
var ResourceFixturesJSONSchema string

//go:embed jsonschema/cerbos/policy/v1/TestFixture/AuxData.schema.json
var AuxDataFixturesJSONSchema string

func ServeSvcSwagger(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

//nolint:gosec

func ServeUI(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func cleanup(r *http.Request) { _ = "STUB: not implemented"; return }

type swaggerMod struct {
	err    error
	schema []byte
}

func newSwaggerMod() *swaggerMod { _ = "STUB: not implemented"; return nil }

func (sm *swaggerMod) setVersion(version string) *swaggerMod { _ = "STUB: not implemented"; return nil }

func (sm *swaggerMod) setHost(host string) *swaggerMod { _ = "STUB: not implemented"; return nil }

func (sm *swaggerMod) setScheme(scheme string) *swaggerMod { _ = "STUB: not implemented"; return nil }

func (sm *swaggerMod) build() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
