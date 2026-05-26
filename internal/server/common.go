// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/auxdata"
	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/svc"
)

// CoreComponents holds the shared components needed for both server and Lambda function initialization.
type CoreComponents struct {
	Engine     *engine.Engine
	AuxData    *auxdata.AuxData
	AuditLog   audit.Log
	Store      storage.Store
	ReqLimits  svc.RequestLimits
	SuggestHub bool
}

// InitializeCerbosCore performs the common initialization steps shared between server and Lambda function.
func InitializeCerbosCore(ctx context.Context) (*CoreComponents, error) {
	_ = "STUB: not implemented"
	// create audit log
	return nil, nil
}

// create store

// Overlay needs to take precedence over BinaryStore in this type switch,
// as our overlay store implements BinaryStore also

// create wrapped policy loader

// create compile manager

//nolint:nestif

// create engine

// initialize aux data
