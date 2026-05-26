// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/cenkalti/backoff/v5"
	logsv1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/logs/v1"
	"github.com/cerbos/cloud-api/logcap"
)

const (
	initialBackoffInterval     = 1 * time.Second
	backoffRandomizationFactor = 0.5
	backoffMultiplier          = 1.5
	maxBackoffInterval         = 2 * time.Minute
)

type ErrIngestBackoff struct {
	underlying error
	Backoff    time.Duration
}

func (e ErrIngestBackoff) Error() string { _ = "STUB: not implemented"; return "" }

type IngestSyncer interface {
	Sync(context.Context, *logsv1.IngestBatch) error
}

type wrappedBackOff struct {
	backoff.BackOff
	mu sync.Mutex
}

func (exp *wrappedBackOff) NextBackOff() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (exp *wrappedBackOff) Reset() { _ = "STUB: not implemented"; return }

type Impl struct {
	client *logcap.Client
	log    *zap.Logger
	wbo    *wrappedBackOff
}

func NewIngestSyncer(logger *zap.Logger) (*Impl, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *Impl) Sync(ctx context.Context, batch *logsv1.IngestBatch) error {
	_ = "STUB: not implemented"
	return nil
}

// Hard failure: use exponential backoff
