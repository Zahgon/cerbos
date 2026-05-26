// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	vtgrpc "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc/encoding"

	// Import the default grpc encoding to ensure that it gets replaced by this codec.
	"google.golang.org/grpc/encoding/gzip"
	_ "google.golang.org/grpc/encoding/proto"
)

const name = "proto"

func init() {
	// Register the codec to use VT where possible for optimized marshaling/unmarshaling.
	encoding.RegisterCodec(Codec{vtcodec: vtgrpc.Codec{}})
	// Register gzip compressor.
	encoding.RegisterCompressor(encoding.GetCompressor(gzip.Name))
}

// Codec implements the grpc Codec interface to delegate encoding to VT where possible.
type Codec struct {
	vtcodec vtgrpc.Codec
}

func (c Codec) Name() string { _ = "STUB: not implemented"; return "" }

func (c Codec) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c Codec) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }
