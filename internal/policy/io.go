// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package policy

import (
	"io"
	"io/fs"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/parser"
)

func ReadPolicyFromFile(fsys fs.FS, path string) (*policyv1.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadPolicy reads a policy from the given reader.
func ReadPolicy(src io.Reader) (*policyv1.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadPolicyWithSourceContext reads a policy and returns it along with information about its source.
func ReadPolicyWithSourceContext(fsys fs.FS, path string) (*policyv1.Policy, parser.SourceCtx, error) {
	_ = "STUB: not implemented"
	return nil, *new(parser.SourceCtx), nil
}

func ReadPolicyWithSourceContextFromReader(src io.Reader) (*policyv1.Policy, parser.SourceCtx, error) {
	_ = "STUB: not implemented"
	return nil, *new(parser.SourceCtx), nil
}

// TODO: Temporary restriction during parser migration to protoyaml.

// FindPolicy finds a policy by ID from the given reader.
func FindPolicy(src io.Reader, modID namer.ModuleID) (*policyv1.Policy, parser.SourceCtx, error) {
	_ = "STUB: not implemented"
	return nil, *new(parser.SourceCtx), nil
}

// WritePolicy writes a policy as YAML to the destination.
func WritePolicy(dest io.Writer, p *policyv1.Policy) error { _ = "STUB: not implemented"; return nil }

// WriteBinaryPolicy writes a policy as binary (protobuf encoding).
func WriteBinaryPolicy(dest io.Writer, p *policyv1.Policy) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadBinaryPolicy reads a policy from binary (protobuf encoding).
func ReadBinaryPolicy(src io.Reader) (*policyv1.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
