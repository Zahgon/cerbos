// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package blob

import (
	"context"
	"time"

	"github.com/cerbos/cerbos/internal/storage"
)

const (
	confKey                = storage.ConfKey + "." + DriverName
	defaultDownloadTimeout = 60 * time.Second
	defaultRequestTimeout  = 5 * time.Second
)

// Conf is required (if driver is set to 'blob') configuration for cloud storage driver.
// +desc=This section is required only if storage.driver is blob.
type Conf struct {
	// DownloadTimeout specifies the timeout for downloading from cloud storage.
	DownloadTimeout *time.Duration `yaml:"downloadTimeout,omitempty" conf:",example=30s"`
	// RequestTimeout specifies the timeout for an HTTP request.
	RequestTimeout *time.Duration `yaml:"requestTimeout,omitempty" conf:",example=10s"`
	// Bucket URL (Examples: s3://my-bucket?region=us-west-1 gs://my-bucket).
	Bucket string `yaml:"bucket" conf:"required,example=\"s3://my-bucket-name?region=us-east-2\""`
	// Prefix specifies a subdirectory to download.
	Prefix string `yaml:"prefix,omitempty" conf:",example=policies"`
	// WorkDir is the local path to check out policies to.
	WorkDir string `yaml:"workDir" conf:",example=${HOME}/tmp/cerbos/work"`
	// UpdatePollInterval specifies the interval to poll the cloud storage. Set to 0 to disable.
	UpdatePollInterval time.Duration `yaml:"updatePollInterval" conf:",example=15s"`
}

func (conf *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (conf *Conf) Validate() error { _ = "STUB: not implemented"; return nil }

func pd(d time.Duration) *time.Duration { _ = "STUB: not implemented"; return nil }

func (conf *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

func (conf *Conf) getCloneCtx(parent context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func GetConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }
