// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/sourcegraph/conc/pool"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	"github.com/cerbos/cerbos/internal/config"
)

const (
	KindAccess   = "access"
	KindDecision = "decision"
)

var (
	ErrIteratorClosed = errors.New("iterator closed")

	backendsMu sync.RWMutex
	backends   = map[string]Constructor{}
)

type Info interface {
	Backend() string
	Enabled() bool
}

type Log interface {
	Info
	io.Closer
	WriteAccessLogEntry(context.Context, AccessLogEntryMaker) error
	WriteDecisionLogEntry(context.Context, DecisionLogEntryMaker) error
}

type QueryableLog interface {
	Log
	LastNAccessLogEntries(context.Context, uint) AccessLogIterator
	LastNDecisionLogEntries(context.Context, uint) DecisionLogIterator
	AccessLogEntriesBetween(context.Context, time.Time, time.Time) AccessLogIterator
	DecisionLogEntriesBetween(context.Context, time.Time, time.Time) DecisionLogIterator
	AccessLogEntryByID(context.Context, ID) AccessLogIterator
	DecisionLogEntryByID(context.Context, ID) DecisionLogIterator
}

// AccessLogEntryMaker is a lazy constructor for access log entries.
type AccessLogEntryMaker func() (*auditv1.AccessLogEntry, error)

// DecisionLogEntryMaker is a lazy constructor for decision log entries.
type DecisionLogEntryMaker func() (*auditv1.DecisionLogEntry, error)

type AccessLogIterator interface {
	Next() (*auditv1.AccessLogEntry, error)
}

type DecisionLogIterator interface {
	Next() (*auditv1.DecisionLogEntry, error)
}

// Constructor for backends.
type Constructor func(context.Context, *config.Wrapper, DecisionLogEntryFilter) (Log, error)

// RegisterBackend registers an audit log backend.
func RegisterBackend(name string, cons Constructor) { _ = "STUB: not implemented"; return }

// GetBackend returns the constructor for the given driver.
func GetBackend(name string) (Constructor, error) {
	_ = "STUB: not implemented"
	return *new(Constructor), nil
}

// NewLog creates a new audit log.
func NewLog(ctx context.Context) (Log, error) { _ = "STUB: not implemented"; return *new(Log), nil }

func NewLogFromConf(ctx context.Context, confW *config.Wrapper) (Log, error) {
	_ = "STUB: not implemented"
	return *new(Log), nil
}

// NewNopLog returns an audit log that does nothing.
func NewNopLog() Log { _ = "STUB: not implemented"; return *new(Log) }

func newLogWrapper(conf *Conf, backend Log) *logWrapper { _ = "STUB: not implemented"; return nil }

// logWrapper wraps the backends and enforces the config options.
type logWrapper struct {
	conf    *Conf
	backend Log
	pool    *pool.Pool
}

func (lw *logWrapper) Backend() string { _ = "STUB: not implemented"; return "" }

func (lw *logWrapper) Enabled() bool { _ = "STUB: not implemented"; return false }

func (lw *logWrapper) WriteAccessLogEntry(ctx context.Context, entry AccessLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (lw *logWrapper) WriteDecisionLogEntry(ctx context.Context, entry DecisionLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (lw *logWrapper) Close() error { _ = "STUB: not implemented"; return nil }

type queryableLogWrapper struct {
	*logWrapper
	queryable QueryableLog
}

func (qlw *queryableLogWrapper) LastNAccessLogEntries(ctx context.Context, n uint) AccessLogIterator {
	_ = "STUB: not implemented"
	return *new(AccessLogIterator)
}

func (qlw *queryableLogWrapper) LastNDecisionLogEntries(ctx context.Context, n uint) DecisionLogIterator {
	_ = "STUB: not implemented"
	return *new(DecisionLogIterator)
}

func (qlw *queryableLogWrapper) AccessLogEntriesBetween(ctx context.Context, from, to time.Time) AccessLogIterator {
	_ = "STUB: not implemented"
	return *new(AccessLogIterator)
}

func (qlw *queryableLogWrapper) DecisionLogEntriesBetween(ctx context.Context, from, to time.Time) DecisionLogIterator {
	_ = "STUB: not implemented"
	return *new(DecisionLogIterator)
}

func (qlw *queryableLogWrapper) AccessLogEntryByID(ctx context.Context, id ID) AccessLogIterator {
	_ = "STUB: not implemented"
	return *new(AccessLogIterator)
}

func (qlw *queryableLogWrapper) DecisionLogEntryByID(ctx context.Context, id ID) DecisionLogIterator {
	_ = "STUB: not implemented"
	return *new(DecisionLogIterator)
}

// nopAccessLogIterator implements an AccessLogIterator that always returns nothing.
type nopAccessLogIterator struct{}

func (n nopAccessLogIterator) Next() (*auditv1.AccessLogEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// nopDecisionLogIterator implements a DecisionLogIterator that always returns nothing.
}

type nopDecisionLogIterator struct{}

func (n nopDecisionLogIterator) Next() (*auditv1.DecisionLogEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
