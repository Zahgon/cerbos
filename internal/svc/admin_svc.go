// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"

	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	svcv1 "github.com/cerbos/cerbos/api/genpb/cerbos/svc/v1"
	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/storage"
)

var _ svcv1.CerbosAdminServiceServer = (*CerbosAdminService)(nil)

var (
	errAuthRequired = status.Error(codes.Unauthenticated, "authentication required")
	authSep         = []byte(":")
)

// CerbosAdminService implements the Cerbos administration service.
type CerbosAdminService struct {
	sfGroup  singleflight.Group
	store    storage.Store
	auditLog audit.Log
	*svcv1.UnimplementedCerbosAdminServiceServer
	adminUser       string
	adminPasswdHash []byte
}

func NewCerbosAdminService(store storage.Store, auditLog audit.Log, adminUser string, adminPasswdHash []byte) *CerbosAdminService {
	_ = "STUB: not implemented"
	return nil
}

func (cas *CerbosAdminService) AddOrUpdatePolicy(ctx context.Context, req *requestv1.AddOrUpdatePolicyRequest) (*responsev1.AddOrUpdatePolicyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) AddOrUpdateSchema(ctx context.Context, req *requestv1.AddOrUpdateSchemaRequest) (*responsev1.AddOrUpdateSchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) InspectPolicies(ctx context.Context, req *requestv1.InspectPoliciesRequest) (*responsev1.InspectPoliciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filters are not scalable for non-mutable stores.

func (cas *CerbosAdminService) ListPolicies(ctx context.Context, req *requestv1.ListPoliciesRequest) (*responsev1.ListPoliciesResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We've historically supported ListPolicies on non-mutable stores, but later introduced filters are not scalable.
// Therefore, if any of the filters in question are passed and the store is not mutable, we reject the request.

func (cas *CerbosAdminService) GetPolicy(ctx context.Context, req *requestv1.GetPolicyRequest) (*responsev1.GetPolicyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:dupl
func (cas *CerbosAdminService) DeletePolicy(ctx context.Context, req *requestv1.DeletePolicyRequest) (*responsev1.DeletePolicyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:dupl
func (cas *CerbosAdminService) DisablePolicy(ctx context.Context, req *requestv1.DisablePolicyRequest) (*responsev1.DisablePolicyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) EnablePolicy(ctx context.Context, req *requestv1.EnablePolicyRequest) (*responsev1.EnablePolicyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) ListSchemas(ctx context.Context, _ *requestv1.ListSchemasRequest) (*responsev1.ListSchemasResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) GetSchema(ctx context.Context, req *requestv1.GetSchemaRequest) (*responsev1.GetSchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) DeleteSchema(ctx context.Context, req *requestv1.DeleteSchemaRequest) (*responsev1.DeleteSchemaResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) PurgeStoreRevisions(ctx context.Context, req *requestv1.PurgeStoreRevisionsRequest) (*responsev1.PurgeStoreRevisionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cas *CerbosAdminService) ReloadStore(ctx context.Context, req *requestv1.ReloadStoreRequest) (*responsev1.ReloadStoreResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

func (cas *CerbosAdminService) ListAuditLogEntries(req *requestv1.ListAuditLogEntriesRequest, stream svcv1.CerbosAdminService_ListAuditLogEntriesServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (cas *CerbosAdminService) getAuditLogStream(ctx context.Context, req *requestv1.ListAuditLogEntriesRequest) (auditLogStream, error) {
	_ = "STUB: not implemented"
	return *new(auditLogStream), nil
}

type auditLogStream func() (*responsev1.ListAuditLogEntriesResponse, error)

func mkAccessLogStream(it audit.AccessLogIterator) auditLogStream {
	_ = "STUB: not implemented"
	return *new(auditLogStream)
}

func mkDecisionLogStream(it audit.DecisionLogIterator) auditLogStream {
	_ = "STUB: not implemented"
	return *new(auditLogStream)
}

func (cas *CerbosAdminService) checkCredentials(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd
