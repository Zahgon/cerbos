// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package telemetry

import "time"

const (
	confKey               = "telemetry"
	defaultReportInterval = 1 * time.Hour
)

// Conf for telemetry reporting.
type Conf struct {
	// StateDir is used to persist state to avoid repeatedly sending the data over and over again.
	StateDir string `yaml:"stateDir" conf:",example=${HOME}/.config/cerbos"`
	// Disabled sets whether telemetry collection is disabled or not.
	Disabled bool `yaml:"disabled" conf:",example=false"`
	// ReportInterval is the interval between telemetry pings.
	ReportInterval time.Duration `yaml:"reportInterval" conf:",example=1h"`
}

func (c *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (c *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

func (c *Conf) Validate() (errs error) { _ = "STUB: not implemented"; return nil }
