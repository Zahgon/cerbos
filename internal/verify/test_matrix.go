// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	effectv1 "github.com/cerbos/cerbos/api/genpb/cerbos/effect/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

type testMatrixKey struct {
	Principal string
	Resource  string
}

type testMatrixExpectations struct {
	actions map[string]effectv1.Effect
	outputs map[string]*policyv1.Test_OutputEntries
}

type testMatrixElement struct {
	Expected *testMatrixExpectations
	testMatrixKey
}

func (r *testSuiteRun) buildTestMatrix(table *policyv1.TestTable) ([]testMatrixElement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) buildExpectationLookup(table *policyv1.TestTable) (map[testMatrixKey]*testMatrixExpectations, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func outputExpectations(expectation *policyv1.TestTable_Expectation) map[string]*policyv1.Test_OutputEntries {
	_ = "STUB: not implemented"
	return nil
}

func mergeExpectations(key testMatrixKey, target *testMatrixExpectations, actions map[string]effectv1.Effect, outputs map[string]*policyv1.Test_OutputEntries) (*testMatrixExpectations, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeEffectExpectations(key testMatrixKey, target, source map[string]effectv1.Effect) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeOutputExpectations(key testMatrixKey, target, source map[string]*policyv1.Test_OutputEntries) (map[string]*policyv1.Test_OutputEntries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeOutputEntries(key testMatrixKey, action string, target, source *policyv1.Test_OutputEntries) (*policyv1.Test_OutputEntries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *testSuiteRun) collectFixtures(fixture string, fixtures, groups []string, lookup func(string) ([]string, error)) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildDefaultExpectation(table *policyv1.TestTable) *testMatrixExpectations {
	_ = "STUB: not implemented"
	return nil
}
