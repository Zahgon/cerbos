// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package hub

import (
	"context"
	"io"

	"github.com/cerbos/cloud-api/credentials"
	bundlev1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v1"
	bundlev2 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v2"
	"github.com/spf13/afero"
	"go.uber.org/zap"

	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/storage"
)

const (
	manifestFileName = "MANIFEST"
	policyDir        = "policies/"
	schemaDir        = "_schemas/"
)

type cleanupFn func() error

type OpenOpts struct {
	Credentials   *credentials.Credentials
	ScratchFS     afero.Fs
	BundlePath    string
	Source        string
	EncryptionKey []byte
	CacheSize     uint
}

type LegacyBundle struct {
	bundleFS afero.Fs
	manifest *bundlev2.Manifest
	cleanup  cleanupFn
	path     string
}

func toManifestV2(manifest *bundlev1.Manifest) *bundlev2.Manifest {
	_ = "STUB: not implemented"
	return nil
}

func OpenLegacy(opts OpenOpts) (*LegacyBundle, error) { _ = "STUB: not implemented"; return nil, nil }

func decryptLegacyBundle(opts OpenOpts, logger *zap.Logger) (string, int64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func OpenLegacyV2(opts OpenOpts) (*LegacyBundle, error) { _ = "STUB: not implemented"; return nil, nil }

func decryptLegacyBundleV2(opts OpenOpts, logger *zap.Logger) (string, int64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func archiveToFS(opts OpenOpts, archivePath string, archiveSize int64, logger *zap.Logger) (afero.Fs, cleanupFn, error) {
	_ = "STUB: not implemented"
	return *new(afero.Fs), *new(cleanupFn), nil
}

// Because we use random strings to avoid a clash, clean up the file

func loadManifest(bundleFS afero.Fs) (*bundlev1.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadManifestV2(bundleFS afero.Fs) (*bundlev2.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readManifestFile(bundleFS afero.Fs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) ID() string { _ = "STUB: not implemented"; return "" }

func (lb *LegacyBundle) Type() bundlev2.BundleType {
	_ = "STUB: not implemented"
	return *new(bundlev2.BundleType)
}

func (lb *LegacyBundle) GetFirstMatch(_ context.Context, candidates []namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) GetAll(_ context.Context) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAllMatching attempts to retrieve all policies from the passed modIDs, unlike `GetFirstMatch` which returns the first
// of the passed candidates, this function returns list of all available modules from the provided IDs.
func (lb *LegacyBundle) GetAllMatching(_ context.Context, modIDs []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) getMatch(id namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) loadPolicySet(idHex, fileName string) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) ListPolicyIDs(_ context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) ListSchemaIDs(_ context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *LegacyBundle) LoadSchema(_ context.Context, path string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// TODO(cell): Should we write the schema to scratch dir and create a reader for that instead?

func (lb *LegacyBundle) Release() error { _ = "STUB: not implemented"; return nil }

func (lb *LegacyBundle) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (lb *LegacyBundle) Close() error { _ = "STUB: not implemented"; return nil }
