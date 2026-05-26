// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package disk

import (
	"context"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/storage/index"
)

const (
	defaultBufferSize uint = 8
	// defaultCooldownPeriod is the amount of time to wait before triggering a notification to update.
	// This is necessary because many file update events can fire even for a seemingly simple operation on a file.
	// We want the system to settle down before performing an expensive update operation.
	defaultCooldownPeriod = 2 * time.Second
)

func watchDir(ctx context.Context, dir string, idx index.Index, sub *storage.SubscriptionManager, cooldownPeriod time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// We need to manually traverse the tree to add all directories because fsnotify package does not support recursive monitoring.
// See https://github.com/fsnotify/fsnotify/issues/18 for more details.

//nolint:gosec

type dirWatch struct {
	lastEventTime time.Time
	watcher       *fsnotify.Watcher
	idx           index.Index
	log           *zap.SugaredLogger
	eventBatch    map[string]struct{}
	*storage.SubscriptionManager
	dir            string
	cooldownPeriod time.Duration
	mu             sync.RWMutex
}

func (dw *dirWatch) listen(ctx context.Context) { _ = "STUB: not implemented"; return }

func (dw *dirWatch) processEvent(event fsnotify.Event) { _ = "STUB: not implemented"; return }

func (dw *dirWatch) processError(err error) { _ = "STUB: not implemented"; return }

func (dw *dirWatch) triggerUpdate() { _ = "STUB: not implemented"; return }

// The deleted files need to be processed first because we could have a duplicate definition when a file is renamed otherwise.

// We need to manually add newly created directories.
// See https://github.com/fsnotify/fsnotify/issues/18 for more details.

func (dw *dirWatch) shouldUpdate() bool { _ = "STUB: not implemented"; return false }

// TODO: use ReadPolicyFromFile instead.
func readPolicy(path string) (*policyv1.Policy, error) { _ = "STUB: not implemented"; return nil, nil }
