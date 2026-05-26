// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/alecthomas/kong"
	"github.com/gobwas/glob"
)

type splitCmd struct {
	IgnoreFile string `type:"existingfile" optional:""`
	Kind       string
	Index      int
	Total      int
}

func (cmd *splitCmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

func readTestTimes() (testTimesByKind, error) {
	_ = "STUB: not implemented"
	return *new(testTimesByKind), nil
}

type packageSet map[string]struct{}

func (pkgs packageSet) Packages() []string { _ = "STUB: not implemented"; return nil }

func listPackages(ignoreFile string) (packageSet, error) {
	_ = "STUB: not implemented"
	return *new(packageSet), nil
}

type globList []glob.Glob

func (gl globList) Matches(value string) bool { _ = "STUB: not implemented"; return false }

func loadIgnoreList(ignoreFile string) (globList, error) {
	_ = "STUB: not implemented"
	return *new(globList), nil
}
