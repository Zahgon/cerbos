// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package reload

import (
	"github.com/alecthomas/kong"

	cmdclient "github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const help = `# Reload the store
cerbosctl store reload

# Reload the store and wait until it finishes
cerbosctl store reload --wait`

type Cmd struct { //betteralign:ignore
	Wait bool `help:"Wait until the reloading process finishes"`
}

func (c *Cmd) Run(k *kong.Kong, ctx *cmdclient.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
