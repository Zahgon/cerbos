// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package revisions

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const purgeCmdHelp = `# Purge all store revisions

cerbosctl store revisions purge

# Purge store revisions but keep last 2 revisions of each policy

cerbosctl store revisions purge --keep-last=2`

type PurgeCmd struct { //betteralign:ignore
	KeepLast uint32 `help:"Keep last N revisions. If not specified or set to zero, all revisions will be deleted."` //nolint:revive
}

func (c *PurgeCmd) Run(k *kong.Kong, ctx *client.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PurgeCmd) Help() string { _ = "STUB: not implemented"; return "" }
