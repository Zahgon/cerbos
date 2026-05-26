// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const (
	adminSvcDisabled      = "Admin service is disabled by the configuration"
	playgroundSvcDisabled = "Playground service is disabled by the configuration"
	unknownSvc            = "Unknown service"
)

type methodNameCtxKeyType struct{}

var methodNameCtxKey = &methodNameCtxKeyType{}

func RequestMetadataUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	// New logging interceptor doesn't have access to method name so we save it to context for later use.
	return *new(any), nil
}

//nolint:mnd

// Fields are key-value pairs. Because we are adding "meta" and "http", the expected length is 4.
//nolint:mnd

func RequestLogger(log *zap.Logger, msg string) logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

func PayloadLogger(conf *Conf) logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

// accessLogExclude decides which methods to exclude from being logged to the access log.
func accessLogExclude(method string) bool { _ = "STUB: not implemented"; return false }

// prettyJSON instructs grpc-gateway to output pretty JSON when the query parameter is present.
func prettyJSON(h http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func customHTTPResponseCode(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func withCORS(conf *Conf, handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// The cors library's defaults don't include user-agent so we explicitly add it here.

func handleUnknownServices(_ any, stream grpc.ServerStream) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:govet

//nolint:mnd

func handleRoutingError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	_ = "STUB: not implemented"
	return
}

//nolint:govet

func cerbosVersionUnaryServerInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
