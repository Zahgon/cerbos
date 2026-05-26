// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package schema

const (
	confKey            = "schema"
	defaultEnforcement = EnforcementNone
	defaultCacheSize   = 1024
)

// Conf is optional configuration for schema validation.
type Conf struct {
	// Enforcement defines level of the validations. Possible values are none, warn, reject.
	Enforcement Enforcement `yaml:"enforcement" conf:",example=reject"`
	// CacheSize defines the number of schemas to cache in memory.
	CacheSize uint `yaml:"cacheSize" conf:",example=1024"`
}

func (c *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (c *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

// Enforcement level for schema validation.
type Enforcement string

const (
	EnforcementNone   Enforcement = "none"   // No enforcement made.
	EnforcementWarn   Enforcement = "warn"   // In case schema is not validated, display a warning.
	EnforcementReject Enforcement = "reject" // In case schema is not validated, reject.
)

func NewConf(enforcement Enforcement) *Conf { _ = "STUB: not implemented"; return nil }

func GetConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }
