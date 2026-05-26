// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package run

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/alecthomas/kong"
	"github.com/go-cmd/cmd"
)

const (
	help = `
Launches a command within the context of a Cerbos PDP. The policies are loaded by default from a directory named "policies" in the current working directory. The launched application can access Cerbos endpoints using the values from CERBOS_HTTP or CERBOS_GRPC environment variables.

If a file named ".cerbos.yaml" exists in the current working directory, it will be used as the configuration file for the PDP. You can override the config file and/or other configuration options using the flags described below.

Examples:

# Launch Go tests within a Cerbos context

cerbos run -- go test ./...

# Start Cerbos with a custom configuration file and run Python tests within the context

cerbos run --config=myconf.yaml -- python -m unittest

# Silence Cerbos log output

cerbos run --log-level=error -- curl -I http://127.0.0.1:3592/_cerbos/health
	`

	confDefault = `
server:
  httpListenAddr: "127.0.0.1:3592"
  grpcListenAddr: "127.0.0.1:3593"
storage:
  driver: "disk"
  disk:
    directory: %q
    watchForChanges: true
`
)

type Cmd struct { //betteralign:ignore
	LogLevel string         `help:"Log level (${enum})" default:"info" enum:"debug,info,warn,error"`
	Config   string         `help:"Path to config file" type:"existingfile" placeholder:".cerbos.yaml"`
	Set      []string       `help:"Config overrides" placeholder:"server.adminAPI.enabled=true"`
	Command  []string       `help:"Command to run" arg:"" passthrough:"" required:""`
	Timeout  time.Duration  `help:"Cerbos startup timeout" default:"30s"`
	wg       sync.WaitGroup `kong:"-"`
}

func (c *Cmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (c *Cmd) loadConfig() error {
	_ = "STUB: not implemented"
	// load any config overrides
	return nil
}

// load configuration
//nolint:nestif

//nolint:mnd

func (c *Cmd) startPDP(ctx context.Context) (*pdpInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

func (c *Cmd) prepCommand(pdp *pdpInstance, stdout, stderr io.Writer) *cmd.Cmd {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmd) goroutine(fn func()) { _ = "STUB: not implemented"; return }

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }

type pdpInstance struct {
	errors   chan error
	stopFn   context.CancelFunc
	client   *http.Client
	httpAddr string
	grpcAddr string
}

func (pdp *pdpInstance) waitForReady(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
