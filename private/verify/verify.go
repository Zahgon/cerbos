// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"io/fs"
	"time"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/verify"
	"github.com/cerbos/cerbos/private/compile"
	"github.com/cerbos/cerbos/private/engine"
)

// Files runs tests using the policy files in the given file system.
func Files(ctx context.Context, fsys fs.FS, idx compile.Index, trace bool) (*policyv1.TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bundle runs tests using the given policy bundle.
func Bundle(ctx context.Context, params engine.BundleParams, testsDir string, trace bool) (*policyv1.TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BundleStream runs tests using the given policy bundle and streams results.
func BundleStream(ctx context.Context, params engine.BundleParams, testsDir string, trace bool) (int, <-chan *policyv1.TestResults_Suite, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

type CheckOptions interface {
	Globals() map[string]any
	NowFunc() func() time.Time
	DefaultPolicyVersion() string
	LenientScopeSearch() bool
}

type Checker interface {
	Check(ctx context.Context, inputs []*enginev1.CheckInput, opts CheckOptions) ([]*enginev1.CheckOutput, error)
}

type Opt func(config *verify.Config)

func WithExcludedResourcePolicyFQNs(fqns ...string) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithExcludedPrincipalPolicyFQNs(fqns ...string) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithWorkers(count uint) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithCustomChecker(ctx context.Context, fsys fs.FS, eng Checker, opts ...Opt) (*policyv1.TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type checkFunc func(context.Context, []*enginev1.CheckInput, CheckOptions) ([]*enginev1.CheckOutput, error)

func (f checkFunc) Check(ctx context.Context, inputs []*enginev1.CheckInput, opts ...evaluator.CheckOpt) ([]*enginev1.CheckOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
