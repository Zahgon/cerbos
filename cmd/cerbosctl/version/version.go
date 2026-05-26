// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"github.com/alecthomas/kong"

	cmdclient "github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

type Cmd struct { //betteralign:ignore
	Client kong.VersionFlag `help:"Only show cerbosctl version"`
}

func (c *Cmd) Run(k *kong.Kong, ctx *cmdclient.Context) error {
	_ = "STUB: not implemented"
	return nil
}
