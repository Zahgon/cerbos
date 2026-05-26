// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build tests

package blob

import (
	"context" //nolint:gosec
	"sync/atomic"
	"testing"
	"time"

	"gocloud.dev/blob"
)

const (
	seaweedUsername = "weedadmin"
	seaweedPassword = "weedadmin"
)

const timeout = 5 * time.Minute

func SeaweedFSBucketURL(bucketName, endpoint string) string { _ = "STUB: not implemented"; return "" }

type UploadParam struct {
	BucketURL    string
	BucketPrefix string
	Username     string
	Password     string //nolint:gosec
	Directory    string
}

//nolint:revive
func CopyDirToBucket(tb testing.TB, ctx context.Context, param UploadParam) *blob.Bucket {
	_ = "STUB: not implemented"
	return nil
}

func newSeaweedFSBucket(t *testing.T, seaweedFS *SeaweedFS, path, prefix string) *blob.Bucket {
	_ = "STUB: not implemented"
	return nil
}

//nolint:revive
func uploadDirToBucket(tb testing.TB, ctx context.Context, dir string, bucket *blob.Bucket) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bucketAdd(tb testing.TB, bucket *blob.Bucket, key string, data []byte) {
	_ = "STUB: not implemented"
	return
}

//nolint:gosec

func bucketDelete(tb testing.TB, bucket *blob.Bucket, key string) {
	_ = "STUB: not implemented"
	return
}

type SeaweedFS struct {
	endpoint string
	buckets  atomic.Int64
}

func (s *SeaweedFS) CreateBucket(t *testing.T) string { _ = "STUB: not implemented"; return "" }

func StartSeaweedFS(t *testing.T) *SeaweedFS { _ = "STUB: not implemented"; return nil }

// exponential backoff-retry, because the application in the container might not be ready to accept connections yet

//nolint:gosec
