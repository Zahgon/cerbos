// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package hub

import (
	"errors"

	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/hub"
	"github.com/cerbos/cerbos/internal/storage"
)

const (
	confKey          = storage.ConfKey + "." + DriverName
	defaultCacheSize = 1024
)

var ErrNoSource = errors.New("at least one of local or remote sources must be defined")

// Conf is required (if driver is set to 'hub') configuration for hub storage driver.
// +desc=This section is required only if storage.driver is hub.
type Conf struct {
	// Remote holds configuration for remote bundle source. Takes precedence over local if both are defined.
	Remote *RemoteSourceConf `yaml:"remote"`
	// Local holds configuration for local bundle source.
	Local *LocalSourceConf `yaml:"local"`
	// Credentials holds Cerbos Hub credentials.
	Credentials *hub.CredentialsConf `yaml:"credentials" conf:",ignore"`
	// CacheSize defines the number of policies to cache in memory.
	CacheSize uint `yaml:"cacheSize" conf:",example=1024"`
}

// LocalSourceConf holds configuration for local bundle store.
type LocalSourceConf struct {
	// BundlePath is the full path to the local bundle file.
	BundlePath string `yaml:"bundlePath" conf:"required,example=/path/to/bundle.crbp"`
	// EncryptionKey is encryption key to decode the bundle. It must be string encoded.
	EncryptionKey string `yaml:"encryptionKey" conf:",ignore"`
	// TempDir is the directory to use for temporary files.
	TempDir string `yaml:"tempDir" conf:",example=${TEMP}"`
}

// RemoteSourceConf holds configuration for remote bundle store.
type RemoteSourceConf struct {
	// Connection defines settings for the remote server connection.
	Connection *hub.ConnectionConf `yaml:"connection" conf:",ignore"`
	// BundleLabel to fetch from the server.
	BundleLabel string `yaml:"bundleLabel" conf:"required,example=latest"`
	// DeploymentID to fetch from the server.
	DeploymentID string `yaml:"deploymentID" conf:",ignore"`
	// PlaygroundID to fetch from the server.
	PlaygroundID string `yaml:"playgroundID" conf:",ignore"`
	// CacheDir is the directory to use for caching downloaded bundles.
	CacheDir string `yaml:"cacheDir" conf:",example=${XDG_CACHE_DIR}"`
	// TempDir is the directory to use for temporary files.
	TempDir string `yaml:"tempDir" conf:",example=${TEMP}"`
	// DisableAutoUpdate sets whether new bundles should be automatically downloaded and applied.
	DisableAutoUpdate bool `yaml:"disableAutoUpdate" conf:",example=false"`
	// DisableBootstrap makes the PDP always fetch bundles using the API. If the API is down, the PDP won't be able to start.
	DisableBootstrap bool `yaml:"disableBootstrap" conf:",ignore"`
}

func (conf *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (conf *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

func (conf *Conf) Validate() (outErr error) { _ = "STUB: not implemented"; return nil }

func (conf *Conf) validateCredentials() error { _ = "STUB: not implemented"; return nil }

func (lc *LocalSourceConf) validate() error { _ = "STUB: not implemented"; return nil }

func (lc *LocalSourceConf) setDefaultsForUnsetFields() error { _ = "STUB: not implemented"; return nil }

func (rc *RemoteSourceConf) validate() error { _ = "STUB: not implemented"; return nil }

func (rc *RemoteSourceConf) setDefaultsForUnsetFields() error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

func GetConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }

func GetConfFromWrapper(confW *config.Wrapper) (*Conf, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
