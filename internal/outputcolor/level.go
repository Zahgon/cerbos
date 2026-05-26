// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package outputcolor

import (
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/jwalton/go-supportscolor"
)

type Level uint8

const (
	None    = Level(supportscolor.None)
	Basic   = Level(supportscolor.Basic)
	Ansi256 = Level(supportscolor.Ansi256)
	Ansi16m = Level(supportscolor.Ansi16m)
)

func DefaultLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

var TypeMapper = kong.TypeMapper(reflect.TypeFor[*Level](), kong.MapperFunc(decode))

func (l *Level) Resolve(disable bool) Level { _ = "STUB: not implemented"; return *new(Level) }

func (l Level) Enabled() bool { _ = "STUB: not implemented"; return false }

func decode(ctx *kong.DecodeContext, target reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func scan(ctx *kong.DecodeContext) (*Level, error) { _ = "STUB: not implemented"; return nil, nil }

func parse(v any) (*Level, error) { _ = "STUB: not implemented"; return nil, nil }

func pointer(level Level) *Level { _ = "STUB: not implemented"; return nil }
