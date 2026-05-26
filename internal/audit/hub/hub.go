// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sourcegraph/conc/pool"
	"go.uber.org/zap"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/audit/local"
	"github.com/cerbos/cerbos/internal/config"
	logsv1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/logs/v1"
)

const (
	Backend = "hub"

	maxAllowedBatchSize = 1024
)

type syncPrefix []byte

var (
	SyncStatusPrefix   = syncPrefix("bs")   // "b" for contiguity with audit log keys in LSM, "s" because "sync"
	AccessSyncPrefix   = syncPrefix("bsac") // these need to be len(4) to correctly reuse `local.GenKey`
	DecisionSyncPrefix = syncPrefix("bsde")
)

func init() {
	audit.RegisterBackend(Backend, func(ctx context.Context, confW *config.Wrapper, decisionFilter audit.DecisionLogEntryFilter) (audit.Log, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read hub audit log configuration: %w", err)
		}

		logger := zap.L().Named("auditlog").With(zap.String("backend", Backend))

		syncer, err := NewIngestSyncer(logger)
		if err != nil {
			return nil, err
		}

		var pipeLog audit.Log
		if conf.PipeOutput.Enabled {
			cons, err := audit.GetBackend(conf.PipeOutput.Backend)
			if err != nil {
				return nil, err
			}

			pipeLog, err = cons(ctx, confW, decisionFilter)
			if err != nil {
				return nil, fmt.Errorf("failed to construct pipe output backend: %w", err)
			}
		}

		return NewLog(conf, decisionFilter, syncer, logger, pipeLog)
	})
}

type options struct {
	maxBatchSize int
}

type Opt func(*options)

func WithMaxBatchSize(maxBatchSize int) Opt { _ = "STUB: not implemented"; return *new(Opt) }

type Log struct {
	syncer          IngestSyncer
	pipeLog         audit.Log
	cancel          context.CancelFunc
	filter          *AuditLogFilter
	oversizedFilter *AuditLogFilter
	pool            *pool.ContextPool
	logger          *zap.Logger
	*local.Log
	minFlushInterval  time.Duration
	flushTimeout      time.Duration
	maxBatchSize      int
	maxBatchSizeBytes int
	numGo             int
}

func NewLog(conf *Conf, decisionFilter audit.DecisionLogEntryFilter, syncer IngestSyncer, logger *zap.Logger, pipeLog audit.Log, opts ...Opt) (*Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Log) WriteAccessLogEntry(ctx context.Context, record audit.AccessLogEntryMaker) error {
	_ = "STUB: not implemented" //nolint:dupl
	return nil
}

func (l *Log) WriteDecisionLogEntry(ctx context.Context, record audit.DecisionLogEntryMaker) error {
	_ = "STUB: not implemented" //nolint:dupl
	return nil
}

func (l *Log) syncLoop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *Log) schedule() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (l *Log) streamLogs() error {
	_ = "STUB: not implemented"
	// We use two streams: one for access logs, and one for decision logs, as this allows us to
	// avoid the penalty of per-key string inspection when inferring the type down the line.
	return nil
}

var keysPool = &sync.Pool{}

func (l *Log) streamPrefix(ctx context.Context, kind logsv1.IngestBatch_EntryKind, prefix syncPrefix) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forcetypeassert

// Retrieve byte-size from key.
// If we find a legacy key (without the message size stored in the final 4 bytes), we need
// to establish the message size now. Retrieving and sizing each entry individually is annoyingly
// inefficient, but given that this'll only happen "once" when dealing with old, unprocessed
// events, it's probably OK.
// If it's oversized, we just replace the key in the DB and skip, so that it can be picked up on
// the next iteration.
// TODO: rip this if-else out in the future

//nolint:nestif

// Write the new size-integrated key to Badger
//nolint:exhaustive

// Delete the legacy (potentially oversized) key

// Skip this event for now. The new, correctly filtered event
// will be processed in a future run.

// these legacy keys will be processed again on the next sync run

func (l *Log) syncThenDelete(ctx context.Context, kind logsv1.IngestBatch_EntryKind, syncKeys [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Log) getIngestBatchEntries(syncKeys [][]byte, kind logsv1.IngestBatch_EntryKind) ([]*logsv1.IngestBatch_Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Log) Backend() string { _ = "STUB: not implemented"; return "" }

func (l *Log) Close() (outErr error) { _ = "STUB: not implemented"; return nil }
