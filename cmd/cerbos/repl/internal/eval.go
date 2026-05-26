// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"errors"

	"github.com/pterm/pterm"
)

var errEvalErrorPresent = errors.New("failed expression(s) present")

type evalOutput struct {
	tree  pterm.LeveledList
	level int
}

func (eo *evalOutput) append(text string) { _ = "STUB: not implemented"; return }

func (eo *evalOutput) appendAndLevelUp(text string) { _ = "STUB: not implemented"; return }

func buildEvalOutput(e *eval) *evalOutput { _ = "STUB: not implemented"; return nil }

func successText(success bool) string { _ = "STUB: not implemented"; return "" }

func doBuildEvalOutput(eo *evalOutput, e *eval) { _ = "STUB: not implemented"; return }

type evalType string

const (
	evalTypeAll  evalType = "all"
	evalTypeAny  evalType = "any"
	evalTypeExpr evalType = "expr"
	evalTypeNone evalType = "none"
)

type eval struct {
	err      error
	evalType evalType
	expr     string
	evals    []*eval
	success  bool
}

func (e *eval) append(eval *eval) { _ = "STUB: not implemented"; return }
