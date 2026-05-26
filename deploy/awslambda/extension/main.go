// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const cerbosHTTPAddr = "unix:/tmp/cerbos.http.sock"

func main() {
	gw, err := NewGateway(cerbosHTTPAddr)
	if err != nil {
		log.Print("failed to create a gateway")
		return
	}
	lambda.StartHandler(gw)
}

var ErrNotStarted = errors.New("timeout exceeded starting Cerbos")

type Gateway struct {
	httpClient    *http.Client
	cerbosAddress *url.URL
	socketPath    string
}

// NewGateway creates a new Gateway instance.
func NewGateway(addr string) (*Gateway, error) { _ = "STUB: not implemented"; return nil, nil }

func (g *Gateway) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newRequest returns a new http.Request from the given Lambda event.
func (g *Gateway) newRequest(ctx context.Context, e events.APIGatewayV2HTTPRequest) (*http.Request, error) {
	_ = "STUB: not implemented"
	// path
	return nil, nil
}

// base64 encoded body

// remote addr

// header fields

// content-length

// custom fields

// xray support

// host

func MkGatewayResponse(hresp *http.Response) (res *events.APIGatewayV2HTTPResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copy headers

// see https://aws.amazon.com/blogs/compute/simply-serverless-using-aws-lambda-to-expose-custom-cookies-with-api-gateway/

// isBinary checks content type of the returns true if it describes binary data
// It uses a non-exhaustive list of binary content types.
func isBinary(h http.Header) (bool, error) { _ = "STUB: not implemented"; return false, nil }
