// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"errors"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/flagset"
)

var errInvalidCredentials = errors.New("invalid credentials: username and password must be non-empty strings")

type Context struct {
	Client      *cerbos.GRPCClient
	AdminClient *cerbos.GRPCAdminClient
}

func GetAdminClient(globals *flagset.Globals) (*cerbos.GRPCAdminClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetClient(globals *flagset.Globals) (*cerbos.GRPCClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
