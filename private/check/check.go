// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
	"io/fs"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	internalengine "github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/verify"
	"github.com/cerbos/cerbos/private/compile"
)

type TestFixtureGetter struct {
	fsys  fs.FS
	cache map[string]*verify.TestFixture
}

func NewTestFixtureGetter(fsys fs.FS) *TestFixtureGetter { _ = "STUB: not implemented"; return nil }

func (g *TestFixtureGetter) PreCacheTestFixtures() error { _ = "STUB: not implemented"; return nil }

// We need to search the filesystem for any directory `**/testdata`, omitting nested matches, e.g. `**/testdata/**/testdata`

func (g *TestFixtureGetter) LoadTestFixture(path string) (fixture *verify.TestFixture) {
	_ = "STUB: not implemented"
	return nil
}

type TestFixtureCtx struct {
	Fixture *verify.TestFixture
	Path    string
}

func (g *TestFixtureGetter) GetAllTestFixtures() []*TestFixtureCtx {
	_ = "STUB: not implemented"
	return nil
}

func Check(ctx context.Context, conf *evaluator.Conf, idx compile.Index, inputs []*enginev1.CheckInput) ([]*enginev1.CheckOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CheckOutputWithTraces struct {
	CheckOutput *enginev1.CheckOutput
	Traces      []*enginev1.Trace
}

func CheckWithTraces(ctx context.Context, conf *evaluator.Conf, idx compile.Index, inputs []*enginev1.CheckInput) ([]CheckOutputWithTraces, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newEngine(ctx context.Context, conf *evaluator.Conf, idx compile.Index) (*internalengine.Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
