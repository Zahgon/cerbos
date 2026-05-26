// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"context"
	"io"
	"net/url"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
	"google.golang.org/protobuf/types/known/structpb"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/observability/logging"
)

const (
	Directory = "_schemas"
	URLScheme = "cerbos"
)

var alwaysValidResult = &ValidationResult{Reject: false}

type ValidationResult struct {
	Errors ValidationErrorList
	Reject bool
}

func (vr *ValidationResult) add(errs ...ValidationError) { _ = "STUB: not implemented"; return }

type Manager interface {
	managerLoader
	ValidateCheckInput(context.Context, *policyv1.Schemas, *enginev1.CheckInput) (*ValidationResult, error)
	ValidatePlanResourcesInput(context.Context, *policyv1.Schemas, *enginev1.PlanResourcesInput) (*ValidationResult, error)
}

type managerLoader interface {
	LoadSchema(context.Context, string) (*jsonschema.Schema, error)
}

type Loader interface {
	LoadSchema(context.Context, string) (io.ReadCloser, error)
}

type Resolver func(context.Context, string) (io.ReadCloser, error)

func NewNopManager() NopManager { _ = "STUB: not implemented"; return *new(NopManager) }

type NopManager struct{}

func (NopManager) ValidateCheckInput(_ context.Context, _ *policyv1.Schemas, _ *enginev1.CheckInput) (*ValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (NopManager) ValidatePlanResourcesInput(_ context.Context, _ *policyv1.Schemas, _ *enginev1.PlanResourcesInput) (*ValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (NopManager) LoadSchema(_ context.Context, _ string) (*jsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewStatic(schemas map[uint64]*policyv1.Schemas, rawSchemas map[string]*runtimev1.RuleTable_JSONSchema) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

func NewStaticFromConf(conf *Conf, schemas map[uint64]*policyv1.Schemas, rawSchemas map[string]*runtimev1.RuleTable_JSONSchema) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

type StaticManager struct {
	conf            *Conf
	log             *logging.Logger
	compiledSchemas map[string]*jsonschema.Schema
	loader          managerLoader
}

func (m *StaticManager) preCompileSchemas(schemas map[uint64]*policyv1.Schemas, rawSchemas map[string]*runtimev1.RuleTable_JSONSchema) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *StaticManager) LoadSchema(ctx context.Context, url string) (*jsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *StaticManager) ValidateCheckInput(ctx context.Context, schemas *policyv1.Schemas, input *enginev1.CheckInput) (*ValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *StaticManager) ValidatePlanResourcesInput(ctx context.Context, schemas *policyv1.Schemas, input *enginev1.PlanResourcesInput) (*ValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resource attributes are optional for query planning, so ignore errors from required properties

func (m *StaticManager) validate(ctx context.Context, schemas *policyv1.Schemas, principalAttr, resourceAttr map[string]*structpb.Value, actions []string, resourceErrorFilter validationErrorFilter) (*ValidationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *StaticManager) validateAttr(ctx context.Context, src ErrSource, schemaRef *policyv1.Schemas_Schema, attr map[string]*structpb.Value, actions []string, errorFilter validationErrorFilter) error {
	_ = "STUB: not implemented"
	return nil
}

// check whether the current actions are excluded from validation

func filterActionsToValidate(ignore, actions []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func StaticResolver(loader Loader) Resolver { _ = "STUB: not implemented"; return *new(Resolver) }

func loadCerbosURL(ctx context.Context, u *url.URL, loader Loader) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type notFoundErr struct {
	// url is the URL failed to load
	url string
	// scheme is the scheme of the URL without any colons or slashes
	scheme string
	// fullPath is the resolved file system path.
	fullPath string
}

func (e notFoundErr) Error() string { _ = "STUB: not implemented"; return "" }
