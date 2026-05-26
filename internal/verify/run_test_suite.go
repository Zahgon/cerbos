// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"context"
	"errors"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

var errUsedDefaultNow = errors.New("a policy used a time-based condition, but `now` was not provided in the test options")

func runTestSuite(ctx context.Context, eng Checker, filter *testFilter, file string, suite *policyv1.TestSuite, fixture *TestFixture, trace, skipBatching bool) *policyv1.TestResults_Suite {
	_ = "STUB: not implemented"
	return nil
}

type testSuiteRun struct {
	Suite           *policyv1.TestSuite
	Fixture         *TestFixture
	PrincipalGroups map[string][]string
	ResourceGroups  map[string][]string
}

func (r *testSuiteRun) checkUniqueTestNames() error { _ = "STUB: not implemented"; return nil }

func (r *testSuiteRun) getTests() ([]*policyv1.Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) buildTests(table *policyv1.TestTable) ([]*policyv1.Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) buildTest(table *policyv1.TestTable, matrixElement testMatrixElement) (*policyv1.Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) lookupPrincipal(name string) (*enginev1.Principal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) lookupPrincipalGroup(name string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) lookupResource(name string) (*enginev1.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) lookupResourceGroup(name string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) lookupAuxData(name string) (*enginev1.AuxData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func runTest(ctx context.Context, eng Checker, test *policyv1.Test, actions []string, trace bool) map[string]*policyv1.TestResults_Details {
	_ = "STUB: not implemented"
	return nil
}

func performCheck(ctx context.Context, eng Checker, inputs []*enginev1.CheckInput, options *policyv1.TestOptions, trace bool) (_ []*enginev1.CheckOutput, traces []*enginev1.Trace, _ error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
