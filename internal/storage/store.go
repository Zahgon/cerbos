// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package storage

import (
	"context"
	"errors"
	"io"
	"sync"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	schemav1 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/namer"
	"github.com/cerbos/cerbos/internal/policy"
)

var (
	driversMu sync.RWMutex
	drivers   = map[string]Constructor{}
)

var (
	ErrPolicyIDCollision    = errors.New("policy ID collision")
	ErrUnsupportedOperation = errors.New("operation not supported")
)

// InvalidPolicyError is a custom error to signal that a policy is invalid.
type InvalidPolicyError struct {
	Err     error
	Message string
}

func (ipe InvalidPolicyError) Error() string { _ = "STUB: not implemented"; return "" }

func (ipe InvalidPolicyError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func NewInvalidPolicyError(err error, msg string, args ...any) InvalidPolicyError {
	_ = "STUB: not implemented"
	return *new(InvalidPolicyError)
}

// InvalidSchemaError is a custom error to signal that a schema is invalid.
type InvalidSchemaError struct {
	Err     error
	Message string
}

func (ise InvalidSchemaError) Error() string { _ = "STUB: not implemented"; return "" }

func (ise InvalidSchemaError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func NewInvalidSchemaError(err error, msg string, args ...any) InvalidSchemaError {
	_ = "STUB: not implemented"
	return *new(InvalidSchemaError)
}

// Constructor is a constructor function for a storage driver.
type Constructor func(context.Context, *config.Wrapper) (Store, error)

// RegisterDriver registers a storage driver.
func RegisterDriver(name string, cons Constructor) { _ = "STUB: not implemented"; return }

// GetDriverConstructor registers a storage driver.
func GetDriverConstructor(name string) (Constructor, error) {
	_ = "STUB: not implemented"
	return *new(Constructor), nil
}

// New returns a storage driver implementation based on the configured driver.
func New(ctx context.Context) (Store, error) { _ = "STUB: not implemented"; return *new(Store), nil }

// NewFromConf returns a storage driver implementation based on the provided configuration.
func NewFromConf(ctx context.Context, confWrapper *config.Wrapper) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

type ListPolicyIDsParams struct {
	NameRegexp      string
	ScopeRegexp     string
	VersionRegexp   string
	IDs             []string
	IncludeDisabled bool
}

// Store is the common interface implemented by storage backends.
type Store interface {
	// Driver is the name of the storage backend implementation.
	Driver() string
	// InspectPolicies returns inspection results for the policies in the store.
	InspectPolicies(context.Context, ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error)
	// ListPolicyIDs returns the policy IDs in the store.
	ListPolicyIDs(context.Context, ListPolicyIDsParams) ([]string, error)
	// ListSchemaIDs returns the schema ids in the store.
	ListSchemaIDs(context.Context) ([]string, error)
	// LoadSchema loads the given schema from the store.
	LoadSchema(context.Context, string) (io.ReadCloser, error)
	// Source returns metadata for inclusion in audit logs.
	Source() *auditv1.PolicySource
}

// SourceStore is implemented by stores that have policies in their source format (uncompiled).
type SourceStore interface {
	Store
	// GetFirstMatch searches for the given module IDs in order and returns the first one found.
	GetFirstMatch(context.Context, []namer.ModuleID) (*policy.CompilationUnit, error)
	// GetAll returns all modules that exist within the policy store
	GetAll(context.Context) ([]*policy.CompilationUnit, error)
	// GetAllMatching returns all modules that exist for the provided module IDs
	GetAllMatching(context.Context, []namer.ModuleID) ([]*policy.CompilationUnit, error)
	// GetCompilationUnits gets the compilation units for the given module IDs.
	GetCompilationUnits(context.Context, ...namer.ModuleID) (map[namer.ModuleID]*policy.CompilationUnit, error)
	// GetDependents returns the dependents of the given modules.
	GetDependents(context.Context, ...namer.ModuleID) (map[namer.ModuleID][]namer.ModuleID, error)
	// LoadPolicy loads the given policy from the store
	LoadPolicy(context.Context, ...string) ([]*policy.Wrapper, error)
}

// BinaryStore is implemented by stores that have pre-compiled policies in binary format.
type BinaryStore interface {
	Store
	Subscribable
	// GetFirstMatch searches for the given module IDs in order and returns the first one found.
	GetFirstMatch(context.Context, []namer.ModuleID) (*runtimev1.RunnablePolicySet, error)
	// GetAll returns all modules that exist within the policy store
	GetAll(context.Context) ([]*runtimev1.RunnablePolicySet, error)
	// GetAllMatching returns all modules that exist for the provided module IDs
	GetAllMatching(context.Context, []namer.ModuleID) ([]*runtimev1.RunnablePolicySet, error)
}

// MutableStore is a store that allows mutations.
type MutableStore interface {
	Store
	AddOrUpdate(context.Context, ...policy.Wrapper) error
	AddOrUpdateSchema(context.Context, ...*schemav1.Schema) error
	Disable(context.Context, ...string) (uint32, error)
	Enable(context.Context, ...string) (uint32, error)
	DeleteSchema(context.Context, ...string) (uint32, error)
	Delete(context.Context, ...string) (uint32, error)
	PurgeRevisions(context.Context, uint32) (uint32, error)
}

// Verifiable stores allow querying whether the requirements for the store are met.
type Verifiable interface {
	CheckSchema(ctx context.Context) error
}

// Reloadable stores allow reloading their contents.
type Reloadable interface {
	Reload(context.Context) error
}

// Instrumented stores expose repository stats.
type Instrumented interface {
	RepoStats(context.Context) RepoStats
}

// Subscribable is an interface for managing subscriptions to storage events.
type Subscribable interface {
	// Subscribe adds a subscriber to listen for storage notifications.
	Subscribe(Subscriber)
	// Unsubscribe removes a subscriber.
	Unsubscribe(Subscriber)
}

// Subscriber is the interface implemented by storage subscribers.
type Subscriber interface {
	SubscriberID() string
	OnStorageEvent(...Event)
}

// EventKind identifies the kind of storage event such as addition or deletion.
type EventKind int

const (
	EventAddOrUpdatePolicy EventKind = iota
	EventDeleteOrDisablePolicy
	EventAddOrUpdateSchema
	EventDeleteSchema
	EventReload
	EventNop
)

// Event is an event detected by the storage layer.
type Event struct {
	OldPolicyID *namer.ModuleID
	SchemaFile  string
	Dependents  []namer.ModuleID
	Kind        EventKind
	PolicyID    namer.ModuleID
}

func (evt Event) String() string { _ = "STUB: not implemented"; return "" }

// NewPolicyEvent creates a new storage event for a policy.
func NewPolicyEvent(kind EventKind, policyID namer.ModuleID) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

// NewSchemaEvent creates a new storage event for a schema.
func NewSchemaEvent(kind EventKind, schemaFile string) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

// NewReloadEvent creates a new reload event.
func NewReloadEvent() Event { _ = "STUB: not implemented"; return *new(Event) }

type RepoStats struct {
	PolicyCount           map[policy.Kind]int
	ConditionCount        map[policy.Kind]int
	RuleCount             map[policy.Kind]int
	MaxConditionCount     map[policy.Kind]int
	MaxRuleCount          map[policy.Kind]int
	AvgConditionCount     map[policy.Kind]float64
	AvgRuleCount          map[policy.Kind]float64
	DistinctActionCount   int
	DistinctResourceCount int
	SchemaCount           int
	HasOutput             bool
	HasScopedPolicies     bool
}
