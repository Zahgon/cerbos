// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package files

import (
	"io/fs"

	"github.com/cerbos/cerbos/internal/util"
)

type callback func(file Found) error

func Find(paths []string, recursive bool, fileType util.IndexedFileType, callback callback) error {
	_ = "STUB: not implemented"
	return nil
}

func find(path string, recursive bool, fileType util.IndexedFileType, callback callback) error {
	_ = "STUB: not implemented"
	return nil
}

func doFind(fsys fs.FS, fileType util.IndexedFileType, recursive bool, callback callback) error {
	_ = "STUB: not implemented"
	return nil
}

func isSupportedFile(fileName string, fileType util.IndexedFileType) bool {
	_ = "STUB: not implemented"
	return false
}
