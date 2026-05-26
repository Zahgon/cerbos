// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package storage

const ConfKey = "storage"

// Conf is required configuration for storage.
// +desc=This section is required. The field driver must be set to indicate which driver to use.
type Conf struct {
	confHolder
}

// confHolder exists to avoid a recursive loop in the UnmarshalYAML method below.
type confHolder struct {
	// Driver defines which storage driver to use.
	Driver string `yaml:"driver" conf:"required,example=\"disk\""`
}

func (c *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (c *Conf) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	// We want to avoid defining all the storage driver configuration structs as fields of the Conf
	// struct to maintain the "plugin" nature of those drivers (and avoid circular package references).
	// However, the strict YAML parser throws an error if it sees undefined fields. This is a slightly
	// inefficient workaround to get over that issue.
	return nil
}

func GetConf() (*Conf, error) { _ = "STUB: not implemented"; return nil, nil }
