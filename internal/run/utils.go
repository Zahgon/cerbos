// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package run

import (
	"context"
	"net/http"
	"time"
)

const (
	requestTimeout = 100 * time.Millisecond
	retryInterval  = 141 * time.Millisecond
)

func WaitForReady(ctx context.Context, errors <-chan error, client *http.Client, httpAddr string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkHealth(client *http.Client, healthURL string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec
