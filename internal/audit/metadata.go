// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit

import (
	"context"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
)

var alwaysExcludeMetadataKeys = map[string]struct{}{
	"authorization":  {},
	"grpc-trace-bin": {},
}

type MetadataExtractor func(context.Context) map[string]*auditv1.MetaValues

func NewMetadataExtractor() (MetadataExtractor, error) {
	_ = "STUB: not implemented"
	return *new(MetadataExtractor), nil
}

func NewMetadataExtractorFromConf(conf *Conf) MetadataExtractor {
	_ = "STUB: not implemented"
	return *new(MetadataExtractor)
}

func sliceToLookupMap(slice []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }
