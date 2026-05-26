// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"io"
	"io/fs"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	sourcev1 "github.com/cerbos/cerbos/api/genpb/cerbos/source/v1"
)

type ReadError struct {
	Errors []*sourcev1.Error
}

func (ReadError) Error() string { _ = "STUB: not implemented"; return "" }

func Wrap(p *policyv1.Policy) *sourcev1.PolicyWrapper { _ = "STUB: not implemented"; return nil }

func ReadFromFile(fsys fs.FS, path string) (*policyv1.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Read(src io.Reader) (*policyv1.Policy, error) { _ = "STUB: not implemented"; return nil, nil }

func IDFromPolicyKey(key string) uint64 { _ = "STUB: not implemented"; return 0 }
