// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package awslambda

import (
	"github.com/aws/aws-lambda-go/events"
	"google.golang.org/protobuf/proto"

	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
)

// parseRequestBody handles the common body parsing logic.
func parseRequestBody(event events.APIGatewayV2HTTPRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// APIGatewayEventToCheckResourcesRequest converts an API Gateway event to a CheckResourcesRequest.
func APIGatewayEventToCheckResourcesRequest(event events.APIGatewayV2HTTPRequest) (*requestv1.CheckResourcesRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// APIGatewayEventToPlanResourcesRequest converts an API Gateway event to a PlanResourcesRequest.
func APIGatewayEventToPlanResourcesRequest(event events.APIGatewayV2HTTPRequest) (*requestv1.PlanResourcesRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResponseToAPIGateway converts a protobuf response to an API Gateway response.
func ResponseToAPIGateway[T proto.Message](resp T) (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse), nil
}

func CheckResourcesResponseToAPIGateway(resp *responsev1.CheckResourcesResponse) (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse), nil
}

func PlanResourcesResponseToAPIGateway(resp *responsev1.PlanResourcesResponse) (events.APIGatewayV2HTTPResponse, error) {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse), nil
}

func ErrorToAPIGateway(message string, statusCode int) events.APIGatewayV2HTTPResponse {
	_ = "STUB: not implemented"
	return *new(events.APIGatewayV2HTTPResponse)
}
