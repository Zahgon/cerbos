// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package flagset

import (
	"errors"
	"time"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
)

var errMoreThanOneFilter = errors.New("more than one filter specified: choose from either `tail`, `between`, `since` or `lookup`")

type AuditFilters struct {
	Lookup  string        `help:"View a specific record using the Cerbos Call ID"`
	Between timerange     `help:"View records captured between two timestamps. The timestamps must be formatted as ISO-8601"`
	Since   time.Duration `help:"View records from X hours/minutes/seconds ago to now. Unit suffixes are: h=hours, m=minutes s=seconds"`
	Tail    uint16        `help:"View the last N records"`
}

func (af *AuditFilters) Validate() error { _ = "STUB: not implemented"; return nil }

func (af *AuditFilters) GenOptions() cerbos.AuditLogOptions {
	_ = "STUB: not implemented"
	return *new(cerbos.AuditLogOptions)
}
