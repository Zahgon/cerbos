// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"text/template"

	"github.com/cerbos/cerbos/hack/tools/generate-npm-packages/templates"
)

var (
	binaries = []string{
		"cerbos",
		"cerbosctl",
	}

	platforms = []templates.Platform{
		{OS: "darwin", Arch: "arm64"},
		{OS: "darwin", Arch: "x64"},
		{OS: "linux", Arch: "arm64"},
		{OS: "linux", Arch: "x64"},
	}
)

const (
	antoraConfigFile = "docs/antora.yml"
	outDir           = "npm/packages"
	dirPerm          = 0o755
)

type AntoraConfig struct {
	Version    string
	Prerelease string
}

func main() {
	err := generatePackages()
	if err != nil {
		log.Fatalln(err)
	}
}

func generatePackages() error { _ = "STUB: not implemented"; return nil }

func readVersion() (string, error) { _ = "STUB: not implemented"; return "", nil }

func generatePackage(name string, templates *template.Template, data any) error {
	_ = "STUB: not implemented"
	return nil
}

func writeFile(packageDir string, template *template.Template, data any) (err error) {
	_ = "STUB: not implemented"
	return nil
}
