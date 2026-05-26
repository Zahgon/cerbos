// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"github.com/alecthomas/kong"
	"github.com/cerbos/cerbos-sdk-go/cerbos/hub"
)

const (
	defaultMessage = "Uploaded using cerbosctl"
	defaultName    = "cerbosctl"
	defaultSource  = "cerbosctl"
)

const replaceFilesHelp = `
Replaces or deletes all files in the remote store so that it only contains the files provided.

The following exit codes have a special meaning.
	- 6: The version condition supplied using --version-must-eq wasn't satisfied

# Upload a local directory

cerbosctl hub store replace-files /path/to/dir

# Upload a local zip archive

cerbosctl hub store replace-files /path/to/archive.zip
`

type ReplaceFilesCmd struct { //betteralign:ignore
	Output        `embed:""`
	Path          string `arg:"" type:"path" help:"Path to a directory or a zip file containing the contents to upload" required:""`
	ChangeDetails `embed:""`
	VersionMustEq int64 `help:"Require that the store is at this version before committing the change" optional:""`
}

func (*ReplaceFilesCmd) Help() string { _ = "STUB: not implemented"; return "" }

func (rfc *ReplaceFilesCmd) Run(k *kong.Kong, cmd *Cmd) error {
	_ = "STUB: not implemented"
	return nil
}

func replaceFiles(storeClient *hub.StoreClient, storeID, path string, cd ChangeDetails, versionMustEq int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:nestif

func changeDetailsFromGit(path string) (*changeDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
