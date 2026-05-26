// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package repl

import (
	"io"

	"github.com/alecthomas/kong"
	"github.com/peterh/liner"
)

type Cmd struct { //betteralign:ignore
	History string `help:"Path to history file" type:"path"`
}

func (c *Cmd) clear(stdout io.Writer) { _ = "STUB: not implemented"; return }

func (c *Cmd) Run(k *kong.Kong) error { _ = "STUB: not implemented"; return nil }

func getHistoryFile(path string) string { _ = "STUB: not implemented"; return "" }

//nolint:mnd

func loadHistory(reader *liner.State, histFile string) error { _ = "STUB: not implemented"; return nil }

func writeHistory(reader *liner.State, histFile string) error {
	_ = "STUB: not implemented"
	return nil
}
