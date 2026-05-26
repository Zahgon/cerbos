// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"
	"io/fs"
	"time"

	"go.uber.org/zap"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	svcv1 "github.com/cerbos/cerbos/api/genpb/cerbos/svc/v1"
	"github.com/cerbos/cerbos/internal/auxdata"
	"github.com/cerbos/cerbos/internal/compile"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/storage/index"
)

const playgroundRequestTimeout = 60 * time.Second

var _ svcv1.CerbosPlaygroundServiceServer = (*CerbosPlaygroundService)(nil)

// CerbosPlaygroundService implements the playground API.
type CerbosPlaygroundService struct {
	*svcv1.UnimplementedCerbosPlaygroundServiceServer
	auxData   *auxdata.AuxData
	reqLimits RequestLimits
}

func NewCerbosPlaygroundService(reqLimits RequestLimits) *CerbosPlaygroundService {
	_ = "STUB: not implemented"
	return nil
}

func (cs *CerbosPlaygroundService) PlaygroundValidate(ctx context.Context, req *requestv1.PlaygroundValidateRequest) (*responsev1.PlaygroundValidateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *CerbosPlaygroundService) PlaygroundTest(ctx context.Context, req *requestv1.PlaygroundTestRequest) (*responsev1.PlaygroundTestResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *CerbosPlaygroundService) PlaygroundEvaluate(ctx context.Context, req *requestv1.PlaygroundEvaluateRequest) (*responsev1.PlaygroundEvaluateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *CerbosPlaygroundService) PlaygroundProxy(ctx context.Context, req *requestv1.PlaygroundProxyRequest) (*responsev1.PlaygroundProxyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func doCompile(ctx context.Context, log *zap.Logger, files []*requestv1.File) (*components, *responsev1.PlaygroundFailure, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func buildIndex(ctx context.Context, log *zap.Logger, files []*requestv1.File) (index.Index, error) {
	_ = "STUB: not implemented"
	return *new(index.Index), nil
}

func buildFS(log *zap.Logger, files []*requestv1.File) (fs.FS, error) {
	_ = "STUB: not implemented"
	return *new(fs.FS), nil
}

//nolint:mnd

func processLintErrors(ctx context.Context, errs *index.BuildError) *responsev1.PlaygroundFailure {
	_ = "STUB: not implemented"
	return nil
}

//nolint:prealloc

//nolint:staticcheck

func processCompileErrors(ctx context.Context, errs *compile.ErrorSet) *responsev1.PlaygroundFailure {
	_ = "STUB: not implemented"
	return nil
}

func processEngineOutput(_ context.Context, playgroundID string, outputs []*enginev1.CheckOutput) (*responsev1.PlaygroundEvaluateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type components struct {
	idx       index.Index
	store     storage.SourceStore
	schemaMgr schema.Manager
}

func (c *components) mkEngine(ctx context.Context) (*engine.Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
