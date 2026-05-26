// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build e2e

package e2e

import (
	"context"
	"flag"
	"fmt"
	"strconv"
	"testing"
	"time"
)

const (
	cerbosHostNameEnvVar = "E2E_CERBOS_HOST"
	HTTPPort             = 3592
	GRPCPort             = 3593
	HealthEndpoint       = "/_cerbos/health"
)

var conf = Config{}

//nolint:mnd
func init() {
	srcRoot, err := findSourceRoot()
	if err != nil {
		panic(fmt.Errorf("failed to determine source root: %w", err))
	}

	noCleanup, err := strconv.ParseBool(envOrDefault("E2E_NO_CLEANUP", "false"))
	if err != nil {
		noCleanup = false
	}

	flag.StringVar(&conf.RunID, "run-id", envOrDefault("E2E_RUN_ID", RandomStr(5)), "Run ID for this test run")
	flag.StringVar(&conf.SourceRoot, "source-root", srcRoot, "Directory containing the Cerbos source code")
	flag.StringVar(&conf.CerbosImgRepo, "cerbos-img-repo", "ghcr.io/cerbos/cerbos", "Cerbos image repo")
	flag.StringVar(&conf.CerbosImgTag, "cerbos-img-tag", "dev", "Cerbos image tag")
	flag.DurationVar(&conf.CommandTimeout, "command-timeout", 3*time.Minute, "Command execution timeout")
	flag.BoolVar(&conf.NoCleanup, "no-cleanup", noCleanup, "Do not cleanup after tests")
}

func findSourceRoot() (string, error) { _ = "STUB: not implemented"; return "", nil }

func envOrDefault(envVar, def string) string { _ = "STUB: not implemented"; return "" }

type Config struct {
	RunID          string        `json:"run_id"`
	SourceRoot     string        `json:"source_root"`
	CerbosImgRepo  string        `json:"cerbos_img_repo"`
	CerbosImgTag   string        `json:"cerbos_img_tag"`
	CommandTimeout time.Duration `json:"command_timeout"`
	NoCleanup      bool          `json:"no_cleanup"`
}

func NewCtx(t *testing.T, contextID string, noTLS bool) Ctx {
	_ = "STUB: not implemented"
	return *new(Ctx)
}

type Ctx struct {
	ContextID   string
	NoTLS       bool
	ComputedEnv map[string]string
	*testing.T
	Config
}

func (c Ctx) Environ() []string { _ = "STUB: not implemented"; return nil }

//nolint:prealloc

// Remove conflicts. Env vars that already exist take precedence over our defaults.

func (c Ctx) CerbosHost() string { _ = "STUB: not implemented"; return "" }

func (c Ctx) Namespace() string { _ = "STUB: not implemented"; return "" }

func (c Ctx) GRPCAddr() string { _ = "STUB: not implemented"; return "" }

func (c Ctx) HTTPAddr() string { _ = "STUB: not implemented"; return "" }

func (c Ctx) HealthURL() string { _ = "STUB: not implemented"; return "" }

func (c Ctx) CommandTimeoutCtx() (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}
