// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package awslambda

func MkConfStorageOverrides(cwd string, confOverrides map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func mkConfServerOverrides(confOverrides map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func mkConfStorageHubOverrides(tmpDir string, confOverrides map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

const tempDir string = "/tmp" // Lambda tempDir

func GetConfOverrides(confOverrides map[string]any) error { _ = "STUB: not implemented"; return nil }

func HubStorageDriver(confOverrides map[string]any) bool { _ = "STUB: not implemented"; return false }
