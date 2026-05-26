// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package local

import (
	"sync"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
)

type collector interface {
	add([]byte) error
	done(error)
}

type accessLogEntryCollector struct {
	err      error
	buffer   chan *auditv1.AccessLogEntry
	mu       sync.RWMutex
	doneOnce sync.Once
}

func newAccessLogEntryCollector() *accessLogEntryCollector { _ = "STUB: not implemented"; return nil }

func (a *accessLogEntryCollector) add(v []byte) error { _ = "STUB: not implemented"; return nil }

func (a *accessLogEntryCollector) done(err error) { _ = "STUB: not implemented"; return }

func (a *accessLogEntryCollector) Next() (*auditv1.AccessLogEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type decisionLogEntryCollector struct {
	err      error
	buffer   chan *auditv1.DecisionLogEntry
	mu       sync.RWMutex
	doneOnce sync.Once
}

func newDecisionLogEntryCollector() *decisionLogEntryCollector {
	_ = "STUB: not implemented"
	return nil
}

func (d *decisionLogEntryCollector) add(v []byte) error { _ = "STUB: not implemented"; return nil }

// convert old format records to new format
//nolint:staticcheck

func (d *decisionLogEntryCollector) done(err error) { _ = "STUB: not implemented"; return }

func (d *decisionLogEntryCollector) Next() (*auditv1.DecisionLogEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
