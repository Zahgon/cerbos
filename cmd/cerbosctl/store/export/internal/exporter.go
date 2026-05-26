// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
)

func NewExporter(path string) (Exporter, error) {
	_ = "STUB: not implemented"
	return *new(Exporter), nil
}

type Exporter interface {
	WriteJSON(name string, jsonData []byte) error
	WriteYAML(name string, jsonData []byte) error
	io.Closer
}

func newTarExporter(w io.WriteCloser) *tarExporter { _ = "STUB: not implemented"; return nil }

type tarExporter struct {
	archiveWriter *tar.Writer
	writer        io.WriteCloser
}

func (e *tarExporter) WriteJSON(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tarExporter) WriteYAML(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tarExporter) Close() error { _ = "STUB: not implemented"; return nil }

func (e *tarExporter) write(name string, data []byte) error { _ = "STUB: not implemented"; return nil }

func newGzipExporter(w io.WriteCloser) *gzipExporter { _ = "STUB: not implemented"; return nil }

type gzipExporter struct {
	archiveWriter *tar.Writer
	gzipWriter    *gzip.Writer
	writer        io.WriteCloser
}

func (e *gzipExporter) WriteJSON(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *gzipExporter) WriteYAML(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *gzipExporter) Close() error { _ = "STUB: not implemented"; return nil }

func (e *gzipExporter) write(name string, data []byte) error { _ = "STUB: not implemented"; return nil }

func newZipExporter(w io.WriteCloser) *zipExporter { _ = "STUB: not implemented"; return nil }

type zipExporter struct {
	archiveWriter *zip.Writer
	writer        io.WriteCloser
}

func (e *zipExporter) WriteJSON(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *zipExporter) WriteYAML(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *zipExporter) Close() error { _ = "STUB: not implemented"; return nil }

func (e *zipExporter) write(name string, data []byte) error { _ = "STUB: not implemented"; return nil }

func newDirectoryExporter(path string) *directoryExporter { _ = "STUB: not implemented"; return nil }

type directoryExporter struct {
	path string
}

func (e *directoryExporter) WriteJSON(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *directoryExporter) WriteYAML(name string, jsonData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *directoryExporter) Close() error { _ = "STUB: not implemented"; return nil }

func (e *directoryExporter) write(name string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
