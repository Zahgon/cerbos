// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package enable

import (
	"github.com/alecthomas/kong"

	internalclient "github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const policyCmdHelp = `# Enable policies
cerbosctl enable policies derived_roles.my_derived_roles
cerbosctl enable policy derived_roles.my_derived_roles
cerbosctl enable p derived_roles.my_derived_roles

# Enable multiple policies
cerbosctl enable policies derived_roles.my_derived_roles resource.leave_request.default
cerbosctl enable policy derived_roles.my_derived_roles resource.leave_request.default
cerbosctl enable p derived_roles.my_derived_roles resource.leave_request.default`

type Cmd struct { //betteralign:ignore
	Policy PolicyCmd `cmd:"" aliases:"policies,p"`
}

type PolicyCmd struct { //betteralign:ignore
	PolicyIds []string `arg:"" name:"id" help:"list of policy ids to enable"` //nolint:revive
}

func (c *Cmd) Run(k *kong.Kong, ctx *internalclient.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *PolicyCmd) Help() string { _ = "STUB: not implemented"; return "" }
