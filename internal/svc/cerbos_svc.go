// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"

	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	svcv1 "github.com/cerbos/cerbos/api/genpb/cerbos/svc/v1"
	"github.com/cerbos/cerbos/internal/auxdata"
	"github.com/cerbos/cerbos/internal/engine"
)

var _ svcv1.CerbosServiceServer = (*CerbosService)(nil)

// CerbosService implements the policy checking service.
type CerbosService struct {
	eng     *engine.Engine
	auxData *auxdata.AuxData
	*svcv1.UnimplementedCerbosServiceServer
	reqLimits RequestLimits
}

type RequestLimits struct {
	MaxActionsPerResource  uint
	MaxResourcesPerRequest uint
}

func NewCerbosService(eng *engine.Engine, auxData *auxdata.AuxData, reqLimits RequestLimits) *CerbosService {
	_ = "STUB: not implemented"
	return nil
}

func (cs *CerbosService) PlanResources(ctx context.Context, request *requestv1.PlanResourcesRequest) (*responsev1.PlanResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

// CheckResourceSet checks a batch of homogenous resources.
//
// Deprecated: Since 0.16.0. Use CheckResources instead.
func (cs *CerbosService) CheckResourceSet(ctx context.Context, req *requestv1.CheckResourceSetRequest) (*responsev1.CheckResourceSetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckResourceBatch checks a batch of heterogenous resources.
//
// Deprecated: Since 0.16.0. Use CheckResources instead.
func (cs *CerbosService) CheckResourceBatch(ctx context.Context, req *requestv1.CheckResourceBatchRequest) (*responsev1.CheckResourceBatchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CheckResources checks a batch of heterogenous resources.
func (cs *CerbosService) CheckResources(ctx context.Context, req *requestv1.CheckResourcesRequest) (*responsev1.CheckResourcesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cs *CerbosService) checkNumResourcesLimit(n int) error { _ = "STUB: not implemented"; return nil }

func (cs *CerbosService) checkNumActionsLimit(n int) error { _ = "STUB: not implemented"; return nil }

func (CerbosService) ServerInfo(_ context.Context, _ *requestv1.ServerInfoRequest) (*responsev1.ServerInfoResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
