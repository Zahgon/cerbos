// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package compilation

import (
	"github.com/cerbos/cerbos/cmd/cerbos/compile/internal/flagset"
	"github.com/cerbos/cerbos/internal/compile"
	"github.com/cerbos/cerbos/internal/outputcolor"
	"github.com/cerbos/cerbos/internal/printer"
)

func Display(p *printer.Printer, errs compile.ErrorSet, output flagset.OutputFormat, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func displayJSON(p *printer.Printer, errs compile.ErrorSet, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func displayList(p *printer.Printer, errs compile.ErrorSet) error {
	_ = "STUB: not implemented"
	return nil
}
