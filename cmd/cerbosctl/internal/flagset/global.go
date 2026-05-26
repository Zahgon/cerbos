// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import "github.com/cerbos/cerbos-sdk-go/cerbos"

const maxRecvMsgSizeBytes = 25 * 1024 * 1024 // 25MiB

type Globals struct {
	Server        string `help:"Address of the Cerbos server" env:"CERBOS_SERVER" default:"localhost:3593"`
	Username      string `help:"Admin username" env:"CERBOS_USERNAME"`
	Password      string `help:"Admin password" env:"CERBOS_PASSWORD"` //nolint:gosec
	CaCert        string `help:"Path to the CA certificate for verifying server identity"`
	TLSClientCert string `name:"client-cert" help:"Path to the TLS client certificate"`
	TLSClientKey  string `name:"client-key" help:"Path to the TLS client key"`
	Insecure      bool   `help:"Skip validating server certificate"`
	Plaintext     bool   `help:"Use plaintext protocol without TLS"`
}

func (g *Globals) ToClientOpts() []cerbos.Opt { _ = "STUB: not implemented"; return nil }
