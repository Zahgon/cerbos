// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	telemetryv1 "github.com/cerbos/cerbos/api/genpb/cerbos/telemetry/v1"
	"github.com/cerbos/cerbos/internal/storage"
)

func buildServerLaunch(store storage.Store) *telemetryv1.ServerLaunch {
	_ = "STUB: not implemented"
	return nil
}

func extractSource() *telemetryv1.ServerLaunch_Source { _ = "STUB: not implemented"; return nil }

func extractFeatures(store storage.Store) *telemetryv1.ServerLaunch_Features {
	_ = "STUB: not implemented"
	return nil
}

// avoid an import cycle by not using server.Conf to retrieve this value

//nolint:nestif

func extractStats(stats storage.RepoStats) *telemetryv1.ServerLaunch_Stats {
	_ = "STUB: not implemented"
	return nil
}
