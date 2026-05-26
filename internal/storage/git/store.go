// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package git

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/storage/index"
	"github.com/cerbos/cerbos/internal/util"
)

const DriverName = "git"

var driverAttr = policy.SourceDriver(DriverName)

var (
	_ storage.SourceStore  = (*Store)(nil)
	_ storage.Reloadable   = (*Store)(nil)
	_ storage.Subscribable = (*Store)(nil)
)

func init() {
	storage.RegisterDriver(DriverName, func(ctx context.Context, confW *config.Wrapper) (storage.Store, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read git configuration: %w", err)
		}

		return NewStore(ctx, conf)
	})
}

type Store struct {
	log  *zap.SugaredLogger
	conf *Conf
	idx  index.Index
	repo *git.Repository
	sf   singleflight.Group
	*storage.SubscriptionManager
	subDir         string
	currCommitHash string
	mu             sync.RWMutex
}

func NewStore(ctx context.Context, conf *Conf) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) init(ctx context.Context) error {
	if s.conf.ScratchDir != "" {
		s.log.Warnf("ScratchDir storage option is deprecated and will be removed in a future release")
	}

	finfo, err := os.Stat(s.conf.CheckoutDir)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to stat %s: %w", s.conf.CheckoutDir, err)
	} else if finfo != nil && !finfo.IsDir() {
		return fmt.Errorf("not a directory: %s", s.conf.CheckoutDir)
	}

	loadAndStartPoller := func() error {
		if err := s.loadAll(ctx); err != nil {
			return err
		}

		go s.pollForUpdates(ctx)

		return nil
	}

	// if the directory does not exist, create it and clone the repo
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(s.conf.CheckoutDir, 0o744); err != nil { //nolint:mnd
			return fmt.Errorf("failed to create directory %s: %w", s.conf.CheckoutDir, err)
		}

		if err := s.cloneRepo(ctx); err != nil {
			return err
		}

		return loadAndStartPoller()
	}

	// check whether the directory is empty
	empty, err := isEmptyDir(s.conf.CheckoutDir)
	if err != nil {
		return err
	}

	// if empty, clone the repo
	if empty {
		if err := s.cloneRepo(ctx); err != nil {
			return err
		}

		return loadAndStartPoller()
	}

	// if not empty, assume it is a git repo and try to pull the latest changes
	if _, err := s.pullAndCompare(ctx); err != nil {
		return err
	}

	s.setCurrentCommitHash()

	return loadAndStartPoller()
}

func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }

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

func (s *Store) currentCommitHash() string { _ = "STUB: not implemented"; return "" }

func (s *Store) setCurrentCommitHash() { _ = "STUB: not implemented"; return }

func isEmptyDir(dir string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *Store) cloneRepo(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Store) loadAll(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Store) pullAndCompare(ctx context.Context) (object.Changes, error) {
	_ = "STUB: not implemented"
	return *new(object.Changes), nil
}

// open the repo if it's not already open.

// Make sure we are in the correct branch and get the current HEAD

// Now pull from remote

// branch is already up-to-date: nothing to do.

// compare the head with the prev state.

func (s *Store) openRepo() error { _ = "STUB: not implemented"; return nil }

func (s *Store) ensureCorrectBranch() (plumbing.Hash, error) {
	_ = "STUB: not implemented"
	return *new(plumbing.Hash), nil
}

func (s *Store) compareWithHEAD(ctx context.Context, prevHash plumbing.Hash) (object.Changes, error) {
	_ = "STUB: not implemented"
	return *new(object.Changes), nil
}

func (s *Store) getTreeForHash(hash plumbing.Hash) (*object.Tree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Store) updateIndex(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Store) normalizePath(path string) (string, util.IndexedFileType) {
	_ = "STUB: not implemented"
	return "", *new(util.IndexedFileType)
}

// not in policies directory

func (s *Store) applyIndexUpdate(ce object.ChangeEntry, eventKind storage.EventKind, headHash plumbing.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Store) readPolicyFromBlob(hash, headHash plumbing.Hash) (*policyv1.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func commitHashAttr(hash plumbing.Hash) policy.SourceAttribute {
	_ = "STUB: not implemented"
	return *new(policy.SourceAttribute)
}

func (s *Store) pollForUpdates(ctx context.Context) { _ = "STUB: not implemented"; return }
