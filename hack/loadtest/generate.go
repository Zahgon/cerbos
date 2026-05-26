// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build loadtest

package main

import (
	"embed"
	"path/filepath"
	"text/template"

	"github.com/alecthomas/kong"
)

const (
	filesDir     = "files"
	policiesDir  = "policies"
	requestsDir  = "requests"
	schemasDir   = "schemas"
	templatesDir = "templates"
)

//go:embed templates
var fsys embed.FS

var tmplOutConf = map[string]string{
	policiesDir: policiesDir,
	schemasDir:  filepath.Join(policiesDir, "_schemas"),
	requestsDir: requestsDir,
}

type cmd struct {
	Out   string `default:"work" help:"Directory to output the generated files" type:"path"`
	Count int    `default:"100" help:"Number of copies to generate from each template"`
	Set   string `default:"classic" help:"Policy template set to use (classic, multitenant)"`
}

type templateArgs struct {
	RequestID string
	N         int
}

func (ta templateArgs) NameMod(n string) string { _ = "STUB: not implemented"; return "" }

type renderFunc func(templateArgs) error

func main() {
	ctx := kong.Parse(&cmd{},
		kong.Description("Generate load test data"),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run())
}

func (c *cmd) Run() error { _ = "STUB: not implemented"; return nil }

func prepOutDirs(out string) error { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func dirExistsInEmbed(dir string) bool { _ = "STUB: not implemented"; return false }

// copyStaticFiles copies files from templates/<set>/files/ to the output directory.
// It is a no-op if the files/ directory does not exist for the given set.
func copyStaticFiles(set, out string) error { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func copyEmbedFile(srcPath, destPath string) error { _ = "STUB: not implemented"; return nil }

func createRenderer(tmplDir, outDir string) (renderFunc, error) {
	_ = "STUB: not implemented"
	return *new(renderFunc), nil
}

func mkRenderFunc(out string, tmpl *template.Template) renderFunc {
	_ = "STUB: not implemented"
	return *new(renderFunc)
}

func renderFile(fileName string, tmpl *template.Template, args templateArgs) error {
	_ = "STUB: not implemented"
	return nil
}
