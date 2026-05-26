// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	"context"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"

	svcv1 "github.com/cerbos/cerbos/api/genpb/authzen/authorization/v1"
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/auxdata"
	"github.com/cerbos/cerbos/internal/engine"
)

var _ svcv1.AuthorizationServiceServer = (*AuthzenAuthorizationService)(nil)

// AuthzenAuthorizationService implements the policy checking service.
type AuthzenAuthorizationService struct {
	eng     *engine.Engine
	auxData *auxdata.AuxData
	*svcv1.UnimplementedAuthorizationServiceServer
	reqLimits RequestLimits
}

func NewAuthzenAuthorizationService(eng *engine.Engine, auxData *auxdata.AuxData, reqLimits RequestLimits) *AuthzenAuthorizationService {
	_ = "STUB: not implemented"
	return nil
}

func (aas *AuthzenAuthorizationService) AccessEvaluation(ctx context.Context, r *svcv1.AccessEvaluationRequest) (*svcv1.AccessEvaluationResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract auxData

// Build engine input

// Call engine

// Build CheckResourcesResponse structure for context

// Convert to value for context

const (
	evalSemanticExecuteAll          = "execute_all"
	evalSemanticPermitOnFirstPermit = "permit_on_first_permit"
)

func (aas *AuthzenAuthorizationService) AccessEvaluationBatch(ctx context.Context, r *svcv1.AccessEvaluationBatchRequest) (*svcv1.AccessEvaluationBatchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate total resources

// Merge each evaluation with defaults to create complete requests

// Original index for maintaining order

// evaluations with default values taken into account

// Group evaluations by (subject, auxData) to process efficiently

// Process each group

// Group by resource since a ResourceEntry can have multiple actions

// Validate total resources

// Build engine inputs and track mapping

// Call engine

// Build CheckResourcesResponse for this group

// Convert to value for context

// Map results back to original positions

// TODO(db): share this function with CerbosService?
func buildCheckResourcesResponse(requestID string, inputs []*enginev1.CheckInput, outputs []*enginev1.CheckOutput, includeMeta bool) *responsev1.CheckResourcesResponse {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

func merge[T any](defaults, override *T) *T { _ = "STUB: not implemented"; return nil }

func cerbosProp(s string) string { _ = "STUB: not implemented"; return "" }

func lookup[T any](m map[string]*T, k string) *T { _ = "STUB: not implemented"; return nil }

func lookupOrEmptyString(m map[string]*structpb.Value, k string) string {
	_ = "STUB: not implemented"
	return ""
}

func (aas *AuthzenAuthorizationService) checkTotalLimit(n int) error {
	_ = "STUB: not implemented"
	return nil
}

func (aas *AuthzenAuthorizationService) checkNumResourcesLimit(n int) error {
	_ = "STUB: not implemented"
	return nil
}

func (aas *AuthzenAuthorizationService) checkNumActionsLimit(n int) error {
	_ = "STUB: not implemented"
	return nil
}

func toResource(res *svcv1.Resource) *enginev1.Resource { _ = "STUB: not implemented"; return nil }

func toPrincipal(subj *svcv1.Subject) *enginev1.Principal { _ = "STUB: not implemented"; return nil }

func messageToValue(msg protoreflect.Message) (*structpb.Value, error) {
	_ = "STUB: not implemented"
	// If the message is already a structpb.Value, return it as-is
	return nil, nil
}

func mkStructPBValue(repeated bool, v protoreflect.Value, f func(v protoreflect.Value) *structpb.Value) *structpb.Value {
	_ = "STUB: not implemented"
	return nil
}

func valueToStructValue(fd protoreflect.FieldDescriptor, v protoreflect.Value) (*structpb.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func valueToMessage(value *structpb.Value, msg protoreflect.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func structToMessage(s *structpb.Struct, msg protoreflect.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func structValueToProtoValue(value *structpb.Value, fd protoreflect.FieldDescriptor, msg protoreflect.Message) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

// Regular message

// default to first enum value

func metaRequested(m map[string]*structpb.Value) bool { _ = "STUB: not implemented"; return false }

func (aas *AuthzenAuthorizationService) extractAuxData(ctx context.Context, m map[string]*structpb.Value) (*enginev1.AuxData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (aas *AuthzenAuthorizationService) Metadata(ctx context.Context, _ *svcv1.MetadataRequest) (*svcv1.MetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
