// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package schema

import (
	"context"
	"io"
	"net/url"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"

	// Register the http and https loaders.
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"

	"github.com/cerbos/cerbos/internal/cache"
	"github.com/cerbos/cerbos/internal/storage"
)

const fileURLScheme = "file"

func NewFromConf(_ context.Context, loader Loader, conf *Conf) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

type manager struct {
	StaticManager
	cache    *cache.Cache[string, *cacheEntry]
	resolver Resolver
}

func New(ctx context.Context, loader Loader) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

func NewEphemeral(resolver Resolver) Manager { _ = "STUB: not implemented"; return *new(Manager) }

func (m *manager) LoadSchema(ctx context.Context, schemaURL string) (*jsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *manager) loadSchemaFromStore(ctx context.Context, schemaURL string) (*jsonschema.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *manager) SubscriberID() string { _ = "STUB: not implemented"; return "" }

func (m *manager) OnStorageEvent(events ...storage.Event) { _ = "STUB: not implemented"; return }

//nolint:exhaustive

type cacheEntry struct {
	schema *jsonschema.Schema
	err    error
}

func DefaultResolver(loader Loader) Resolver { _ = "STUB: not implemented"; return *new(Resolver) }

func loadFileURL(u *url.URL) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

//nolint:gosec

func loadHTTPURL(ctx context.Context, u *url.URL) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

//nolint:gosec
