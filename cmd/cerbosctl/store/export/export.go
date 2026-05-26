// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package export

import (
	"github.com/alecthomas/kong"

	internalclient "github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
)

const help = `# Export policies and schemas from the store into a directory
cerbosctl store export path/to/dir

# Export policies and schemas from the store into a zip archive
cerbosctl store export archive.zip

# Export policies and schemas from the store into a gzip archive
cerbosctl store export path/to/archive.gzip
cerbosctl store export path/to/archive.tar.gz`

type Cmd struct { //betteralign:ignore
	Path string `arg:"" help:"Path to write policies and schemas" type:"path"`
}

func (c *Cmd) Run(k *kong.Kong, clientCtx *internalclient.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
