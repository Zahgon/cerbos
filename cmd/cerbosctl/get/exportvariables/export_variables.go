// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package exportvariables

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const help = `# List exported variables
cerbosctl get export_variables
cerbosctl get ev

# List and filter exported variables
cerbosctl get export_variables --name my_variables

# List and sort exported variables by column
cerbosctl get export_variables --sort-by policyId
cerbosctl get export_variables --sort-by name

# List all policies including disabled policies
cerbosctl get export_variables --include-disabled

# Get exported variables policy definition (disk, git, blob stores)
cerbosctl get export_variables my_variables.yaml

# Get exported variables policy definition (mutable stores)
cerbosctl get export_variables export_variables.my_variables

# Get exported variables policy definition as yaml
cerbosctl get export_variables export_variables.my_variables -oyaml

# Get exported variables policy definition as json
cerbosctl get export_variables export_variables.my_variables -ojson

# Get exported variables policy definition as pretty json
cerbosctl get export_variables export_variables.my_variables -oprettyjson`

type Cmd struct { //betteralign:ignore
	flagset.Format
	flagset.Sort
	flagset.Filters
}

func (c *Cmd) Run(k *kong.Kong, ctx *client.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
