// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package auxdata

import (
	"context"
	"errors"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
)

var ErrFailedToExtractJWT = errors.New("failed to extract JWT")

type AuxData struct {
	jwt *jwtHelper
}

func New(ctx context.Context) (*AuxData, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFromConf(ctx context.Context, conf *Conf) *AuxData { _ = "STUB: not implemented"; return nil }

func NewWithoutVerification(ctx context.Context) *AuxData { _ = "STUB: not implemented"; return nil }

// Extract auxiliary data and convert to format expected by the engine.
func (ad *AuxData) Extract(ctx context.Context, adProto *requestv1.AuxData) (*enginev1.AuxData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
