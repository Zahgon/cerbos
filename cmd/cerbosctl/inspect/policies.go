// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package inspect

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/inspect/internal/flagset"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const (
	help = `# Inspect policies

cerbosctl inspect policies

# Inspect policies, print no headers

cerbosctl inspect policies`
)

//nolint:govet
type PoliciesCmd struct { //betteralign:ignore
	flagset.Filters
	flagset.Format
}

func (c *PoliciesCmd) Run(k *kong.Kong, cctx *client.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PoliciesCmd) Help() string { _ = "STUB: not implemented"; return "" }
