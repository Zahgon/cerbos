// Copyright 2021-2026 Zenauth Ltd.

package main

import (
	"embed"

	"github.com/alecthomas/kong"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
)

//go:embed templates
var templatesFS embed.FS

const changelogDirName = ".changelog"

type args struct {
	Add      addCmd      `cmd:""`
	Generate generateCmd `cmd:""`
}

type addCmd struct {
	Type        string `help:"Type of the entry" enum:"breaking,chore,deprecation,docs,enhancement,feature,fix" required:""`
	Description string `help:"Change description" required:""`
}

type generateCmd struct {
	From       string `help:"Reference to start of change log" required:""`
	NewVersion string `help:"New release version" required:""`
}

func main() {
	var args args
	ctx := kong.Parse(&args,
		kong.Name("changelog"),
		kong.Description("Manage changelog"),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run())
}

func (ac *addCmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

type ChangelogSection struct {
	Title   string
	Type    string
	Entries []entry
}

type Changelog struct {
	Version  string
	Sections []ChangelogSection
}

func (gc *generateCmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

type entry struct {
	Name        string
	Type        string
	Description string
}

func getChangelogEntries(workingDir string, fromRevision plumbing.Revision) (map[string][]entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTree(repo *git.Repository, revision plumbing.Revision) (*object.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createEntry(fileName string) (entry, error) {
	_ = "STUB: not implemented"
	return *new(entry), nil
}

// Use Asciidoc continuation for complex Descriptions.
