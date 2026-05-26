// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package decisions

import (
	"io"

	"github.com/alecthomas/kong"
	cmdclient "github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/flagset"

	"github.com/alecthomas/chroma/v2"
	"github.com/rivo/tview"
	"google.golang.org/protobuf/proto"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
)

var help = `Requires audit logging to be enabled on the server. Supports several ways of filtering the data.

tail: View the last N records
between: View records captured between two timestamps. The timestamps must be formatted as ISO-8601
since: View records from X hours/minutes/seconds ago to now. Unit suffixes are: h=hours, m=minutes s=seconds
lookup: View a specific record using the Cerbos Call ID

# View the last 10 records
cerbosctl decisions --tail=10

# View the logs from midnight 2021-07-01 to midnight 2021-07-02
cerbosctl decisions --between=2021-07-01T00:00:00Z,2021-07-02T00:00:00Z

# View the logs from midnight 2021-07-01 to now
cerbosctl decisions --between=2021-07-01T00:00:00Z

# View the logs from 3 hours ago to now
cerbosctl decisions --since=3h --raw

# View a specific log entry by call ID
cerbosctl decisions--lookup=01F9Y5MFYTX7Y87A30CTJ2FB0S`

type Cmd struct { //betteralign:ignore
	flagset.AuditFilters
}

func (c *Cmd) Run(_ *kong.Kong, ctx *cmdclient.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }

func (c *Cmd) Validate() error { _ = "STUB: not implemented"; return nil }

const (
	browserKey      = "browser"
	checkDetailsKey = "checkDetails"
	planDetailsKey  = "planDetails"
)

type decisionsUI struct {
	app          *tview.Application
	tabs         *tview.Pages
	browser      *browserPanel
	checkDetails *checkDetailsPanel
	planDetails  *planDetailsPanel
}

type browserPanel struct {
	entriesTable *tview.Table
	jsonView     *tview.TextView
	focusOrder   []tview.Primitive
}

type checkDetailsPanel struct {
	inputsList       *tview.List
	principalView    *tview.TextView
	resourceView     *tview.TextView
	actionsTable     *tview.Table
	derivedRolesView *tview.TextView
	focusOrder       []tview.Primitive
}

type planDetailsPanel struct {
	principalView *tview.TextView
	resourceView  *tview.TextView
	planView      *tview.TextView
	debugView     *tview.TextView
	focusOrder    []tview.Primitive
}

func mkUI(entries []*auditv1.DecisionLogEntry) *decisionsUI { _ = "STUB: not implemented"; return nil }

//nolint:mnd
func mkBrowserPanel(ui *decisionsUI, entries []*auditv1.DecisionLogEntry) {
	_ = "STUB: not implemented"
	return
}

//nolint:mnd
func populateEntriesTable(ui *decisionsUI, entries []*auditv1.DecisionLogEntry) {
	_ = "STUB: not implemented"
	return
}

//nolint:mnd
func mkCheckDetailsPanel(ui *decisionsUI) { _ = "STUB: not implemented"; return }

//nolint:mnd
func mkPlanDetailsPanel(ui *decisionsUI) { _ = "STUB: not implemented"; return }

func (d *decisionsUI) Start() error { _ = "STUB: not implemented"; return nil }

func (d *decisionsUI) entrySelectedFunc(row, _ int) { _ = "STUB: not implemented"; return }

func (d *decisionsUI) showDetailsPanel(entry *auditv1.DecisionLogEntry) {
	_ = "STUB: not implemented"
	return
}

//nolint:mnd
func (d *decisionsUI) showCheckDetailsPanel(entry *auditv1.DecisionLogEntry_CheckResources) {
	_ = "STUB: not implemented"
	return
}

func (d *decisionsUI) showPlanDetailsPanel(entry *auditv1.DecisionLogEntry_PlanResources) {
	_ = "STUB: not implemented"
	return
}

func (bp *browserPanel) switchFocus(backward bool) tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (cdp *checkDetailsPanel) switchFocus(backward bool) tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (pdp *planDetailsPanel) switchFocus(backward bool) tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func switchFocus(items []tview.Primitive, backward bool) tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func keyCode(key string) string { _ = "STUB: not implemented"; return "" }

func keyDesc(key, desc string) string { _ = "STUB: not implemented"; return "" }

type prettyJSON struct {
	lexer     chroma.Lexer
	formatter chroma.Formatter
	style     *chroma.Style
}

func newPrettyJSON() *prettyJSON { _ = "STUB: not implemented"; return nil }

func (p *prettyJSON) write(out io.Writer, msg proto.Message) { _ = "STUB: not implemented"; return }
