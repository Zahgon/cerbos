// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit

import (
	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
)

type DecisionLogEntryFilter func(*auditv1.DecisionLogEntry) *auditv1.DecisionLogEntry

func NewDecisionLogEntryFilter() (DecisionLogEntryFilter, error) {
	_ = "STUB: not implemented"
	return *new(DecisionLogEntryFilter), nil
}

func NewDecisionLogEntryFilterFromConf(conf *Conf) DecisionLogEntryFilter {
	_ = "STUB: not implemented"
	return *new(DecisionLogEntryFilter)
}

func buildCheckResourcesFilter(f CheckResourcesFilter) func(*auditv1.DecisionLogEntry_CheckResources) *auditv1.DecisionLogEntry_CheckResources {
	_ = "STUB: not implemented"
	return nil
}

func buildPlanResourcesFilter(f PlanResourcesFilter) func(*auditv1.DecisionLogEntry_PlanResources) *auditv1.DecisionLogEntry_PlanResources {
	_ = "STUB: not implemented"
	return nil
}
