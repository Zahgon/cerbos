// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"

func addResult(suite *policyv1.TestResults_Suite, name *policyv1.Test_TestName, action string, details *policyv1.TestResults_Details) {
	_ = "STUB: not implemented"
	return
}

func addTestCase(suite *policyv1.TestResults_Suite, name string) *policyv1.TestResults_TestCase {
	_ = "STUB: not implemented"
	return nil
}

func addPrincipal(testCaseResult *policyv1.TestResults_TestCase, name string) *policyv1.TestResults_Principal {
	_ = "STUB: not implemented"
	return nil
}

func addResource(principal *policyv1.TestResults_Principal, name string) *policyv1.TestResults_Resource {
	_ = "STUB: not implemented"
	return nil
}

func addAction(resource *policyv1.TestResults_Resource, name string) *policyv1.TestResults_Action {
	_ = "STUB: not implemented"
	return nil
}
