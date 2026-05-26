// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"iter"

	"github.com/alecthomas/kong"

	storev1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/store/v1"
)

const addFilesHelp = `
The following exit codes have a special meaning.
	- 6: The version condition supplied using --version-must-eq wasn't satisfied

# Upload foo.yaml and all files in the bar directory

cerbosctl hub store add-files foo.yaml bar

# Upload bar.yaml, renaming it to foo/bar.yaml in the store

cerbosctl hub store add-files --message="Adding foo/bar.yaml" foo/bar.yaml=bar.yaml
`

type AddFilesCmd struct { //betteralign:ignore
	filesToAdd    map[string]string
	Output        `embed:""`
	Message       string   `help:"Commit message for this change" default:"Uploaded using cerbosctl"`
	Paths         []string `arg:"" help:"List of files or directories to add to the store. To rename how the file appears in the store, use store_path=actual_path as the input format." required:""`
	VersionMustEq int64    `help:"Require that the store is at this version before committing the change" optional:""`
}

func (*AddFilesCmd) Help() string { _ = "STUB: not implemented"; return "" }

func (afc *AddFilesCmd) Validate() error { _ = "STUB: not implemented"; return nil }

//nolint:nestif

func (afc *AddFilesCmd) Run(k *kong.Kong, cmd *Cmd) error { _ = "STUB: not implemented"; return nil }

func (afc *AddFilesCmd) batch() iter.Seq2[[]*storev1.FileOp, error] {
	_ = "STUB: not implemented"
	return nil
}
