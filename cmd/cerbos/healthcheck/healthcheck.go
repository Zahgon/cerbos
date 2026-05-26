// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package healthcheck

import (
	"context"
	"crypto/tls"
	"io"
	"time"

	"github.com/alecthomas/kong"
	"github.com/cerbos/cerbos/internal/server"
)

const (
	defaultGRPCHostPort = "127.0.0.1:3593"
	defaultHTTPHostPort = "127.0.0.1:3592"
	grpcKind            = "grpc"
	httpKind            = "http"

	help = `
Performs a healthcheck on a Cerbos PDP. This can be used as a Docker HEALTHCHECK command.
When the path to the Cerbos config file is provided via the '--config' flag or the CERBOS_CONFIG environment variable, the healthcheck will be automatically configured based on the settings from the file.
By default, the gRPC endpoint will be checked using the gRPC healthcheck protocol. This is usually sufficient for most cases as the Cerbos REST API is built on top of the gRPC API as well.

Examples:

# Check gRPC endpoint

cerbos healthcheck --config=/path/to/.cerbos.yaml

# Check HTTP endpoint and ignore server certificate verification

cerbos healthcheck --config=/path/to/.cerbos.yaml --kind=http --insecure

# Check the HTTP endpoint of a specific host with no TLS.

cerbos healthcheck --kind=http --host-port=10.0.1.5:3592 --no-tls
`
)

type Cmd struct { //betteralign:ignore
	Config   string        `help:"Cerbos config file" group:"config" xor:"hostport,cacert,notls" env:"CERBOS_CONFIG"`
	Kind     string        `help:"Healthcheck kind (${enum})" default:"grpc" enum:"grpc,http" env:"CERBOS_HC_KIND"`
	HostPort string        `help:"Host and port to connect to" group:"manual" xor:"hostport" env:"CERBOS_HC_HOSTPORT"`
	CACert   string        `help:"Path to CA cert for validating server cert" type:"existingfile" group:"manual" xor:"cacert" env:"CERBOS_HC_CACERT"`
	NoTLS    bool          `help:"Don't use TLS" group:"manual" xor:"notls" env:"CERBOS_HC_NOTLS"`
	Insecure bool          `help:"Do not verify server certificate" default:"false" env:"CERBOS_HC_INSECURE"`
	Timeout  time.Duration `help:"Healthcheck timeout" default:"2s" env:"CERBOS_HC_TIMEOUT"`
}

type checker interface {
	check(context.Context, io.Writer) error
}

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }

func (c *Cmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

func (c *Cmd) buildCheck() (checker, error) { _ = "STUB: not implemented"; return *new(checker), nil }

func (c *Cmd) doBuildCheckFromConf(serverConf *server.Conf) (checker, error) {
	_ = "STUB: not implemented"
	return *new(checker), nil
}

func (c *Cmd) doBuildCheckManual() (checker, error) {
	_ = "STUB: not implemented"
	return *new(checker), nil
}

func mkTLSConfig(tc *server.TLSConf, insecure bool) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type grpcCheck struct {
	tlsConf *tls.Config
	addr    string
}

func (gc grpcCheck) check(ctx context.Context, out io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

type httpCheck struct {
	tlsConf *tls.Config
	url     string
}

func newHTTPCheck(hostPort string, tlsConf *tls.Config) httpCheck {
	_ = "STUB: not implemented"
	return *new(httpCheck)
}

func (hc httpCheck) check(ctx context.Context, out io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

//nolint:gosec
