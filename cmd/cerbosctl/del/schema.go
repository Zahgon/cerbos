// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package del

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const schemaCmdHelp = `# Delete schemas

cerbosctl delete schemas principal.json

cerbosctl delete schema principal.json

cerbosctl delete s principal.json

# Delete multiple schemas

cerbosctl delete schemas principal.json leave_request.json

cerbosctl delete schema principal.json leave_request.json

cerbosctl delete s principal.json leave_request.json`

type SchemaCmd struct { //betteralign:ignore
	SchemaIds []string `arg:"" name:"id" help:"list of schema ids to delete"` //nolint:revive
}

func (c *SchemaCmd) Run(k *kong.Kong, ctx *client.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SchemaCmd) Help() string { _ = "STUB: not implemented"; return "" }
