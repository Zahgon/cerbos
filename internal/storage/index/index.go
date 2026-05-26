// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package index

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"sync"

	"golang.org/x/sync/singleflight"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/parser"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
)

var (
	// ErrDuplicatePolicy signals that there are duplicate policy definitions.
	ErrDuplicatePolicy = errors.New("duplicate policy definitions")
	// ErrInvalidEntry signals that the index entry is invalid.
	ErrInvalidEntry = errors.New("invalid index entry")
	// ErrPolicyNotFound signals that the policy does not exist.
	ErrPolicyNotFound = errors.New("policy not found")
)

type ModuleIDSet map[namer.ModuleID]struct{}

type Entry struct {
	File   string
	Policy policy.Wrapper
}

type Index interface {
	io.Closer
	storage.Instrumented
	GetFirstMatch([]namer.ModuleID) (*policy.CompilationUnit, error)
	GetAll(context.Context) ([]*policy.CompilationUnit, error)
	GetAllMatching([]namer.ModuleID) ([]*policy.CompilationUnit, error)
	GetCompilationUnits(...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error)
	GetDependents(...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error)
	AddOrUpdate(Entry) (storage.Event, error)
	Delete(Entry) (storage.Event, error)
	GetFiles() []string
	GetAllCompilationUnits(context.Context) <-chan *policy.CompilationUnit
	GetAllCompilationUnitsWithCount(context.Context) (int, <-chan *policy.CompilationUnit)
	Clear() error
	InspectPolicies(context.Context, ...string) (map[string]*responsev1.InspectPoliciesResponse_Result, error)
	ListPolicyIDs(context.Context, ...string) ([]string, error)
	ListSchemaIDs(context.Context) ([]string, error)
	LoadSchema(context.Context, string) (io.ReadCloser, error)
	LoadPolicy(context.Context, ...string) ([]*policy.Wrapper, error)
	Reload(context.Context) ([]storage.Event, error)
}

type index struct {
	fsys         fs.FS
	sfGroup      singleflight.Group
	fileToModID  map[string]namer.ModuleID
	executables  ModuleIDSet
	dependents   map[namer.ModuleID]ModuleIDSet
	dependencies map[namer.ModuleID]ModuleIDSet
	modIDToFile  map[namer.ModuleID]string
	schemaLoader *SchemaLoader
	buildOpts    buildOptions
	stats        storage.RepoStats
	mu           sync.RWMutex
}

func (idx *index) GetFiles() []string { _ = "STUB: not implemented"; return nil }

func (idx *index) GetFirstMatch(candidates []namer.ModuleID) (*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) GetAll(ctx context.Context) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) GetAllMatching(modIDs []namer.ModuleID) ([]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) GetCompilationUnits(ids ...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add dependencies

func (idx *index) addDepsToCompilationUnit(cu *policy.CompilationUnit, id namer.ModuleID) error {
	_ = "STUB: not implemented"
	return nil
}

func (idx *index) loadPolicy(id namer.ModuleID) (*policyv1.Policy, parser.SourceCtx, error) {
	_ = "STUB: not implemented"
	return nil, *new(parser.SourceCtx), nil
}

func (idx *index) GetDependents(ids ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) addTransitiveDependents(dependents map[namer.ModuleID]struct{}, id namer.ModuleID) {
	_ = "STUB: not implemented"
	return
}

func (idx *index) AddOrUpdate(entry Entry) (evt storage.Event, err error) {
	_ = "STUB: not implemented"
	return *new(storage.Event), nil
}

// Is this is a duplicate of another file?

// if this is an existing file, clear its state first

// go through the dependencies and remove self from the dependents list of each dependency.

// add to index

func (idx *index) addDep(child, parent namer.ModuleID) {
	_ = "STUB: not implemented"
	// When we compile a policy, we need to load the dependencies (imported constants, variables, and derived roles).
	return
}

// if a derived role or variable export changes, we need to recompile all the policies that import it (dependents).

func (idx *index) Delete(entry Entry) (storage.Event, error) {
	_ = "STUB: not implemented"
	return *new(storage.Event), nil
}

// nothing to do because we don't have that file in the index.

// go through the dependencies and remove self from the dependents list for each dependency.

func (idx *index) GetAllCompilationUnits(ctx context.Context) <-chan *policy.CompilationUnit {
	_ = "STUB: not implemented"
	return nil
}

func (idx *index) GetAllCompilationUnitsWithCount(ctx context.Context) (int, <-chan *policy.CompilationUnit) {
	_ = "STUB: not implemented"
	return 0, nil
}

// is this a policy that is referenced by another one? If so, it will be implicitly compiled.

// No implicit compilation for this policy so add it to the list

func (idx *index) Clear() error { _ = "STUB: not implemented"; return nil }

type meta struct {
	Dependencies []string
	Dependents   []string
}

func (idx *index) Inspect() map[string]meta { _ = "STUB: not implemented"; return nil }

func (idx *index) InspectPolicies(ctx context.Context, file ...string) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) ListPolicyIDs(_ context.Context, filteredFiles ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) ListSchemaIDs(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) LoadSchema(ctx context.Context, url string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (idx *index) LoadPolicy(_ context.Context, file ...string) ([]*policy.Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) RepoStats(_ context.Context) storage.RepoStats {
	_ = "STUB: not implemented"
	return *new(storage.RepoStats)
}

func (idx *index) Reload(ctx context.Context) ([]storage.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (idx *index) Close() error { _ = "STUB: not implemented"; return nil }
