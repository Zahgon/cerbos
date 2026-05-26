// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	"context"
	"io"

	"go.uber.org/zap"

	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/ruletable"
	"github.com/cerbos/cerbos/internal/schema"
	"github.com/cerbos/cerbos/internal/storage"
	bundlev2 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v2"
)

const cerbosSchemaPrefix = schema.URLScheme + ":///"

type RuleTableBundle struct {
	ruleTable *ruletable.RuleTable
}

func OpenRuleTableBundle(opts OpenOpts) (*RuleTableBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptRuleTableBundle(opts OpenOpts, logger *zap.Logger) (*runtimev1.RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rtb *RuleTableBundle) ID() string { _ = "STUB: not implemented"; return "" }

func (*RuleTableBundle) Type() bundlev2.BundleType {
	_ = "STUB: not implemented"
	return *new(bundlev2.BundleType)
}

func (*RuleTableBundle) GetFirstMatch(_ context.Context, _ []namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*RuleTableBundle) GetAll(_ context.Context) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*RuleTableBundle) GetAllMatching(_ context.Context, _ []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rtb *RuleTableBundle) InspectPolicies(ctx context.Context, _ storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rtb *RuleTableBundle) ListPolicyIDs(_ context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rtb *RuleTableBundle) ListSchemaIDs(_ context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rtb *RuleTableBundle) LoadSchema(_ context.Context, path string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (rtb *RuleTableBundle) GetRuleTable() (*ruletable.RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rtb *RuleTableBundle) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (rtb *RuleTableBundle) Release() error { _ = "STUB: not implemented"; return nil }

func (*RuleTableBundle) Close() error { _ = "STUB: not implemented"; return nil }
