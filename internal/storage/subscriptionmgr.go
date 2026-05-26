// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package storage

import (
	"context"
	"sync"
	"testing"
	"time"
)

const eventBufferSize = 32

type SubscriptionManager struct {
	eventChan   chan Event
	subscribers map[string]Subscriber
	mu          sync.RWMutex
	once        sync.Once
}

func NewSubscriptionManager(ctx context.Context) *SubscriptionManager {
	_ = "STUB: not implemented"
	return nil
}

func (sm *SubscriptionManager) handleEvents(ctx context.Context) { _ = "STUB: not implemented"; return }

func (sm *SubscriptionManager) distributeEvent(evt Event) { _ = "STUB: not implemented"; return }

// TODO(cell) Use work pool to notify multiple subscribers in parallel.

// Notify sends the events to all subscribers.
func (sm *SubscriptionManager) NotifySubscribers(events ...Event) {
	_ = "STUB: not implemented"
	return
}

// TODO(cell) drop event if not published within a reasonable time period.

func (sm *SubscriptionManager) Subscribe(s Subscriber) { _ = "STUB: not implemented"; return }

func (sm *SubscriptionManager) Unsubscribe(s Subscriber) { _ = "STUB: not implemented"; return }

func (sm *SubscriptionManager) shutdown() { _ = "STUB: not implemented"; return }

// TestSubscription is a helper to test subscriptions.
func TestSubscription(s Subscribable) func(*testing.T, time.Duration, ...Event) {
	_ = "STUB: not implemented"
	return nil
}

// sort top‐level Events by PolicyID.hash, then by SchemaFile for schema events

// For events with same PolicyID hash (like schema events), sort by SchemaFile

// sort Dependents by hash

// allow comparing ModuleID despite its unexported field

type subscriber struct {
	stream chan Event
	events []Event
	mu     sync.RWMutex
}

func (s *subscriber) SubscriberID() string { _ = "STUB: not implemented"; return "" }

func (s *subscriber) OnStorageEvent(evt ...Event) { _ = "STUB: not implemented"; return }

func (s *subscriber) Events() []Event { _ = "STUB: not implemented"; return nil }

func (s *subscriber) Clear() { _ = "STUB: not implemented"; return }
