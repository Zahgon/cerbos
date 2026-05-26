// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package policy

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
	"github.com/cerbos/cerbos/internal/policy"
)

func DoCmd(k *kong.Kong, ac *cerbos.GRPCAdminClient, kind policy.Kind, filters *flagset.Filters, format *flagset.Format, sort *flagset.Sort, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

func List(k *kong.Kong, c *cerbos.GRPCAdminClient, filters *flagset.Filters, format *flagset.Format, sortFlags *flagset.Sort, kind policy.Kind) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

// no version or scope

func Get(k *kong.Kong, c *cerbos.GRPCAdminClient, format *flagset.Format, kind policy.Kind, ids ...string) error {
	_ = "STUB: not implemented"
	return nil
}
