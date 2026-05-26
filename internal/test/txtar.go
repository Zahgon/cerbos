// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"io/fs"
	"testing"
)

func ExtractTxtArchiveToFS(t *testing.T, path string) fs.FS {
	_ = "STUB: not implemented"
	return *new(fs.FS)
}

func ExtractTxtArchiveToDir(t *testing.T, path, out string) { _ = "STUB: not implemented"; return }
