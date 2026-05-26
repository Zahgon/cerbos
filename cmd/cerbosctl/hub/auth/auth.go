// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package auth

import (
	"crypto/tls"

	"github.com/alecthomas/kong"
)

type Cmd struct {
	LogLevel      string `help:"Log level (${enum})" default:"info" enum:"debug,info,warn,error"`
	APIEndpoint   string `name:"api-endpoint" default:"https://api.cerbos.cloud" env:"CERBOS_HUB_API_ENDPOINT"`
	ClientID      string `name:"client-id" help:"Client ID of the access credential" env:"CERBOS_HUB_CLIENT_ID" and:"client-id,client-secret"`
	ClientSecret  string `name:"client-secret" help:"Client secret of the access credential" env:"CERBOS_HUB_CLIENT_SECRET" and:"client-id,client-secret"` //nolint:gosec
	TLSCACert     string `name:"tls-ca-cert" hidden:"" help:"Path to the CA certificate for verifying server identity" type:"existingfile" env:"CERBOS_HUB_TLS_CA_CERT"`
	TLSClientCert string `name:"tls-client-cert" hidden:"" help:"Path to the TLS client certificate" type:"existingfile" env:"CERBOS_HUB_TLS_CLIENT_CERT" and:"tls-client-key"`
	TLSClientKey  string `name:"tls-client-key" hidden:"" help:"Path to the TLS client key" type:"existingfile" env:"CERBOS_HUB_TLS_CLIENT_KEY" and:"tls-client-cert"`
	TLSInsecure   bool   `name:"tls-insecure" hidden:"" help:"Skip validating server certificate" env:"CERBOS_HUB_TLS_INSECURE"`
}

func (c *Cmd) Run(k *kong.Kong, cmd *Cmd) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (c *Cmd) buildTLSConf() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }
