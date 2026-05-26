// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

type FilterDimension string

const (
	FilterDimensionSuite     FilterDimension = "suite"
	FilterDimensionTest      FilterDimension = "test"
	FilterDimensionPrincipal FilterDimension = "principal"
	FilterDimensionResource  FilterDimension = "resource"
	FilterDimensionAction    FilterDimension = "action"

	expectedKeyValueParts = 2
)

var validDimensions = map[FilterDimension]struct{}{
	FilterDimensionSuite:     {},
	FilterDimensionTest:      {},
	FilterDimensionPrincipal: {},
	FilterDimensionResource:  {},
	FilterDimensionAction:    {},
}

type FilterConfig struct {
	Suite     []string
	Test      []string
	Principal []string
	Resource  []string
	Action    []string
}

func (fc *FilterConfig) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func toFilterDimension(s string) FilterDimension {
	_ = "STUB: not implemented"
	return *new(FilterDimension)
}

func parseDimension(s string) (FilterDimension, []string, error) {
	_ = "STUB: not implemented"
	return *new(FilterDimension), nil, nil
}

// ParseFilterConfig parses a filter string in the format:
// "test=glob1,glob2;principal=glob3;resource=glob4;action=glob5".
func ParseFilterConfig(filter string) (*FilterConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseGlobs(s string) []string { _ = "STUB: not implemented"; return nil }

func (fc *FilterConfig) Merge(other *FilterConfig) *FilterConfig {
	_ = "STUB: not implemented"
	return nil
}
