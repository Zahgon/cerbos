// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	"time"

	"github.com/cerbos/cloud-api/credentials"
)

type EnvVarKey int

const (
	BundleLabelKey EnvVarKey = iota
	ClientIDKey
	ClientSecretKey
	OfflineKey
	PDPIDKey
	WorkspaceSecretKey
	BundleVersionKey
	DeploymentIDKey
	PlaygroundIDKey
)

var envVars = map[EnvVarKey][]string{
	BundleLabelKey:     {"CERBOS_HUB_BUNDLE", "CERBOS_CLOUD_BUNDLE"},
	ClientIDKey:        {"CERBOS_HUB_CLIENT_ID", "CERBOS_CLOUD_CLIENT_ID"},
	ClientSecretKey:    {"CERBOS_HUB_CLIENT_SECRET", "CERBOS_CLOUD_CLIENT_SECRET"},
	OfflineKey:         {"CERBOS_HUB_OFFLINE", "CERBOS_CLOUD_OFFLINE"},
	PDPIDKey:           {"CERBOS_HUB_PDP_ID", "CERBOS_PDP_ID"},
	WorkspaceSecretKey: {"CERBOS_HUB_WORKSPACE_SECRET", "CERBOS_CLOUD_SECRET_KEY"},
	BundleVersionKey:   {"CERBOS_HUB_BUNDLE_VERSION"},
	DeploymentIDKey:    {"CERBOS_HUB_DEPLOYMENT_ID"},
	PlaygroundIDKey:    {"CERBOS_HUB_PLAYGROUND_ID"},
}

func GetEnv(key EnvVarKey) string { _ = "STUB: not implemented"; return "" }

const (
	confKey                  = "hub"
	defaultAPIEndpoint       = "https://api.cerbos.cloud"
	defaultBootstrapHost     = "https://cdn.cerbos.cloud"
	defaultHeartbeatInterval = 180 * time.Second
	defaultMaxRetryWait      = 120 * time.Second
	defaultMinRetryWait      = 1 * time.Second
	defaultNumRetries        = 5
	minHeartbeatInterval     = 30 * time.Second
)

type Conf struct {
	// Credentials holds Cerbos Hub client credentials.
	Credentials CredentialsConf `yaml:"credentials"`
	// Connection holds advanced connection settings for Cerbos Hub.
	Connection ConnectionConf `yaml:"connection" conf:",ignore"`
}

func (conf *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (conf *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

func (conf *Conf) Validate() (outErr error) { _ = "STUB: not implemented"; return nil }

// CredentialsConf holds credentials for accessing Cerbos Hub.
type CredentialsConf struct {
	// PDPID is the unique identifier for this Cerbos instance. Defaults to the value of the CERBOS_HUB_PDP_ID environment variable.
	PDPID string `yaml:"pdpID" conf:",example=crb-004"`
	// ClientID of the Cerbos Hub credential. Defaults to the value of the CERBOS_HUB_CLIENT_ID environment variable.
	ClientID string `yaml:"clientID" conf:",example=92B0K05B6HOF"`
	// ClientSecret of the Cerbos Hub credential. Defaults to the value of the CERBOS_HUB_CLIENT_SECRET environment variable.
	ClientSecret string `yaml:"clientSecret" conf:",example=${CERBOS_HUB_CLIENT_SECRET}"` //nolint:gosec
	// WorkspaceSecret used to decrypt the bundles. Defaults to the value of the CERBOS_HUB_WORKSPACE_SECRET environment variable.
	WorkspaceSecret string `yaml:"workspaceSecret" conf:",example=${CERBOS_HUB_WORKSPACE_SECRET}"`
	// Deprecated: Use PDPID
	InstanceID string `yaml:"instanceID" conf:",ignore"`
	// Deprecated: Use WorkspaceSecret
	SecretKey string `yaml:"secretKey" conf:",ignore"`
}

func (cc *CredentialsConf) Validate() (outErr error) {
	_ = "STUB: not implemented"
	// SecretKey was renamed to WorkspaceSecret in Cerbos 0.31.0
	return nil
}

// InstanceID was renamed to PDPID in Cerbos 0.31.0

// We don't do any validation here because some fields are optional depending on the use case.

func (cc *CredentialsConf) LoadFromEnv() { _ = "STUB: not implemented"; return }

func (cc CredentialsConf) ToCredentials() (*credentials.Credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConnectionConf holds configuration for the remote connection.
type ConnectionConf struct {
	// TLS defines settings for TLS connections.
	TLS TLSConf `yaml:"tls"`
	// APIEndpoint is the address of the API server.
	APIEndpoint string `yaml:"apiEndpoint" conf:"required,example=https://api.cerbos.cloud"`
	// BootstrapEndpoint is the addresses of the server serving the bootstrap configuration.
	BootstrapEndpoint string `yaml:"bootstrapEndpoint" conf:"required,example=https://cdn.cerbos.cloud"`
	// MinRetryWait is the minimum amount of time to wait between retries.
	MinRetryWait time.Duration `yaml:"minRetryWait" conf:",example=1s"`
	// MaxRetryWait is the maximum amount of time to wait between retries.
	MaxRetryWait time.Duration `yaml:"maxRetryWait" conf:",example=120s"`
	// NumRetries is the number of times to retry before giving up.
	NumRetries uint `yaml:"numRetries" conf:",example=5"`
	// HeartbeatInterval is the interval for sending regular heartbeats.
	HeartbeatInterval time.Duration `yaml:"heartbeatInterval" conf:",example=2m"`
}

func (cc ConnectionConf) IsUnset() bool { _ = "STUB: not implemented"; return false }

func (cc *ConnectionConf) Validate() error { _ = "STUB: not implemented"; return nil }

// TLSConf holds TLS configuration for the remote connection.
type TLSConf struct {
	// Authority overrides the Cerbos Hub server authority if it is different from what is provided in the API and bootstrap endpoints.
	Authority string `yaml:"authority" conf:",example=domain.tld"`
	// CACert is the path to the CA certificate chain to use for certificate verification.
	CACert string `yaml:"caCert" conf:",example=/path/to/CA_certificate"`
}

func GetConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }
