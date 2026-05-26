// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"iter"

	"github.com/alecthomas/kong"
	storev1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/store/v1"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
)

const uploadGitHelp = `
The following exit codes have a special meaning.
	- 6: The version condition supplied using --version-must-eq wasn't satisfied

# Apply the file changes between remote store's current version and the local git repository's HEAD, unless
the remote store doesn't have any Git change details in the latest version in which case fallback to replacing all
files in the remote store by the files in the local git repository.

cerbosctl hub store upload-git
cerbosctl hub store upload-git --to=HEAD
cerbosctl hub store upload-git --to=HEAD --path path/to/git/repository
cerbosctl hub store upload-git --to=HEAD --path path/to/git/repository --subdir policies

# Apply the file changes recorded in the git repo between commit 55a4248 and HEAD

cerbosctl hub store upload-git --from=55a4248
cerbosctl hub store upload-git --from=55a4248 --to=HEAD
cerbosctl hub store upload-git --from=55a4248 --to=HEAD --path path/to/git/repository
cerbosctl hub store upload-git --from=55a4248 --to=HEAD --path path/to/git/repository --subdir policies

# Apply the file changes recorded in the git repo between commit 55a4248 to e746228

cerbosctl hub store upload-git --from=55a4248 --to=e746228
cerbosctl hub store upload-git --from=55a4248 --to=e746228 --path path/to/git/repository
cerbosctl hub store upload-git --from=55a4248 --to=e746228 --path path/to/git/repository --subdir policies
`

type UploadGitCmd struct { //betteralign:ignore
	repository    *git.Repository
	from          *plumbing.Hash
	to            *plumbing.Hash
	To            string `help:"Git revision to end when generating the diff" default:"HEAD"`
	From          string `help:"Git revision to start from when generating the diff (the resolved reference must be the ancestor of the to argument)"`
	Path          string `help:"Path to the git repository" default:"."`
	Subdirectory  string `help:"Subdirectory under the given path to check and upload changes from" aliases:"subdir" default:"."`
	Output        `embed:""`
	ChangeDetails `embed:""`
	VersionMustEq int64 `help:"Require that the store is at this version before committing the change" optional:""`
}

func (*UploadGitCmd) Help() string { _ = "STUB: not implemented"; return "" }

func (ugc *UploadGitCmd) Validate() error { _ = "STUB: not implemented"; return nil }

func (ugc *UploadGitCmd) Run(k *kong.Kong, cmd *Cmd) error { _ = "STUB: not implemented"; return nil }

//nolint:nestif

func (ugc *UploadGitCmd) batch(diffToApply *diff) iter.Seq2[[]*storev1.FileOp, error] {
	_ = "STUB: not implemented"
	return nil
}

func (ugc *UploadGitCmd) diff(r *git.Repository, from, to plumbing.Hash) (*diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ugc *UploadGitCmd) changes(objectChanges object.Changes) ([]*change, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ugc *UploadGitCmd) normalize(name string) (normalized string, skipped bool) {
	_ = "STUB: not implemented"
	return "", false
}

type diff struct {
	hash    plumbing.Hash
	changes []*change
}

type change struct {
	name      string
	path      string
	operation op
}

type op string

const (
	OpUnspecified op = "UNSPECIFIED"
	OpAddOrUpdate op = "ADD_OR_UPDATE"
	OpDelete      op = "DELETE"
)
