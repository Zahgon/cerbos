// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package traces

import (
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/printer"
)

type Map map[string]*enginev1.TraceBatch

func (m *Map) Add(suiteName, principalName, resourceName, actionName string, batch *enginev1.TraceBatch) {
	_ = "STUB: not implemented"
	return
}

func (m *Map) Print(p *printer.Printer) { _ = "STUB: not implemented"; return }
