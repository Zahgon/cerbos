// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit

import (
	"context"

	"google.golang.org/grpc/metadata"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
)

const (
	grpcGWUserAgentKey = "grpcgateway-user-agent"
	userAgentKey       = "user-agent"
	xffKey             = "x-forwarded-for"
	callIDTagKey       = "call_id"

	SetByGRPCGatewayKey = "x-cerbos-set-by-grpc-gateway"
	HTTPRemoteAddrKey   = "x-cerbos-http-remote-addr"
)

var SetByGRPCGatewayVal string

func init() {
	SetByGRPCGatewayVal = generateSetByGRPCGatewayVal()
}

type callIDCtxKeyType struct{}

var callIDCtxKey = callIDCtxKeyType{}

type requestContextKeyType struct{}

var requestContextKey = requestContextKeyType{}

func NewContextWithCallID(ctx context.Context, id ID) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func CallIDFromContext(ctx context.Context) (ID, bool) {
	_ = "STUB: not implemented"
	return *new(ID), false
}

func NewContextWithRequestContext(ctx context.Context, reqCtx *auditv1.RequestContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func RequestContextFromContext(ctx context.Context) *auditv1.RequestContext {
	_ = "STUB: not implemented"
	return nil
}

func PeerFromContext(ctx context.Context) *auditv1.Peer { _ = "STUB: not implemented"; return nil }

func peerFromContext(ctx context.Context) *auditv1.Peer { _ = "STUB: not implemented"; return nil }

func generateSetByGRPCGatewayVal() string { _ = "STUB: not implemented"; return "" }

func checkSetByGRPCGateway(md metadata.MD) bool { _ = "STUB: not implemented"; return false }
