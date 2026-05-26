// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"io"

	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
	"github.com/cerbos/cerbos/internal/policy"
)

func printPolicy(w io.Writer, policies []policy.Wrapper, output flagset.OutputFormat) error {
	_ = "STUB: not implemented"
	return nil
}

func printPolicyJSON(w io.Writer, policies []policy.Wrapper) error {
	_ = "STUB: not implemented"
	return nil
}

func printPolicyPrettyJSON(w io.Writer, policies []policy.Wrapper) error {
	_ = "STUB: not implemented"
	return nil
}

func printPolicyYAML(w io.Writer, policies []policy.Wrapper) error {
	_ = "STUB: not implemented"
	return nil
}

func getHeaders(kind policy.Kind) []string { _ = "STUB: not implemented"; return nil }
