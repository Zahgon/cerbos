// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package util

import (
	"crypto/tls"
	"net/http"

	"google.golang.org/grpc"
)

const unixNetwork = "unix"

// ParseListenAddress parses an address and returns the network type and the address to dial.
// inspired by https://github.com/ghostunnel/ghostunnel/blob/6e58c75c8762fe371c1134e89dd55033a6d577a4/socket/net.go#L31
func ParseListenAddress(listenAddr string) (network, addr string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// DefaultTLSConfig returns the default TLS configuration.
func DefaultTLSConfig() *tls.Config {
	_ = "STUB: not implemented"
	// See https://wiki.mozilla.org/Security/Server_Side_TLS
	return nil
}

func GetFreeListenAddr() (string, error) { _ = "STUB: not implemented"; return "", nil }

//nolint:noctx

func GetFreePort() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func NewInsecureHTTPClient(httpListenAddr string, tlsSpecified bool) (client *http.Client, httpAddr string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

//nolint:gosec

func newTransportForAddress(network, addr string, tlsConfig *tls.Config) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

//nolint:forcetypeassert

// EagerGRPCClient creates a gRPC client and establishes a connection immediately.
func EagerGRPCClient(target string, dialOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
