// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package store

import (
	"archive/zip"
	"os"

	"github.com/alecthomas/kong"

	storev1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/store/v1"
)

const downloadHelp = `

# Download the store to a directory

cerbosctl hub store download /path/to/dir

# Download the store as a zip file

cerbosctl hub store download /path/to/archive.zip

`

type DownloadCmd struct { //betteralign:ignore
	Output     `embed:""`
	OutputPath string `arg:"" type:"path" required:"" help:"Path to write the retrieved files. Must be a path to a directory, zip file or - for stdout."`
}

func (*DownloadCmd) Help() string { _ = "STUB: not implemented"; return "" }

func (dc *DownloadCmd) Run(k *kong.Kong, cmd *Cmd) (outErr error) {
	_ = "STUB: not implemented"
	return nil
}

type fileWriter struct {
	outputRoot *os.Root
	zipWriter  *zip.Writer
	close      func() error
}

func newFileWriter(outputPath string) (*fileWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fw *fileWriter) writeFiles(files []*storev1.File) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:nestif

func (fw *fileWriter) Close() error { _ = "STUB: not implemented"; return nil }

func mkdirAll(root *os.Root, path string) error { _ = "STUB: not implemented"; return nil }
