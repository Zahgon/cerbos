// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package local

import (
	"context"
	"fmt"
	"sync"
	"time"

	badgerv4 "github.com/dgraph-io/badger/v4"
	"go.uber.org/zap"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/config"
)

const (
	badgerDiscardRatio      = 0.5
	goroutineResetThreshold = 1 << 16

	Backend          = "local"
	keyLen           = 24
	keyTSStart       = 4
	keyTSEnd         = 10
	KeyByteSizeStart = 20
	// Messages have a hard limit of 8mb in Hub. We hold back 2MB to allow for additional metadata/tolerances etc.
	MaxAllowedBatchSizeBytes = 6291456 // 6MB

)

var (
	AccessLogPrefix   = []byte("aacc")
	DecisionLogPrefix = []byte("adec")
)

func init() {
	audit.RegisterBackend(Backend, func(_ context.Context, confW *config.Wrapper, decisionFilter audit.DecisionLogEntryFilter) (audit.Log, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read local audit log configuration: %w", err)
		}

		return NewLog(conf, decisionFilter)
	})
}

// Log implements the decisionlog interface with Badger as the backing store.
type Log struct {
	logger                   *zap.Logger
	Db                       *badgerv4.DB
	buffer                   chan *badgerv4.Entry
	stopChan                 chan struct{}
	decisionFilter           audit.DecisionLogEntryFilter
	wg                       sync.WaitGroup
	ttl                      time.Duration
	stopOnce                 sync.Once
	bufferSize, maxBatchSize int
	flushInterval            time.Duration
}

func NewLog(conf *Conf, decisionFilter audit.DecisionLogEntryFilter) (*Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:mnd
//nolint:mnd

func (l *Log) batchWriter(maxBatchSize int, flushInterval time.Duration) {
	_ = "STUB: not implemented"
	return
}

// restart the goroutine with a fresh stack

func (l *Log) gc(gcInterval time.Duration) { _ = "STUB: not implemented"; return }

// restart goroutine with a fresh stack

func (l *Log) Backend() string { _ = "STUB: not implemented"; return "" }

func (l *Log) Enabled() bool {
	_ = "STUB: not implemented"

	// ForceWrite forces a write operation and blocks until completion. It is used only by tests.
	return false
}

func (l *Log) ForceWrite() { _ = "STUB: not implemented"; return }

// Restart the batching goroutine

func (l *Log) WriteAccessLogEntry(ctx context.Context, record audit.AccessLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Log) WriteDecisionLogEntry(ctx context.Context, record audit.DecisionLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Log) Write(ctx context.Context, key, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Log) LastNAccessLogEntries(ctx context.Context, n uint) audit.AccessLogIterator {
	_ = "STUB: not implemented"
	return *new(audit.AccessLogIterator)
}

func (l *Log) LastNDecisionLogEntries(ctx context.Context, n uint) audit.DecisionLogIterator {
	_ = "STUB: not implemented"
	return *new(audit.DecisionLogIterator)
}

func (l *Log) listLastN(ctx context.Context, prefix []byte, n uint, c collector) {
	_ = "STUB: not implemented"
	return
}

func (l *Log) AccessLogEntriesBetween(ctx context.Context, fromTS, toTS time.Time) audit.AccessLogIterator {
	_ = "STUB: not implemented"
	return *new(audit.AccessLogIterator)
}

func (l *Log) DecisionLogEntriesBetween(ctx context.Context, fromTS, toTS time.Time) audit.DecisionLogIterator {
	_ = "STUB: not implemented"
	return *new(audit.DecisionLogIterator)
}

func (l *Log) listBetweenTimestamps(ctx context.Context, prefix []byte, fromTS, toTS time.Time, c collector) {
	_ = "STUB: not implemented"
	return
}

// stop when we have reached a key larger than the end key

func (l *Log) AccessLogEntryByID(ctx context.Context, id audit.ID) audit.AccessLogIterator {
	_ = "STUB: not implemented"
	return *new(audit.AccessLogIterator)
}

func (l *Log) DecisionLogEntryByID(ctx context.Context, id audit.ID) audit.DecisionLogIterator {
	_ = "STUB: not implemented"
	return *new(audit.DecisionLogIterator)
}

func (l *Log) getByID(ctx context.Context, prefix []byte, id audit.ID, c collector) {
	_ = "STUB: not implemented"
	return
}

func (l *Log) Close() error { _ = "STUB: not implemented"; return nil }

func GenKey(prefix []byte, id audit.IDBytes) []byte { _ = "STUB: not implemented"; return nil }

func GenKeyWithByteSize(prefix []byte, id audit.IDBytes, nBytes int) []byte {
	_ = "STUB: not implemented"
	return nil
}

func genKeyForTime(prefix []byte, ts time.Time) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func minScanKeyForTime(prefix []byte, ts time.Time) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:mnd

func maxScanKeyForTime(prefix []byte, ts time.Time) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:mnd

func scanKeyForTime(prefix []byte, ts time.Time, randFiller byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func maxScanKeyForPrefix(prefix []byte) []byte { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func scanKeyForPrefix(prefix []byte, filler byte) []byte { _ = "STUB: not implemented"; return nil }

type batcher struct {
	db      *badgerv4.DB
	batch   []*badgerv4.Entry
	maxSize int
	ptr     int
}

func newBatcher(db *badgerv4.DB, maxSize int) *batcher { _ = "STUB: not implemented"; return nil }

func (b *batcher) add(entry *badgerv4.Entry) error { _ = "STUB: not implemented"; return nil }

func (b *batcher) flush() error { _ = "STUB: not implemented"; return nil }

func newDBLogger(logger *zap.Logger) badgerv4.Logger {
	_ = "STUB: not implemented"
	return *new(badgerv4.Logger)
}

type zapLogger struct {
	*zap.SugaredLogger
}

func (zl zapLogger) Warningf(msg string, args ...any) { _ = "STUB: not implemented"; return }
