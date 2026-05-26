// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package audit

import (
	"bufio"
	"io"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/kong"
	"google.golang.org/protobuf/proto"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	cmdclient "github.com/cerbos/cerbos/cmd/cerbosctl/internal/client"
	"github.com/cerbos/cerbos/cmd/cerbosctl/internal/flagset"
)

var newline = []byte("\n")

const (
	dashLen = 54
	help    = `View audit logs.
Requires audit logging to be enabled on the server. Supports several ways of filtering the data.

tail: View the last N records
between: View records captured between two timestamps. The timestamps must be formatted as ISO-8601
since: View records from X hours/minutes/seconds ago to now. Unit suffixes are: h=hours, m=minutes s=seconds
lookup: View a specific record using the Cerbos Call ID

# View the last 10 access logs
cerbosctl audit --kind=access --tail=10

# View the decision logs from midnight 2021-07-01 to midnight 2021-07-02
cerbosctl audit --kind=decision --between=2021-07-01T00:00:00Z,2021-07-02T00:00:00Z

# View the decision logs from midnight 2021-07-01 to now
cerbosctl audit --kind=decision --between=2021-07-01T00:00:00Z

# View the access logs from 3 hours ago to now as newline-delimited JSON
cerbosctl audit --kind=access --since=3h --raw

# View a specific access log entry by call ID
cerbosctl audit --kind=access --lookup=01F9Y5MFYTX7Y87A30CTJ2FB0S`
)

type Cmd struct { //betteralign:ignore
	Kind string `default:"access" enum:"access,decision" help:"Kind of log entry (${enum})"`
	flagset.AuditFilters
	Raw bool `help:"Output results without formatting or colours"`
}

func (c *Cmd) Run(k *kong.Kong, ctx *cmdclient.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cmd) Help() string { _ = "STUB: not implemented"; return "" }

func (c *Cmd) Validate() error { _ = "STUB: not implemented"; return nil }

func streamLogsToWriter(writer auditLogWriter, entries <-chan *cerbos.AuditLogEntry) error {
	_ = "STUB: not implemented"
	return nil
}

type auditLogWriter interface {
	write(proto.Message) error
	flush()
}

func newRawAuditLogWriter(out io.Writer) *rawAuditLogWriter { _ = "STUB: not implemented"; return nil }

type rawAuditLogWriter struct {
	out io.Writer
}

func (r *rawAuditLogWriter) write(entry proto.Message) error { _ = "STUB: not implemented"; return nil }

func (r *rawAuditLogWriter) flush() { _ = "STUB: not implemented"; return }

type richAuditLogWriter struct {
	out       *bufio.Writer
	lexer     chroma.Lexer
	formatter chroma.Formatter
	rowStyle  func(...string) string
	jsonStyle *chroma.Style
}

func newRichAuditLogWriter(out io.Writer) *richAuditLogWriter {
	_ = "STUB: not implemented"
	return nil
}

func (r *richAuditLogWriter) write(entry proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *richAuditLogWriter) header(h string) { _ = "STUB: not implemented"; return }

func (r *richAuditLogWriter) formattedJSON(msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *richAuditLogWriter) flush() { _ = "STUB: not implemented"; return }
