// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"regexp"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

const (
	SkipReasonName            = "Test name did not match the provided pattern"
	SkipReasonResource        = "Resource matched a policy that was excluded from the bundle"
	SkipReasonPrincipal       = "Principal matched a policy that was excluded from the bundle"
	SkipReasonFilterSuite     = "Suite did not match the test filter"
	SkipReasonFilterTest      = "Test did not match the test filter"
	SkipReasonFilterPrincipal = "Principal did not match the test filter"
	SkipReasonFilterResource  = "Resource did not match the test filter"
	SkipReasonFilterAction    = "No actions matched the test filter"
)

func IsFilterSkipReason(reason string) bool { _ = "STUB: not implemented"; return false }

type testFilter struct {
	excludedResourcePolicyFQNs  map[string]struct{}
	excludedPrincipalPolicyFQNs map[string]struct{}
	includedTestNamesRegexp     *regexp.Regexp
	filter                      *FilterConfig
}

func newTestFilter(conf *Config) (*testFilter, error) { _ = "STUB: not implemented"; return nil, nil }

// Apply checks if the filter matches the given test, returning nil if the test should be run or a "skipped" result otherwise.
func (f *testFilter) Apply(test *policyv1.Test, suite *policyv1.TestSuite) *policyv1.TestResults_Details {
	_ = "STUB: not implemented"
	return nil
}

func (f *testFilter) shouldRunTestNamed(name string) bool { _ = "STUB: not implemented"; return false }

func (f *testFilter) shouldRunTestForResource(resource *enginev1.Resource, options *policyv1.TestOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *testFilter) shouldRunTestForPrincipal(principal *enginev1.Principal, options *policyv1.TestOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func policyVersion(fixture interface{ GetPolicyVersion() string }, options *policyv1.TestOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func scope(fixture interface{ GetScope() string }, options *policyv1.TestOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *testFilter) matchesFilterSuite(name string) bool { _ = "STUB: not implemented"; return false }

func (f *testFilter) matchesFilterTest(name string) bool { _ = "STUB: not implemented"; return false }

func (f *testFilter) matchesFilterPrincipal(principalKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *testFilter) matchesFilterResource(resourceKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *testFilter) matchesFilterActions(actions []string) bool {
	_ = "STUB: not implemented"
	return false
}

// partitionActions returns actions split into matched and skipped slices.
// If no action filter is configured, all actions are returned as matched.
func (f *testFilter) partitionActions(actions []string) (matched, skipped []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// matchesAnyGlob checks if the value matches any of the provided glob patterns.
func matchesAnyGlob(globs []string, value string) bool { _ = "STUB: not implemented"; return false }
