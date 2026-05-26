// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"io/fs"
	"sync"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/evaluator"
)

const (
	parallelismThreshold = 5
	extraWorkersOverCPUs = 4
)

type Config struct {
	ExcludedResourcePolicyFQNs  map[string]struct{}
	ExcludedPrincipalPolicyFQNs map[string]struct{}
	Filter                      *FilterConfig
	IncludedTestNamesRegexp     string
	Workers                     uint
	Trace                       bool
	SkipBatching                bool
}

type Checker interface {
	Check(ctx context.Context, inputs []*enginev1.CheckInput, opts ...evaluator.CheckOpt) ([]*enginev1.CheckOutput, error)
}

func Verify(ctx context.Context, fsys fs.FS, eng Checker, conf Config) (*policyv1.TestResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyStream runs test suites and streams results as each suite completes.
// It returns the number of test suites, a channel of results, and any setup error.
// Callers may cancel the context to stop workers early and avoid unnecessary work.
func VerifyStream(ctx context.Context, fsys fs.FS, eng Checker, conf Config) (int, <-chan *policyv1.TestResults_Suite, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func discoverTestFiles(ctx context.Context, fsys fs.FS) (suiteDefs []string, fixtureDefs map[string]struct{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type fixtureCache struct {
	fsys     fs.FS
	defs     map[string]struct{}
	fixtures map[string]*TestFixture
	mu       sync.Mutex
}

func newFixtureCache(fsys fs.FS, defs map[string]struct{}) *fixtureCache {
	_ = "STUB: not implemented"
	return nil
}

func (c *fixtureCache) get(path string) (*TestFixture, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveWorkerCount(configured uint, numSuites int) int { _ = "STUB: not implemented"; return 0 }

func runConcurrent(ctx context.Context, suiteDefs []string, workers int, runSuite func(string) *policyv1.TestResults_Suite, results chan<- *policyv1.TestResults_Suite) {
	_ = "STUB: not implemented"
	return
}

func appendSuiteResult(results *policyv1.TestResults, suite *policyv1.TestResults_Suite) {
	_ = "STUB: not implemented"
	return
}

func incrementTally(summary *policyv1.TestResults_Summary, result policyv1.TestResults_Result, delta uint32) {
	_ = "STUB: not implemented"
	return
}

func addTally(summary *policyv1.TestResults_Summary, result policyv1.TestResults_Result) *policyv1.TestResults_Tally {
	_ = "STUB: not implemented"
	return nil
}
