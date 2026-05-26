// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package derivedroles

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const help = `# List derived roles
cerbosctl get derived_roles
cerbosctl get derived_role
cerbosctl get dr

# List and filter derived roles
cerbosctl get derived_roles --name my_derived_roles

# List and sort derived roles by column
cerbosctl get derived_roles --sort-by policyId
cerbosctl get derived_roles --sort-by name

# List all policies including disabled policies
cerbosctl get derived_roles --include-disabled
cerbosctl get principal_policy --include-disabled
cerbosctl get resource_policy --include-disabled

# Get derived role policy definition (disk, git, blob stores)
cerbosctl get derived_roles blog_derived_roles.yaml

# Get derived role policy definition (mutable stores)
cerbosctl get derived_roles derived_roles.my_derived_roles

# Get derived role policy definition as yaml
cerbosctl get derived_roles derived_roles.my_derived_roles -oyaml

# Get derived role policy definition as json
cerbosctl get derived_roles derived_roles.my_derived_roles -ojson

# Get derived role policy definition as pretty json
cerbosctl get derived_roles derived_roles.my_derived_roles -oprettyjson`

type Cmd struct { //betteralign:ignore
	flagset.Format
	flagset.Sort
	flagset.Filters
}

func (c *Cmd) Run(k *kong.Kong, ctx *client.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
