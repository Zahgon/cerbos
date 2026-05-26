// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"github.com/alecthomas/kong"
)

const getFilesHelp = `
# Download foo.yaml and bar.yaml to local directory

cerbosctl hub store get-files --output-path=/path/to/dir foo.yaml bar.yaml

# Download foo.yaml and bar.yaml as a zip archive

cerbosctl hub store get-files --output-path=/path/to/archive.zip foo.yaml bar.yaml
`

type GetFilesCmd struct { //betteralign:ignore
	Output     `embed:""`
	OutputPath string   `name:"output-path" short:"O" type:"path" required:"" help:"Path to write the retrieved files. Must be a path to a directory, zip file or - for stdout."`
	Files      []string `arg:"" required:"" help:"List of files to retrieve"`
}

func (*GetFilesCmd) Help() string { _ = "STUB: not implemented"; return "" }

func (gfc *GetFilesCmd) Run(k *kong.Kong, cmd *Cmd) (outErr error) {
	_ = "STUB: not implemented"
	return nil
}
