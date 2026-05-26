// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"io/fs"
	"testing"
	"text/template"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

func RenderTemplate(tb testing.TB, path string, data any) []byte {
	_ = "STUB: not implemented"
	return nil
}

func TemplateFuncs() template.FuncMap { _ = "STUB: not implemented"; return *new(template.FuncMap) }

type templateHelper struct {
	fsys fs.FS
}

func (th *templateHelper) FileBytes(relPath string) []byte { _ = "STUB: not implemented"; return nil }

func (th *templateHelper) FileString(relPath string) string { _ = "STUB: not implemented"; return "" }

func (th *templateHelper) ToPolicyJSON(p *policyv1.Policy) string {
	_ = "STUB: not implemented"
	return ""
}

func (th *templateHelper) ReadPolicy(relPath string) *policyv1.Policy {
	_ = "STUB: not implemented"
	return nil
}
