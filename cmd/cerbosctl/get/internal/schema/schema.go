// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"github.com/alecthomas/kong"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/cerbos/cerbos/cmd/cerbosctl/get/internal/flagset"
)

func List(k *kong.Kong, c *cerbos.GRPCAdminClient, format *flagset.Format) error {
	_ = "STUB: not implemented"
	return nil
}

func Get(k *kong.Kong, c *cerbos.GRPCAdminClient, format *flagset.Format, ids ...string) error {
	_ = "STUB: not implemented"
	return nil
}
