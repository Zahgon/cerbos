// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package del

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const policyCmdHelp = `# Delete policies

cerbosctl delete policies derived_roles.my_derived_roles

cerbosctl delete policy derived_roles.my_derived_roles

cerbosctl delete p derived_roles.my_derived_roles

# Delete multiple policies

cerbosctl delete policies derived_roles.my_derived_roles resource.leave_request.default

cerbosctl delete policy derived_roles.my_derived_roles resource.leave_request.default

cerbosctl delete p derived_roles.my_derived_roles resource.leave_request.default`

type PolicyCmd struct { //betteralign:ignore
	PolicyIds []string `arg:"" name:"id" help:"list of policy ids to delete"` //nolint:revive
}

func (c *PolicyCmd) Run(k *kong.Kong, ctx *client.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PolicyCmd) Help() string { _ = "STUB: not implemented"; return "" }
