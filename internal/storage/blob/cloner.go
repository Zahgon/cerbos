// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package blob

import (
	"context"

	"go.uber.org/zap"
	"gocloud.dev/blob"
)

type Cloner struct {
	bucket *blob.Bucket
	fs     FS
	log    *zap.SugaredLogger
	state  map[string][]string
}

func NewCloner(bucket *blob.Bucket, dir string) (*Cloner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type info struct {
	etag string
	file string
}

type CloneResult struct {
	all            map[string][]string
	addedOrUpdated []info
	deleted        []info
}

func (cr *CloneResult) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (c *Cloner) Clone(ctx context.Context) (*CloneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cloner) downloadToFile(ctx context.Context, key, file string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

func (c *Cloner) Clean() error { _ = "STUB: not implemented"; return nil }
