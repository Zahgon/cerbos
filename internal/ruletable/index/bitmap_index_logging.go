// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package index

import (
	"go.uber.org/zap"
)

// dimStats holds size statistics for a single dimension.
type dimStats struct { //nolint:unused
	Name     string
	Keys     int
	MinWords int
	MaxWords int
	AvgWords int
	MinCard  uint64
	MaxCard  uint64
	AvgCard  uint64
}

func collectBitmapStats(s *dimStats, bm *Bitmap) {
	_ = "STUB: not implemented" //nolint:unused
	return
}

func dimensionStats[T comparable](name string, d dimension[T]) dimStats {
	_ = "STUB: not implemented" //nolint:unused
	return *new(dimStats)
}

func globDimensionStats(name string, gd *globDimension) dimStats {
	_ = "STUB: not implemented" //nolint:unused
	return *new(dimStats)
}

func fqnDimensionStats(name string, d fqnDimension) dimStats {
	_ = "STUB: not implemented" //nolint:unused
	return *new(dimStats)
}

func (idx *bitmapIndex) logStats(log *zap.SugaredLogger) {
	_ = "STUB: not implemented" //nolint:unused
	return
}

//nolint:mnd
