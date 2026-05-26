// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package awslambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/cerbos/cerbos/internal/server"
	"github.com/cerbos/cerbos/internal/svc"
)

// FunctionHandler handles AWS Lambda function invocations.
type FunctionHandler struct {
	svc  *svc.CerbosService
	core *server.CoreComponents
}

func NewFunctionHandler(ctx context.Context) (*FunctionHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *FunctionHandler) Handle(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *

	// Add request ID and other metadata to logging context
	new(events.APIGatewayV2HTTPResponse), nil
}

// TODO: AWS X-Ray trace ID is available in context or headers, not directly in event

func (h *FunctionHandler) Close() error { _ = "STUB: not implemented"; return nil }
