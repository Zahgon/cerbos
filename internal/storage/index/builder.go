// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package index

import (
	"context"
	"fmt"
	"io/fs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	sourcev1 "github.com/cerbos/cerbos/api/genpb/cerbos/source/v1"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/parser"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/util"
)

// errSchemasInWrongDir signals that schemas folder is in the wrong place.
var errSchemasInWrongDir = fmt.Errorf("%s directory must be under the root of the storage directory", util.SchemasDirectory)

const maxLoggableBuildErrors = 5

// BuildError is an error type that contains details about the failures encountered during the index build.
type BuildError struct {
	*runtimev1.IndexBuildErrors
	nErr int
}

func (ibe *BuildError) Error() string { _ = "STUB: not implemented"; return "" }

type buildOptions struct {
	rootDir              string
	sourceAttributes     []policy.SourceAttribute
	buildFailureLogLevel zapcore.Level
}

type BuildOpt func(*buildOptions)

func WithBuildFailureLogLevel(level zapcore.Level) BuildOpt {
	_ = "STUB: not implemented"
	return *new(BuildOpt)
}

func WithRootDir(rootDir string) BuildOpt { _ = "STUB: not implemented"; return *new(BuildOpt) }

func WithSourceAttributes(attrs ...policy.SourceAttribute) BuildOpt {
	_ = "STUB: not implemented"
	return *new(BuildOpt)
}

func mkBuildOpts(opts ...BuildOpt) buildOptions {
	_ = "STUB: not implemented"
	return *new(buildOptions)
}

// Build builds an index from the policy files stored in a directory.
func Build(ctx context.Context, fsys fs.FS, opts ...BuildOpt) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

func build(ctx context.Context, fsys fs.FS, opts buildOptions) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

type indexBuilder struct {
	executables                   ModuleIDSet
	modIDToFile                   map[namer.ModuleID]string
	fileToModID                   map[string]namer.ModuleID
	dependents                    map[namer.ModuleID]ModuleIDSet
	dependencies                  map[namer.ModuleID]ModuleIDSet
	missingScopes                 map[string]map[string]struct{}
	sharedScopePermissionGroups   map[string]map[policyv1.ScopePermissions]struct{}
	conflictingScopes             map[string]struct{}
	missingResourceScopes         map[string]map[string]map[string]struct{} // map[{resource}]map[{scope}]map[{version}]struct{}
	foundRolePolicyResourceScopes map[string]map[string]map[string]struct{} // map[{resource}]map[{scope}]map[{version}]struct{}
	missing                       map[namer.ModuleID][]*runtimev1.IndexBuildErrors_MissingImport
	stats                         *statsCollector
	duplicates                    []*runtimev1.IndexBuildErrors_DuplicateDef
	loadFailures                  []*runtimev1.IndexBuildErrors_LoadFailure
	disabled                      []*runtimev1.IndexBuildErrors_Disabled
}

func newIndexBuilder() *indexBuilder { _ = "STUB: not implemented"; return nil }

func (idx *indexBuilder) addLoadFailure(file string, err error) {
	_ = "STUB: not implemented"
	//nolint:errorlint
	return
}

func (idx *indexBuilder) addErrors(file string, errs []*sourcev1.Error) {
	_ = "STUB: not implemented"
	return
}

func (idx *indexBuilder) addDisabled(file string, srcCtx parser.SourceCtx, p *policyv1.Policy) {
	_ = "STUB: not implemented"
	return
}

func (idx *indexBuilder) addPolicy(file string, srcCtx parser.SourceCtx, p policy.Wrapper) {
	_ = "STUB: not implemented"
	// Is this policy defined elsewhere?
	return
}

// Record that this role policy combination exists

// not executable

// the dependent may not have been loaded by the indexer yet because it's still walking the directory.

// check to see if matching role policies (with a rule for the given resource) reside in any of the missing scopes
//nolint:nestif

func (idx *indexBuilder) addDep(child, parent namer.ModuleID) {
	_ = "STUB: not implemented"
	// When we compile a policy, we need to load the dependencies (imported variables and derived roles).
	return
}

// if a derived role or variable export changes, we need to recompile all the policies that import it (dependents).

func (idx *indexBuilder) build(fsys fs.FS, opts buildOptions) (*index, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func logBuildFailure(logger *zap.Logger, level zapcore.Level, err *BuildError) {
	_ = "STUB: not implemented"
	return
}

func checkValidDir(fsys fs.FS, dir string) error { _ = "STUB: not implemented"; return nil }
