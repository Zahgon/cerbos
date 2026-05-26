// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package epdp

import (
	"context"
	"io/fs"

	"github.com/alecthomas/kong"
)

const listCandidatesCmdHelp = `# List candidates for inclusion in embedded PDP

cerbosctl hub epdp list-candidates ./path/to/repo

cerbosctl hub epdp lc ./path/to/repo

# List candidates, print no headers

cerbosctl hub epdp list-candidates ./path/to/repo --no-headers

cerbosctl hub epdp lc ./path/to/repo --no-headers`

const (
	embeddedPDPKey      = "hub.cerbos.cloud/embedded-pdp"
	policyIDNotFoundYet = ""
)

type ListCandidatesCmd struct { //betteralign:ignore
	Path      string `arg:"" type:"path" help:"Path to repository"`
	NoHeaders bool   `help:"Do not output headers"`
}

func (c *ListCandidatesCmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

func (c *ListCandidatesCmd) Help() string { _ = "STUB: not implemented"; return "" }

func listCandidates(ctx context.Context, fsys fs.FS) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nestif
