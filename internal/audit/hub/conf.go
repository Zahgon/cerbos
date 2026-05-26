// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	"fmt"
	"time"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/audit/local"
)

const (
	confKey = audit.ConfKey + ".hub"

	defaultMinFlushInterval  = 5 * time.Second
	defaultFlushTimeout      = 5 * time.Second
	defaultNumGoRoutines     = 4
	defaultMaxBatchSizeBytes = 4194304 // 4MB

	minMinFlushInterval = 2 * time.Second
	maxFlushTimeout     = 10 * time.Second
	// Arbitrary figure to account for additional metadata in the batch message as we only track the size of each entry at write time.
	// It's not the end of the world if the batch size exceeds the limit (due to this number being set too low), but
	// it reduces the chance of that happening.
	BatchSizeToleranceBytes = 128
)

var (
	errInvalidFlushInterval = fmt.Errorf("flushInterval must be at least %s", minMinFlushInterval)
	errInvalidFlushTimeout  = fmt.Errorf("flushTimeout cannot be more than %s", maxFlushTimeout)
)

type Conf struct {
	PipeOutput PipeOutputConf `yaml:"pipeOutput"`
	// Mask defines a list of attributes to exclude from the audit logs, specified as lists of JSONPaths
	Mask       MaskConf `yaml:"mask"`
	local.Conf `yaml:",inline"`
	Ingest     IngestConf `yaml:"ingest" conf:",ignore"`
}

type IngestConf struct {
	// MaxBatchSizeBytes defines the max cumulative size in bytes for a batch of log entries.
	MaxBatchSizeBytes uint `yaml:"maxBatchSizeBytes" conf:",example=2097152,ignore"`
	// MinFlushInterval is the minimal duration between Ingest requests.
	MinFlushInterval time.Duration `yaml:"minFlushInterval" conf:",example=3s"`
	// FlushTimeout defines the max allowable timeout for each Ingest request.
	FlushTimeout time.Duration `yaml:"flushTimeout" conf:",example=5s"`
	// NumGoRoutines defines the max number of goroutines used when streaming log entries from the local DB.
	NumGoRoutines uint `yaml:"numGoRoutines" conf:",example=8"`
}

type MaskConf struct {
	Peer           []string `yaml:"peer" conf:",example=\n    - address\n    - forwarded_for"`
	Metadata       []string `yaml:"metadata" conf:",example=['authorization']"`
	CheckResources []string `yaml:"checkResources" conf:",example=\n    - inputs[*].principal.attr.foo\n    - inputs[*].auxData\n    - outputs"`
	PlanResources  []string `yaml:"planResources" conf:",example=['input.principal.attr.nestedMap.foo']"`
}

type PipeOutputConf struct {
	Backend string `yaml:"backend" conf:",example=file"`
	Enabled bool   `yaml:"enabled" conf:",example=false"`
}

func (c *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (c *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

func (c *Conf) Validate() (outErr error) { _ = "STUB: not implemented"; return nil }
