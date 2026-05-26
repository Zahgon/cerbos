// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

func checkGroupDefinitions[G any](groups map[string]G, members func(G) []string, exists func(string) bool) (map[string][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func principalGroupMembers(group *policyv1.TestFixtureGroup_Principals) []string {
	_ = "STUB: not implemented"
	return nil
}

func resourceGroupMembers(group *policyv1.TestFixtureGroup_Resources) []string {
	_ = "STUB: not implemented"
	return nil
}

func existsInMap[F any](fixtures map[string]F) func(string) bool {
	_ = "STUB: not implemented"
	return nil
}

func existsFromLookup[F any](lookup func(string) (F, error)) func(string) bool {
	_ = "STUB: not implemented"
	return nil
}
