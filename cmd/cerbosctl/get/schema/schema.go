// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const help = `# List schemas
cerbosctl get schemas
cerbosctl get schema
cerbosctl get s

# Get schema definition
cerbosctl get schemas principal.json`

type Cmd struct { //betteralign:ignore
	flagset.Format

	SchemaIDs []string `arg:"" name:"id" optional:"" help:"list of schema ids to retrieve"` //nolint:revive
}

func (c *Cmd) Run(k *kong.Kong, ctx *client.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
