// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package lint

import (
	"github.com/cerbos/cerbos/cmd/cerbos/compile/internal/flagset"
	"github.com/cerbos/cerbos/internal/outputcolor"
	"github.com/cerbos/cerbos/internal/printer"
	"github.com/cerbos/cerbos/internal/storage/index"
)

func Display(p *printer.Printer, errs *index.BuildError, output flagset.OutputFormat, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func displayJSON(p *printer.Printer, errs *index.BuildError, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func displayList(p *printer.Printer, errs *index.BuildError) error {
	_ = "STUB: not implemented"
	return nil
}
