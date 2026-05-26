// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package util

var (
	AppName   = "cerbos"
	BuildDate = "unknown"
	Commit    = "unknown"
	Version   = "unknown"

	appID = []byte(AppName)
)

const nodeIDLen = 16

func AppVersion() string { _ = "STUB: not implemented"; return "" }

func AppShortVersion() string { _ = "STUB: not implemented"; return "" }
