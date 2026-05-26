// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package internal

import (
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/jmoiron/sqlx"
)

const (
	defaultRetryMaxAttempts uint = 3
)

// ConnPoolConf holds common SQL connection pool settings.
type ConnPoolConf struct {
	MaxLifetime time.Duration `yaml:"maxLifeTime"`
	MaxIdleTime time.Duration `yaml:"maxIdleTime"`
	MaxOpen     uint          `yaml:"maxOpen"`
	MaxIdle     uint          `yaml:"maxIdle"`
}

func (cc *ConnPoolConf) Configure(db *sqlx.DB) { _ = "STUB: not implemented"; return }

// ConnRetryConf holds common retry settings for establishing a database connection.
type ConnRetryConf struct {
	// MaxAttempts is the maximum number of times to attempt to connect before giving up.
	MaxAttempts uint `yaml:"maxAttempts"`
	// InitialInterval is the initial wait period between retry attempts. Subsequent attempts will be longer depending on the attempt number.
	InitialInterval time.Duration `yaml:"initialInterval"`
	// MaxInterval is the maximum amount of time to wait between retry attempts.
	MaxInterval time.Duration `yaml:"maxInterval"`
}

func (rc *ConnRetryConf) Validate() (outErr error) { _ = "STUB: not implemented"; return nil }

func (rc *ConnRetryConf) BackoffOptions() []backoff.RetryOption {
	_ = "STUB: not implemented"
	return nil
}
