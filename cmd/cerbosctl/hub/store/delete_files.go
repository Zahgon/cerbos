// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"github.com/alecthomas/kong"
)

const deleteFilesHelp = `
The following exit codes have a special meaning.
	- 6: The version condition supplied using --version-must-eq wasn't satisfied

# Delete foo/bar.yaml from the remote store

cerbosctl hub store delete-files --message="Deleting foo/bar.yaml" foo/bar.yaml
`

type DeleteFilesCmd struct { //betteralign:ignore
	Output        `embed:""`
	Message       string   `help:"Commit message for this change" default:"Uploaded using cerbosctl"`
	Paths         []string `arg:"" help:"List of paths to delete from the store" required:""`
	VersionMustEq int64    `help:"Require that the store is at this version before committing the change" optional:""`
}

func (*DeleteFilesCmd) Help() string { _ = "STUB: not implemented"; return "" }

func (dfc *DeleteFilesCmd) Run(k *kong.Kong, cmd *Cmd) error { _ = "STUB: not implemented"; return nil }
