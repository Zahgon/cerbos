// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	"github.com/cerbos/cerbos/internal/storage"
)

type (
	ExcludeMethod     func(string) bool
	IncludeKeysMethod func(string) bool
)

func NewUnaryInterceptor(log Log, store storage.Store, exclude ExcludeMethod) (grpc.UnaryServerInterceptor, error) {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor), nil
}

type cerbosAPIResponse interface {
	proto.Message
	GetCerbosCallId() string
}

func setCerbosCallID(callID string, resp any) { _ = "STUB: not implemented"; return }

// don't panic in case there's nil pointer error

type cerbosRequestWithContext interface {
	GetRequestContext() *auditv1.RequestContext
}

func extractRequestContext(req any) *auditv1.RequestContext { _ = "STUB: not implemented"; return nil }
