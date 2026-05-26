// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"github.com/alecthomas/kong"
)

const help = `
Examples:

# Start the server

cerbos server

# Start the server with the Admin API enabled and the 'sqlite' storage driver

cerbos server --set=server.adminAPI.enabled=true --set=storage.driver=sqlite3 --set=storage.sqlite3.dsn=':memory:'`

type LogLevelFlag string

func (ll *LogLevelFlag) Decode(ctx *kong.DecodeContext) error {
	_ = "STUB: not implemented"
	return nil
}

type Cmd struct { //betteralign:ignore
	DebugListenAddr string       `help:"Address to start the gops listener" placeholder:":6666"`
	LogLevel        LogLevelFlag `help:"Log level (${enum})" default:"info" enum:"debug,info,warn,error"`
	Config          string       `help:"Path to config file" optional:"" placeholder:".cerbos.yaml" env:"CERBOS_CONFIG"`
	HubBundle       string       `help:"[Legacy] Use Cerbos Hub to pull the policy bundle with the given label. Overrides the store defined in the configuration." optional:"" env:"CERBOS_HUB_BUNDLE,CERBOS_CLOUD_BUNDLE" group:"hubv1" xor:"hub.deployment-id,hub.playground-id"`
	Hub             HubFlags     `embed:"" prefix:"hub."`
	Set             []string     `help:"Config overrides" placeholder:"server.adminAPI.enabled=true"`
}

type HubFlags struct {
	DeploymentID string `help:"Use Cerbos Hub to pull the policy bundle for the given deployment ID. Overrides the store defined in the configuration." optional:"" env:"CERBOS_HUB_DEPLOYMENT_ID" group:"hubv2" xor:"hub-bundle,hub.playground-id"`
	PlaygroundID string `help:"Use Cerbos Hub to pull the policy bundle for the given playground ID. Overrides the store defined in the configuration." optional:"" env:"CERBOS_HUB_PLAYGROUND_ID" group:"hubv2" xor:"hub-bundle,hub.deployment-id"`
}

func MkHubOverrides(c *Cmd) []string { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Run() error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }
