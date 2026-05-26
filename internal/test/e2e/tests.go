// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build e2e

package e2e

import (
	"testing"
	"time"
)

const (
	AdminSuite         = "admin"
	ChecksSuite        = "checks"
	PlanResourcesSuite = "plan_resources"
	testTimeout        = 90 * time.Second // Things are slower inside Kind

	adminSuiteSleepDuration = time.Millisecond * 500
)

type Opt func(*suiteOpt)

type suiteOpt struct {
	contextID         string
	suites            []string
	computedEnv       func(Ctx) map[string]string
	postSetup         func(Ctx)
	tlsDisabled       bool
	overlayMaxRetries uint
}

func WithContextID(contextID string) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithSuites(suites ...string) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithComputedEnv(fn func(Ctx) map[string]string) Opt {
	_ = "STUB: not implemented"
	return *new(Opt)
}

func WithPostSetup(fn func(Ctx)) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithMutableStoreSuites() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithImmutableStoreSuites() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithTLSDisabled() Opt { _ = "STUB: not implemented"; return *new(Opt) }

func WithOverlayMaxRetries(nRetries uint) Opt { _ = "STUB: not implemented"; return *new(Opt) }

func RunSuites(t *testing.T, opts ...Opt) { _ = "STUB: not implemented"; return }

//nolint:gosec
