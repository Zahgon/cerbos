// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"net/http"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	privatev1 "github.com/cerbos/cerbos/api/genpb/cerbos/private/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
)

const (
	requestTimeout     = 5 * time.Second
	healthPollInterval = 100 * time.Millisecond
	retryBackoffDelay  = 5
)

type AuthCreds struct {
	Username string
	Password string //nolint:gosec
}

func (ac AuthCreds) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (AuthCreds) RequireTransportSecurity() bool { _ = "STUB: not implemented"; return false }

func LoadTestCases(tb testing.TB, suiteSleeps map[string]time.Duration, dirs ...string) *TestRunner {
	_ = "STUB: not implemented"
	return nil
}

//nolint:prealloc

// no point sleeping after the final suite

func readTestCase(tb testing.TB, name string, data []byte) *privatev1.ServerTestCase {
	_ = "STUB: not implemented"
	return nil
}

type TestRunner struct {
	sleeps                 map[int]time.Duration
	Cases                  []*privatev1.ServerTestCase
	Timeout                time.Duration
	HealthPollInterval     time.Duration
	CerbosClientMaxRetries uint
}

// WithCerbosClientRetries is relevant to Overlay storage driver calls (specifically the e2e overlay test).
func (tr *TestRunner) WithCerbosClientRetries(nRetries uint) *TestRunner {
	_ = "STUB: not implemented"
	return nil
}

func (tr *TestRunner) RunGRPCTests(addr string, opts ...grpc.DialOption) func(*testing.T) {
	_ = "STUB: not implemented"
	//nolint:thelper
	return nil
}

func mkGRPCConn(t *testing.T, addr string, opts ...grpc.DialOption) *grpc.ClientConn {
	_ = "STUB: not implemented"
	return nil
}

func (tr *TestRunner) executeGRPCTestCase(grpcConn *grpc.ClientConn, tc *privatev1.ServerTestCase) func(*testing.T) {
	_ = "STUB: not implemented"
	//nolint:thelper
	return nil
}

func (tr *TestRunner) RunHTTPTests(hostAddr string, creds *AuthCreds) func(*testing.T) {
	_ = "STUB: not implemented"
	//nolint:thelper
	return nil
}

func mkHTTPClient(t *testing.T) *http.Client { _ = "STUB: not implemented"; return nil }

//nolint:forcetypeassert
//nolint:gosec

func (tr *TestRunner) executeHTTPTestCase(c *http.Client, hostAddr string, creds *AuthCreds, tc *privatev1.ServerTestCase) func(*testing.T) {
	_ = "STUB: not implemented"
	//nolint:thelper
	return nil
}

//nolint:gosec

func (tr *TestRunner) checkCORS(c *http.Client, hostAddr string) func(*testing.T) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:thelper

//nolint:gosec

func compareProto(t *testing.T, want, have proto.Message) { _ = "STUB: not implemented"; return }

func cmpPlaygroundEvalResult(a, b *responsev1.PlaygroundEvaluateResponse_EvalResult) bool {
	_ = "STUB: not implemented"
	return false
}

func cmpPlaygroundError(a, b *responsev1.PlaygroundFailure_Error) bool {
	_ = "STUB: not implemented"
	return false
}

func cmpValidationError(a, b *schemav1.ValidationError) bool {
	_ = "STUB: not implemented"
	return false
}

func cmpOutputs(a, b *enginev1.OutputEntry) bool { _ = "STUB: not implemented"; return false }

func TraceBatchToTraces(batch protocmp.Message) *privatev1.TestTracesWrapper {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

func grpcHealthCheckPasses(t *testing.T, grpcConn *grpc.ClientConn, reqTimeout time.Duration) func() bool {
	_ = "STUB: not implemented"
	return nil
}

func httpHealthCheckPasses(client *http.Client, url string, reqTimeout time.Duration) func() bool {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec
