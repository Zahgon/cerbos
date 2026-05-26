// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit

const (
	ConfKey = "audit"
)

// Conf is optional configuration for Audit.
type Conf struct {
	confHolder
}

type confHolder struct {
	// Backend states which backend to use for Audits.
	Backend string `yaml:"backend" conf:",example=local"`
	// IncludeMetadataKeys defines which gRPC request metadata keys should be included in the audit logs.
	IncludeMetadataKeys []string `yaml:"includeMetadataKeys" conf:",example=['content-type']"`
	// ExcludeMetadataKeys defines which gRPC request metadata keys should be excluded from the audit logs. Takes precedence over includeMetadataKeys.
	ExcludeMetadataKeys []string `yaml:"excludeMetadataKeys" conf:",example=['authorization']"`
	// Enabled defines whether audit logging is enabled.
	Enabled bool `yaml:"enabled" conf:",example=false"`
	// AccessLogsEnabled defines whether access logging is enabled.
	AccessLogsEnabled bool `yaml:"accessLogsEnabled" conf:",example=false"`
	// DecisionLogsEnabled defines whether logging of policy decisions is enabled.
	DecisionLogsEnabled bool `yaml:"decisionLogsEnabled" conf:",example=false"`
	// DecisionLogFilters define the filters to apply while producing decision logs.
	DecisionLogFilters DecisionLogFilters `yaml:"decisionLogFilters"`
}

type DecisionLogFilters struct {
	// CheckResources defines the filters that apply to CheckResources calls.
	CheckResources CheckResourcesFilter `yaml:"checkResources"`
	// PlanResources defines the filters that apply to PlanResources calls.
	PlanResources PlanResourcesFilter `yaml:"planResources"`
}

type CheckResourcesFilter struct {
	// IgnoreAllowAll ignores responses that don't contain an EFFECT_DENY.
	IgnoreAllowAll bool `yaml:"ignoreAllowAll" conf:",example=false"`
}

type PlanResourcesFilter struct {
	// IgnoreAll prevents any plan responses from being logged. Takes precedence over other filters.
	IgnoreAll bool `yaml:"ignoreAll" conf:",example=false"`
	// IgnoreAlwaysAllow ignores ALWAYS_ALLOWED plans.
	IgnoreAlwaysAllow bool `yaml:"ignoreAlwaysAllow" conf:",example=false"`
}

func (c *Conf) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	// This is a workaround to circumvent strict config parsing.
	// Consider the following:
	//
	// audit:
	//   enabled: true
	//   backend: local
	//   local:
	//     storageDirectory: /var/cerbos
	//
	// Because the config for the "local" backend is nested under the "audit" key, we will
	// need to have "local" defined as a field in the Conf struct when strict parsing is on.
	// However, backends are self-contained plugins and it is not practical to add a new field
	// definition to the audit config struct for each plugin we introduce.
	// This hack is slightly inefficient because it marshals and unmarshals YAML twice. It is an
	// acceptable sacrifice to make because config is only read once on startup.
	return nil
}

func (c *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (c *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

func GetConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }
