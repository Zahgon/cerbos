// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package hub

import (
	"context"
	"errors"
	"io"
	"regexp"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v5"
	bundleapi "github.com/cerbos/cloud-api/bundle"
	bundleapiv2 "github.com/cerbos/cloud-api/bundle/v2"
	"github.com/cerbos/cloud-api/credentials"
	bundlev2 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/bundle/v2"
	hubapi "github.com/cerbos/cloud-api/hub"
	"github.com/spf13/afero"
	"go.uber.org/zap"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/ruletable"
	"github.com/cerbos/cerbos/internal/storage"
)

const (
	defaultReconnectBackoff = 5 * time.Second
	noBundleInitialInterval = 60 * time.Second
	noBundleMaxInterval     = 10 * time.Minute
	noBundleMaxCount        = 10
)

var (
	_ storage.BinaryStore  = (*RemoteSource)(nil)
	_ storage.Reloadable   = (*RemoteSource)(nil)
	_ storage.Instrumented = (*RemoteSource)(nil)

	playgroundLabelPattern = regexp.MustCompile(`^playground/[A-Z0-9]{12}$`)

	ErrOfflineModeNotAvailable = errors.New("offline mode is not available when bundle version is set to 2")
)

type Bundle interface {
	io.Closer
	ID() string
	Release() error
	Type() bundlev2.BundleType
	InspectPolicies(context.Context, storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error)
	ListPolicyIDs(context.Context, storage.ListPolicyIDsParams) ([]string, error)
	ListSchemaIDs(context.Context) ([]string, error)
	LoadSchema(context.Context, string) (io.ReadCloser, error)
	GetFirstMatch(context.Context, []namer.ModuleID) (*runtimev1.RunnablePolicySet, error)
	GetAll(context.Context) ([]*runtimev1.RunnablePolicySet, error)
	GetAllMatching(context.Context, []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error)
	RepoStats(context.Context) storage.RepoStats
}

type cloudAPIClient interface {
	BootstrapBundle(context.Context) (string, bundlev2.BundleType, []byte, error)
	GetBundle(context.Context) (string, bundlev2.BundleType, []byte, error)
	GetCachedBundle() (string, error)
	OpenCredentials() *credentials.Credentials
	WatchBundle(context.Context) (bundleapi.WatchHandle, error)
}

type cloudAPIv1 struct {
	client      ClientV1
	bundleLabel string
	playground  bool
}

func (apiv1 *cloudAPIv1) BootstrapBundle(ctx context.Context) (string, bundlev2.BundleType, []byte, error) {
	_ = "STUB: not implemented"
	return "", *new(bundlev2.BundleType), nil, nil
}

func (apiv1 *cloudAPIv1) GetBundle(ctx context.Context) (string, bundlev2.BundleType, []byte, error) {
	_ = "STUB: not implemented"
	return "", *new(bundlev2.BundleType), nil, nil
}

func (apiv1 *cloudAPIv1) GetCachedBundle() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (apiv1 *cloudAPIv1) OpenCredentials() *credentials.Credentials {
	_ = "STUB: not implemented"
	return nil
}

func (apiv1 *cloudAPIv1) WatchBundle(ctx context.Context) (bundleapi.WatchHandle, error) {
	_ = "STUB: not implemented"
	return *new(bundleapi.WatchHandle), nil
}

type cloudAPIv2 struct {
	client ClientV2
	source bundleapiv2.Source
}

func (apiv2 *cloudAPIv2) BootstrapBundle(ctx context.Context) (string, bundlev2.BundleType, []byte, error) {
	_ = "STUB: not implemented"
	return "", *new(bundlev2.BundleType), nil, nil
}

func (apiv2 *cloudAPIv2) GetBundle(ctx context.Context) (string, bundlev2.BundleType, []byte, error) {
	_ = "STUB: not implemented"
	return "", *new(bundlev2.BundleType), nil, nil
}

func (apiv2 *cloudAPIv2) GetCachedBundle() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (apiv2 *cloudAPIv2) OpenCredentials() *credentials.Credentials {
	_ = "STUB: not implemented"
	return nil
}

func (apiv2 *cloudAPIv2) WatchBundle(ctx context.Context) (bundleapi.WatchHandle, error) {
	_ = "STUB: not implemented"
	return *new(bundleapi.WatchHandle), nil
}

// RemoteSource implements a bundle store that loads bundles from a remote source.
type RemoteSource struct {
	hub       ClientProvider
	scratchFS afero.Fs
	client    cloudAPIClient
	log       *zap.Logger
	conf      *Conf
	bundle    Bundle
	*storage.SubscriptionManager
	bundleVersion bundleapi.Version
	mu            sync.RWMutex
	healthy       bool
}

func NewRemoteSource(conf *Conf) (*RemoteSource, error) { _ = "STUB: not implemented"; return nil, nil }

type ClientProvider interface {
	V1(bundleapi.ClientConf) (ClientV1, error)
	V2(bundleapi.ClientConf) (ClientV2, error)
}

type hubClientProvider struct {
	*hubapi.Hub
}

func (h hubClientProvider) V1(conf bundleapi.ClientConf) (ClientV1, error) {
	_ = "STUB: not implemented"
	return *new(ClientV1), nil
}

func (h hubClientProvider) V2(conf bundleapi.ClientConf) (ClientV2, error) {
	_ = "STUB: not implemented"
	return *new(ClientV2), nil
}

type ClientV1 interface {
	HubCredentials() *credentials.Credentials
	BootstrapBundle(context.Context, string) (string, error)
	GetBundle(context.Context, string) (string, error)
	GetCachedBundle(string) (string, error)
	WatchBundle(context.Context, string) (bundleapi.WatchHandle, error)
}

type ClientV2 interface {
	BootstrapBundle(context.Context, bundleapiv2.Source) (string, bundlev2.BundleType, []byte, error)
	GetBundle(context.Context, bundleapiv2.Source) (string, bundlev2.BundleType, []byte, error)
	GetCachedBundle(bundleapiv2.Source) (string, error)
	WatchBundle(context.Context, bundleapiv2.Source) (bundleapi.WatchHandle, error)
}

func NewRemoteSourceWithHub(conf *Conf, hub ClientProvider) (*RemoteSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) Init(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Ideally we want to be able to automatically switch between online and offline modes.
// That logic is complicated to implement and test in the little time we have. There are open questions
// about expected behaviour as well. For example, is it preferable to use a stale copy from cache or fail fast?
// So, this initial version just provides an escape hatch to manually deal with downtime by putting the PDP
// into offline mode.
// TODO(cell): Implement automatic online/offline mode
// TODO(oguzhan): Get rid of offline mode when we no longer support bundle.Version1.

// fail fast if the service is down

type noBundleBackoff struct {
	backoff backoff.BackOff
	count   uint
}

func (b *noBundleBackoff) NextBackOff() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (b *noBundleBackoff) Reset() { _ = "STUB: not implemented"; return }

func shouldWorkOffline() bool { _ = "STUB: not implemented"; return false }

func (s *RemoteSource) fetchBundle(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RemoteSource) fetchBundleOffline() error {
	_ = "STUB: not implemented"
	// TODO(oguzhan): Get rid of offline mode when we no longer support bundle.Version1.
	return nil
}

func (s *RemoteSource) removeBundle(healthy bool) { _ = "STUB: not implemented"; return }

func (s *RemoteSource) swapBundle(bundlePath string, encryptionKey []byte, bundleType bundlev2.BundleType) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RemoteSource) activeBundleID() string { _ = "STUB: not implemented"; return "" }

func (s *RemoteSource) startWatchLoop(ctx context.Context, noBundleBackoff backoff.BackOff) {
	_ = "STUB: not implemented"
	return
}

// reset backoff if the last call succeeded

func incEventMetric(event string) { _ = "STUB: not implemented"; return }

func (s *RemoteSource) startWatch(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// retry indefinitely

// Returning a nil error causes the connection to be re-established.
// Returning a non-nil error terminates the process.

func (s *RemoteSource) Driver() string { _ = "STUB: not implemented"; return "" }

func (s *RemoteSource) GetRuleTable() (*ruletable.RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) IsHealthy() bool { _ = "STUB: not implemented"; return false }

func (s *RemoteSource) GetFirstMatch(ctx context.Context, candidates []namer.ModuleID) (*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) GetAll(ctx context.Context) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) GetAllMatching(ctx context.Context, modIDs []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) ListSchemaIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RemoteSource) LoadSchema(ctx context.Context, id string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (s *RemoteSource) Reload(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *RemoteSource) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (s *RemoteSource) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

func (s *RemoteSource) SourceKind() string { _ = "STUB: not implemented"; return "" }

func (s *RemoteSource) Close() error { _ = "STUB: not implemented"; return nil }
