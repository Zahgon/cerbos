// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package awslambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"github.com/cerbos/cerbos/internal/svc"
)

// RouteRequest routes the API Gateway request to the appropriate handler.
func RouteRequest(ctx context.Context, event events.APIGatewayV2HTTPRequest, svc *svc.CerbosService) (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse), nil
}

func handleHealthCheck() (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse), nil
}

func handleCheckResources(ctx context.Context, event events.APIGatewayV2HTTPRequest, svc *svc.CerbosService) (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse), nil
}

func handlePlanResources(ctx context.Context, event events.APIGatewayV2HTTPRequest, svc *svc.CerbosService) (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse), nil
}
