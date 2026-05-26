// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package inspect

import (
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
)

func PolicySets() *PolicySet { _ = "STUB: not implemented"; return nil }

type PolicySet struct {
	results map[string]*responsev1.InspectPoliciesResponse_Result
}

// Inspect inspects the given policy set and caches the inspection related information internally.
func (ps *PolicySet) Inspect(pset *runtimev1.RunnablePolicySet) error {
	_ = "STUB: not implemented"
	return nil
}

// Results returns the final inspection results.
func (ps *PolicySet) Results() (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil,

		// listReferencedAttributes inspects the definitions and rules in the given policy set to find references to the attributes.
		nil
}

func (ps *PolicySet) listReferencedAttributes(pset *runtimev1.RunnablePolicySet) ([]*responsev1.InspectPoliciesResponse_Attribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
