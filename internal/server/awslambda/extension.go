// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package awslambda

import (
	"context"
	"net/http"
)

type lambdaExt struct {
	client      *http.Client
	runtimeAPI  string
	extensionID string
}

type RegisterRequest struct {
	Events []string `json:"events"`
}

type EventResponse struct {
	Tracing struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"tracing"`
	EventType          string `json:"eventType"`
	RequestID          string `json:"requestId"` //nolint:tagliatelle
	InvokedFunctionArn string `json:"invokedFunctionArn"`
	DeadlineMs         int64  `json:"deadlineMs"`
}

const (
	extensionNameHeader  = "Lambda-Extension-Name"
	extensionIDHeader    = "Lambda-Extension-Identifier"
	extensionErrorType   = "Lambda-Extension-Function-Error-Type"
	registrationEndpoint = "/2020-01-01/extension/register"
	nextEventEndpoint    = "/2020-01-01/extension/event/next"
	exitErrorEndpoint    = "/2020-01-01/extension/exit/error"
)

const maxBodySize = 1024

func RegisterNewExtension(ctx context.Context, runtimeAPI string) (*lambdaExt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:mnd

//nolint:gosec

func (l *lambdaExt) CheckShutdown(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

//nolint:gosec

func (l *lambdaExt) ReportError(ctx context.Context, err error) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func WaitForReady(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
