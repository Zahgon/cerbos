// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"github.com/alecthomas/kong"

	storev1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/store/v1"
)

const listFilesHelp = `
# List all files

cerbosctl hub store list-files

# List files matching "resource"

cerbosctl hub store list-files --filter=contains:resource
`

type ListFilesCmd struct { //betteralign:ignore
	Output `embed:""`
	Filter string `name:"filter" optional:"" help:"Optional file name filter in the form <operator>:<value>. Supported operators are 'eq', 'in' and 'contains'. For 'in' multiple values can be provided as a comma separated list."`
}

func (*ListFilesCmd) Help() string { _ = "STUB: not implemented"; return "" }

func (lfc *ListFilesCmd) Run(k *kong.Kong, cmd *Cmd) error { _ = "STUB: not implemented"; return nil }

type listFilesOutput struct {
	*storev1.ListFilesResponse
}

func (lfo listFilesOutput) String() string { _ = "STUB: not implemented"; return "" }

func (lfo listFilesOutput) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
