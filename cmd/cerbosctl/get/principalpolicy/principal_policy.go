// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package principalpolicy

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const help = `# List principal policies
cerbosctl get principal_policies
cerbosctl get principal_policy
cerbosctl get pp

# List and filter principal policies
cerbosctl get principal_policies --name donald_duck

# List and sort principal policies by column
cerbosctl get principal_policies --sort-by policyId
cerbosctl get principal_policies --sort-by name
cerbosctl get principal_policies --sort-by version

# Get principal policy definition (disk, git, blob stores)
cerbosctl get principal_policies donald_duck.yaml

# Get principal policy definition (mutable stores)
cerbosctl get principal_policies principal.donald_duck.default

# Get principal policy definition as yaml
cerbosctl get principal_policies principal.donald_duck.default -oyaml

# Get principal policy definition as json
cerbosctl get principal_policies principal.donald_duck.default -ojson

# Get principal policy definition as pretty json
cerbosctl get principal_policies principal.donald_duck.default -oprettyjson`

type Cmd struct { //betteralign:ignore
	flagset.Format
	flagset.Sort
	flagset.Filters
}

func (c *Cmd) Run(k *kong.Kong, ctx *client.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
