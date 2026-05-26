// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package templates

import (
	"embed"
	"text/template"
)

func parse(fsys embed.FS) *template.Template { _ = "STUB: not implemented"; return nil }

var (
	//go:embed binary
	binaryPackageFiles embed.FS

	//go:embed wrapper
	wrapperPackageFiles embed.FS

	BinaryPackage  = parse(binaryPackageFiles)
	WrapperPackage = parse(wrapperPackageFiles)
)

type Platform struct {
	OS   string
	Arch string
}

func (p Platform) String() string { _ = "STUB: not implemented"; return "" }

type BinaryPackageData struct {
	Platform
	Name    string
	Binary  string
	Version string
}

type WrapperPackageData struct {
	Name      string
	Version   string
	Platforms []Platform
}

func (w WrapperPackageData) OptionalDependencies() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (w WrapperPackageData) SupportedPlatforms() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func toJSON(value any, prefix string) (string, error) { _ = "STUB: not implemented"; return "", nil }
