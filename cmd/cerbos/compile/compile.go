// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package compile

import (
	"io/fs"

	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbos/compile/internal/flagset"
	"github.com/cerbos/cerbos/internal/outputcolor"
)

const (
	help = `
Examples:

# Compile and run tests found in /path/to/policy/repo

cerbos compile /path/to/policy/repo

# Compile but skip tests

cerbos compile --skip-tests /path/to/policy/repo

# Compile and run tests matching a filter (globs for suite, test, principal, resource, action)

cerbos compile --test-filter='suite=MySuite;test=album*;principal=alice;resource=my_album;action=view' /path/to/policy/repo

# Multiple filters can be combined (all filter dimensions are merged)

cerbos compile --test-filter='principal=alice,bob' --test-filter='action=view,edit' /path/to/policy/repo
`
)

//nolint:govet // Kong prints fields in order, so we don't want to reorder fields to save bytes.
type Cmd struct { //betteralign:ignore
	Dir           string                            `help:"Policy directory" arg:"" required:"" type:"path"`
	IgnoreSchemas bool                              `help:"Ignore schemas during compilation"`
	Tests         string                            `help:"[Deprecated] Path to the directory containing tests. Defaults to policy directory." type:"path"`
	RunRegexp     string                            `help:"[Deprecated] Run only tests that match this regex" name:"run" hidden:""`
	TestFilter    flagset.TestFilter                `help:"Filter tests by dimensions (suite, test, principal, resource, action). Format: 'dimension=glob1,glob2;...'. Can be specified multiple times." name:"test-filter"`
	SkipTests     bool                              `help:"Skip tests"`
	SkipBatching  bool                              `help:"Skip batching tests"`
	Output        flagset.OutputFormat              `help:"Output format (${enum})" default:"tree" enum:"tree,list,json" short:"o"`
	TestOutput    *flagset.VerificationOutputFormat `help:"Test output format. If unspecified matches the value of the output flag. (tree,list,json,junit)"`
	Color         *outputcolor.Level                `help:"Output color level (auto,never,always,256,16m). Defaults to auto." xor:"color"`
	NoColor       bool                              `help:"Disable colored output" xor:"color"`
	Verbose       bool                              `help:"Verbose output on test failure"`
}

func (c *Cmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

//nolint:nestif

func (c *Cmd) testsDir() (fs.FS, string, error) {
	_ = "STUB: not implemented"
	return *new(fs.FS), "", nil
}

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
