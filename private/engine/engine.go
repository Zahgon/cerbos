// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package engine

import (
	"context"

	"github.com/cerbos/cloud-api/bundle"

	"github.com/cerbos/cerbos/internal/engine"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/internal/storage/hub"
)

type (
	Conf         = evaluator.Conf
	BundleParams = hub.LocalParams
	Engine       = engine.Engine
)

const (
	BundleVersion1 = bundle.Version1
	BundleVersion2 = bundle.Version2
)

func FromBundle(ctx context.Context, params BundleParams) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
