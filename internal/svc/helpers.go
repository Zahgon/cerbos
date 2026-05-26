// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"
)

const (
	metaTagKey         = "grpc.request.meta"
	requestIDTagKey    = "request_id"
	playgroundIDTagKey = "playground_id"
)

func ExtractRequestFields(fullMethod string, req any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func SetHTTPStatusCode(ctx context.Context, code int) { _ = "STUB: not implemented"; return }
