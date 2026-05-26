// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"bufio"
	"errors"
	"io"

	"google.golang.org/protobuf/proto"
)

const (
	bufSize     = 1024 * 4        // 4KiB
	maxFileSize = 1024 * 1024 * 4 // 4MiB
	newline     = '\n'
)

var (
	jsonStart           = []byte("{")
	yamlSep             = []byte("---")
	yamlComment         = []byte("#")
	ErrEmptyFile        = errors.New("empty file")
	ErrMultipleYAMLDocs = errors.New("more than one YAML document detected")
)

func ReadJSONOrYAML(src io.Reader, dest proto.Message) error { _ = "STUB: not implemented"; return nil }

func IsJSON(src []byte) bool { _ = "STUB: not implemented"; return false }

func mkDecoder(src io.Reader) decoder { _ = "STUB: not implemented"; return *new(decoder) }

type decoder interface {
	decode(dest proto.Message) error
}

type decoderFunc func(dest proto.Message) error

func (df decoderFunc) decode(dest proto.Message) error { _ = "STUB: not implemented"; return nil }

func newJSONDecoder(src *bufio.Reader) decoderFunc {
	_ = "STUB: not implemented"
	return *new(decoderFunc)
}

func newYAMLDecoder(src *bufio.Reader) decoderFunc {
	_ = "STUB: not implemented"
	return *new(decoderFunc)
}

// ignore comments

// ignore empty lines at the beginning of the file

func WriteYAML(dest io.Writer, data proto.Message) error { _ = "STUB: not implemented"; return nil }
