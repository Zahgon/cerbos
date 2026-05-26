// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import (
	"github.com/alecthomas/kong"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type timerange struct {
	Values []*timestamppb.Timestamp
}

func (t *timerange) Decode(ctx *kong.DecodeContext) error { _ = "STUB: not implemented"; return nil }

//nolint:mnd

// default to current time if only one timestamp value is provided

func (t timerange) IsSet() bool { _ = "STUB: not implemented"; return false }
