// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package compile

import (
	"context"
	"io/fs"

	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/private/compile"
)

func Compile(ctx context.Context, fsys fs.FS, attrs ...compile.SourceAttribute) (*runtimev1.RuleTable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
