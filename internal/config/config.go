// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	_ "embed"
	"errors"
	"io"
	"sync"

	"go.uber.org/config"
)

const DefaultMarker = "__default__"

//go:embed .cerbos.yaml.gotmpl
var defaultConfTmpl string

var ErrConfigNotLoaded = errors.New("config not loaded")

var conf = &Wrapper{}

type Section interface {
	Key() string
}

type Defaulter interface {
	SetDefaults()
}

type Validator interface {
	Validate() error
}

// Load loads the config file at the given path.
func Load(confFile string, overrides map[string]any) error { _ = "STUB: not implemented"; return nil }

func loadDefault(overrides map[string]any) error { _ = "STUB: not implemented"; return nil }

func LoadReader(reader io.Reader, overrides map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func LoadMap(m map[string]any) error { _ = "STUB: not implemented"; return nil }

func doLoad(sources ...config.YAMLOption) error { _ = "STUB: not implemented"; return nil }

func mkProvider(sources ...config.YAMLOption) (config.Provider, error) {
	_ = "STUB: not implemented"
	return *new(config.Provider), nil
}

//nolint:gocritic

// Global returns the default global config wrapper.
func Global() *Wrapper {
	_ = "STUB: not implemented"

	// Get populates out with the configuration at the given key.
	// Populate out with default values before calling this function to ensure sane defaults if there are any.
	return nil
}

func Get(key string, out any) error { _ = "STUB: not implemented"; return nil }

// GetSection populates a config section.
func GetSection(section Section) error { _ = "STUB: not implemented"; return nil }

func WrapperFromReader(reader io.Reader, overrides map[string]any) (*Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func WrapperFromMap(m map[string]any) (*Wrapper, error) { _ = "STUB: not implemented"; return nil, nil }

func newWrapper(sources ...config.YAMLOption) (*Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Wrapper struct {
	provider config.Provider
	mu       sync.RWMutex
}

func (w *Wrapper) Get(key string, out any) error { _ = "STUB: not implemented"; return nil }

// set defaults if any are specified

// validate if a validate function is available

func (w *Wrapper) GetSection(section Section) error { _ = "STUB: not implemented"; return nil }

func (w *Wrapper) replaceProvider(provider config.Provider) { _ = "STUB: not implemented"; return }
