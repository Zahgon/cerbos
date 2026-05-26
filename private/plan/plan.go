// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package plan

import (
	"context"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/evaluator"
	"github.com/cerbos/cerbos/private/compile"
)

func Resources(ctx context.Context, conf *evaluator.Conf, idx compile.Index, input *enginev1.PlanResourcesInput) (*enginev1.PlanResourcesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
