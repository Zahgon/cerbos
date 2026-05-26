// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package blob

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"gocloud.dev/blob"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/storage/index"
)

const (
	dotcache   = ".cache"
	DriverName = "blob"
)

var (
	_ storage.SourceStore  = (*Store)(nil)
	_ storage.Reloadable   = (*Store)(nil)
	_ storage.Subscribable = (*Store)(nil)
)

var ErrUnsupportedBucketScheme = errors.New("currently only \"s3\" and \"gs\" bucket URL schemes are supported")

var driverSourceAttr = policy.SourceDriver(DriverName)

func init() {
	storage.RegisterDriver(DriverName, func(ctx context.Context, confW *config.Wrapper) (storage.Store, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read blob configuration: %w", err)
		}

		bucket, url, err := newBucket(ctx, conf)
		if err != nil {
			return nil, err
		}

		source := &auditv1.PolicySource{
			Source: &auditv1.PolicySource_Blob_{
				Blob: &auditv1.PolicySource_Blob{
					BucketUrl: url,
					Prefix:    conf.Prefix,
				},
			},
		}

		cacheDir := cacheDir(conf.Bucket, conf.WorkDir)
		workDir := conf.WorkDir

		if err := createOrValidateDir(workDir); err != nil {
			return nil, fmt.Errorf("failed to create work directory: %w", err)
		}

		cloner, err := NewCloner(bucket, cacheDir)
		if err != nil {
			return nil, fmt.Errorf("failed to create cloner: %w", err)
		}

		workFS := newBlobFS(workDir)
		return NewStore(ctx, conf, workFS, cloner, symlinkerFunc(func(destination, source string) error {
			src := filepath.Join(workDir, source)
			dst := filepath.Join(cacheDir, destination)

			return os.Symlink(dst, src)
		}), source)
	})
}

func newBucket(ctx context.Context, conf *Conf) (*blob.Bucket, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func openGSBucket(ctx context.Context, conf *Conf, bucketURL *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The following query parameters are supported:
//
//   - access_id: sets Options.GoogleAccessID
//   - private_key_path: path to read for Options.PrivateKey
//
// Currently their use is limited to SignedURL.

func openS3Bucket(ctx context.Context, conf *Conf, bucketURL *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bucketCloner interface {
	Clean() error
	Clone(ctx context.Context) (*CloneResult, error)
}

type symlinker interface {
	Symlink(destination, source string) error
}

type symlinkerFunc func(destination, source string) error

func (s symlinkerFunc) Symlink(destination, source string) error {
	_ = "STUB: not implemented"
	return nil
}

type Store struct {
	*storage.SubscriptionManager
	log         *zap.SugaredLogger
	conf        *Conf
	source      *auditv1.PolicySource
	idx         index.Index
	cloner      bucketCloner
	symlink     symlinker
	workFS      FS
	currDirName string
	workDir     string
}

func (s *Store) Subscribe(sub storage.Subscriber) { _ = "STUB: not implemented"; return }

func NewStore(ctx context.Context, conf *Conf, workFS FS, cloner bucketCloner, symlink symlinker, source *auditv1.PolicySource) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) init(ctx context.Context) error {
	cr, err := s.clone(ctx)
	if err != nil {
		s.log.Errorw("Failed to clone blob store", "error", err)
		return fmt.Errorf("failed to clone blob store: %w", err)
	}

	idx, dirName, err := s.buildIndex(ctx, cr.all)
	if err != nil {
		return fmt.Errorf("failed to build index from the new set of files: %w", err)
	}

	s.currDirName = dirName
	s.idx = idx
	go s.pollForUpdates(ctx)

	if err := s.cloner.Clean(); err != nil {
		s.log.Warnw("Failed to clean up the cache", "error", err)
	}

	return nil
}

func (s *Store) updateIndex(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

// to prevent duplicating shared dependents across multiple events, we create a single stub event
// which carries all dependents not already included in this batch of events

// we need to emit all events regardless of validity as some subscribers (such as the rule table)
// need to be kept in sync.

func (s *Store) addOrUpdateEvent(etag, file, currDirName string, ts int64) (storage.Event, error) {
	_ = "STUB: not implemented"
	return *new(storage.Event), nil
}

func (s *Store) deleteEvent(file string) (storage.Event, error) {
	_ = "STUB: not implemented"
	return *new(storage.Event), nil
}

func (s *Store) clone(ctx context.Context) (*CloneResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) pollForUpdates(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Store) createSymlink(tmpDir, destination, source string) error {
	_ = "STUB: not implemented"
	return nil
}

// If there are subdirectories in the blob storage we need to create them in the source directory before creation
// of the symlink

func (s *Store) createSymlinks(all map[string][]string, tmpDir string) error {
	_ = "STUB: not implemented"
	return nil
}

type indexBuildError struct {
	err error
	// dir is the temporary work directory which the store tried build an index from the new set of files
	dir string
}

func (e *indexBuildError) Error() string { _ = "STUB: not implemented"; return "" }

// buildIndex creates a new work directory with its name set to current timestamp, creates symlinks targeted to
// s.cacheDir according to the given map 'all' and tries to build a temporary index to see if there are any errors
// with the incoming policies/schemas. If there are no errors returns the index built and the path to the new work directory.
func (s *Store) buildIndex(ctx context.Context, all map[string][]string) (index.Index, string, error) {
	_ = "STUB: not implemented"
	return *new(index.Index), "", nil
}

func (s *Store) prepareWorkDir(ctx context.Context, all map[string][]string) (dir, currDirName string, ts int64, err error) {
	_ = "STUB: not implemented"
	return "", "", 0, nil
}

func (s *Store) buildIndexFromWorkDir(ctx context.Context, dir, currDirName string, ts int64) (idx index.Index, err error) {
	_ = "STUB: not implemented"
	return *new(index.Index), nil
}

func (s *Store) Driver() string { _ = "STUB: not implemented"; return "" }

func (s *Store) GetFirstMatch(_ context.Context, candidates []namer.ModuleID) (*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetAll(ctx context.Context) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetAllMatching(_ context.Context, modIDs []namer.ModuleID) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetCompilationUnits(_ context.Context, ids ...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) GetDependents(_ context.Context, ids ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) InspectPolicies(ctx context.Context, params storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) ListPolicyIDs(ctx context.Context, params storage.ListPolicyIDsParams) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) ListSchemaIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) LoadSchema(ctx context.Context, url string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (s *Store) LoadPolicy(ctx context.Context, file ...string) ([]*policy.Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) RepoStats(ctx context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (s *Store) Reload(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }

func cacheDir(bucketURL, workDir string) string { _ = "STUB: not implemented"; return "" }

func indexBuildTSSourceAttr(ts int64) policy.SourceAttribute {
	_ = "STUB: not implemented"
	return *new(policy.SourceAttribute)
}

func etagSourceAttr(etag string) policy.SourceAttribute {
	_ = "STUB: not implemented"
	return *new(policy.SourceAttribute)
}
