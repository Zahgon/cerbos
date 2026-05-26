// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package printer

import (
	"io"

	"github.com/alecthomas/chroma/v2/styles"
	"google.golang.org/protobuf/proto"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	"github.com/cerbos/cerbos/internal/outputcolor"
)

var style = styles.Get("solarized-dark256")

func New(stdout, stderr io.Writer) *Printer { _ = "STUB: not implemented"; return nil }

type Printer struct {
	stdout io.Writer
	stderr io.Writer
}

func (p *Printer) Println(args ...any) { _ = "STUB: not implemented"; return }

func (p *Printer) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (p *Printer) coloredJSON(data string, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Printer) PrintJSON(val any, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Printer) coloredYAML(data string, colorLevel outputcolor.Level) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Printer) PrintProtoYAML(message proto.Message, colorLevel outputcolor.Level, indent int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Printer) PrintProtoJSON(message proto.Message, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Printer) PrintTraceEntry(definitions []*enginev1.Trace_Component, trace *enginev1.TraceEntry) {
	_ = "STUB: not implemented"
	return
}

func (p *Printer) PrintTrace(trace *enginev1.Trace) { _ = "STUB: not implemented"; return }

func (p *Printer) printTraceEntryComponents(definitions []*enginev1.Trace_Component, componentIndices []uint32) {
	_ = "STUB: not implemented"
	return
}

//nolint:govet

func (p *Printer) printTraceComponents(components []*enginev1.Trace_Component) {
	_ = "STUB: not implemented"
	return
}

//nolint:govet

func (p *Printer) printTraceComponent(component *enginev1.Trace_Component) {
	_ = "STUB: not implemented"
	return
}

//nolint:govet

//nolint:govet

//nolint:govet

//nolint:govet

//nolint:govet

//nolint:govet

func (p *Printer) printTraceEvent(event *enginev1.Trace_Event) { _ = "STUB: not implemented"; return }
