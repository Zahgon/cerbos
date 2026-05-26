// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"io"

	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/cmd/cerbosctl/inspect/internal/flagset"
)

const (
	separator = ","
	width     = 80
)

func Print(w io.Writer, format flagset.Format, response *responsev1.InspectPoliciesResponse) error {
	_ = "STUB: not implemented"
	return nil
}

func printYAML(w io.Writer, results []*responsev1.InspectPoliciesResponse_Result) error {
	_ = "STUB: not implemented"
	return nil
}

func printPrettyJSON(w io.Writer, results []*responsev1.InspectPoliciesResponse_Result) error {
	_ = "STUB: not implemented"
	return nil
}

func printJSON(w io.Writer, results []*responsev1.InspectPoliciesResponse_Result) error {
	_ = "STUB: not implemented"
	return nil
}

func printTable(w io.Writer, noHeaders bool, results []*responsev1.InspectPoliciesResponse_Result) {
	_ = "STUB: not implemented"
	return
}

//nolint:nestif
