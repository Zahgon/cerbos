// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"github.com/alecthomas/participle/v2"
)

func NewParser() (*participle.Parser[REPLDirective], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:govet
type REPLDirective struct { //betteralign:ignore
	Exit  bool           `parser:"@('q'|'quit'|'exit')"`
	Reset bool           `parser:"| @'reset'"`
	Vars  bool           `parser:"| @'vars'"`
	Help  bool           `parser:"| @('h' | 'help')"`
	Rules bool           `parser:"| @'rules'"`
	Load  *LoadDirective `parser:"| @@"`
	Exec  *ExecDirective `parser:"| @@"`
	Let   *LetDirective  `parser:"| @@"`
}

type LetDirective struct {
	Name string `parser:"'let' @Ident"`
	Expr string `parser:"'=' @Any"`
}

type LoadDirective struct {
	Path string `parser:"'load' @(Path|Ident)"`
}

type ExecDirective struct {
	RuleID int `parser:"'exec' '#'@Int"`
}
