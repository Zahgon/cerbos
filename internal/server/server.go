// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	grpcruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	// Import the default grpc encoding to ensure that it gets replaced by VT.
	_ "google.golang.org/grpc/encoding/proto"
	"google.golang.org/grpc/health"

	// Import to register the Badger audit log backend.
	_ "github.com/cerbos/cerbos/internal/audit/local"
	// Import to register the file audit log backend.
	_ "github.com/cerbos/cerbos/internal/audit/file"
	// Import to register the kafka audit log backend.
	_ "github.com/cerbos/cerbos/internal/audit/kafka"
	// Import to register the hub audit log backend.
	_ "github.com/cerbos/cerbos/internal/audit/hub"

	// Import blob to register the storage driver.
	_ "github.com/cerbos/cerbos/internal/storage/blob"

	// Import hub to register the storage driver.
	_ "github.com/cerbos/cerbos/internal/storage/hub"
	// Import mysql to register the storage driver.
	_ "github.com/cerbos/cerbos/internal/storage/db/mysql"
	// Import postgres to register the storage driver.
	_ "github.com/cerbos/cerbos/internal/storage/db/postgres"
	// Import sqlite3 to register the storage driver.
	_ "github.com/cerbos/cerbos/internal/storage/db/sqlite3"
	// Import disk to register the storage driver.
	_ "github.com/cerbos/cerbos/internal/storage/disk"
	// Import git to register the storage driver.
	_ "github.com/cerbos/cerbos/internal/storage/git"
)

const (
	defaultTimeout        = 30 * time.Second
	minGRPCConnectTimeout = 20 * time.Second

	adminEndpoint          = "/admin"
	apiEndpoint            = "/api"
	authzenEndpont         = "/access/v1"
	authzenMetadataEnpoint = "/.well-known/authzen-configuration"
	healthEndpoint         = "/_cerbos/health"
	metricsEndpoint        = "/_cerbos/metrics"
	playgroundEndpoint     = "/api/playground"
	schemaEndpoint         = "/schema/swagger.json"
)

var ErrInvalidStore = errors.New("store does not implement either SourceStore or BinaryStore interfaces")

func Start(ctx context.Context) error {
	_ = "STUB: not implemented"
	// get configuration
	return nil
}

type Server struct {
	conf       *Conf
	cancelFunc context.CancelFunc
	pool       *pool.ContextPool
	health     *health.Server
	tlsConfig  *tls.Config
}

func NewServer(conf *Conf) *Server { _ = "STUB: not implemented"; return nil }

func (s *Server) Start(ctx context.Context, core *CoreComponents) error {
	_ = "STUB: not implemented"
	return nil
}

// It would be nice to have a single port to serve both gRPC and HTTP. Unfortunately, cmux
// can't deal effectively with both gRPC and HTTP/2 when TLS is enabled (see https://github.com/soheilhy/cmux/issues/68).
// Another potential issue with single-port gRPC and HTTP/2 is when a proxy like Envoy is in front of the server it
// would have a connection pool per port and would end up sending HTTP/2 traffic to gRPC and vice-versa.
// This is why we have two dedicated ports for HTTP and gRPC traffic. However, if gRPC traffic is sent to the HTTP port, it
// will still be handled correctly.

// start servers

// mark this service as NOT_SERVING in the gRPC health check.

func (s *Server) initializeTLSConfig(log *zap.Logger) error { _ = "STUB: not implemented"; return nil }

//nolint:nilerr

func (s *Server) createListener(ctx context.Context, listenAddr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (s *Server) startGRPCServer(l net.Listener, core *CoreComponents) (*grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkForUnsafeAdminCredentials(log *zap.Logger, passwordHash []byte) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) mkGRPCServer(log *zap.Logger, core *CoreComponents) (*grpc.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) startHTTPServer(ctx context.Context, l net.Listener, grpcSrv *grpc.Server, suggestHub bool) (*http.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handle gRPC requests that come over http

func mkGatewayMux(grpcConn grpc.ClientConnInterface) *grpcruntime.ServeMux {
	_ = "STUB: not implemented"
	return nil
}

type grpcJSONPb struct {
	grpcruntime.JSONPb
}

var _ grpcruntime.StreamContentType = (*grpcJSONPb)(nil)

func (*grpcJSONPb) StreamContentType(any) string { _ = "STUB: not implemented"; return "" }

func defaultGRPCDialOpts() []grpc.DialOption {
	_ = "STUB: not implemented"
	// see https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md
	return nil
}

func (s *Server) mkGRPCConn() (*grpc.ClientConn, error) { _ = "STUB: not implemented"; return nil, nil }

// we are connecting as localhost which would differ from what the cert is issued for.

// inspired by https://github.com/ghostunnel/ghostunnel/blob/6e58c75c8762fe371c1134e89dd55033a6d577a4/socket/net.go#L100
func (s *Server) parseAndOpen(ctx context.Context, listenAddr string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

//nolint:nestif

//nolint:forcetypeassert

//nolint:mnd
func toUDSFileMode(modeStr string) os.FileMode { _ = "STUB: not implemented"; return *new(os.FileMode) }

// Ignore everything but the last 9 bits which hold the user, group and world perms.

func incomingHeaderMatcher(key string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// The gateway sets its own user agent, so we need to alias it

// The request sent by the gateway will have a different content length

// Reserved for aliasing the incoming User-Agent header

// Translated to X-Forwarded-Host by the gateway

// Connection-specific headers must be removed when translating HTTP/1.x to HTTP/2 (https://httpwg.org/specs/rfc9113.html#ConnectionSpecific)

func setPeerMetadata(_ context.Context, req *http.Request) metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}
