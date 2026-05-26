// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package internal

import (
	"context"

	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
)

func ConcatWithSepFunc(dialect string) func(string, ...any) exp.Expression {
	_ = "STUB: not implemented"
	return nil
}

func mysqlConcatWithSep(sep string, args ...any) exp.Expression {
	_ = "STUB: not implemented"
	return *new(exp.Expression)
}

//nolint:mnd
func ansiConcatWithSep(sep string, args ...any) exp.Expression {
	_ = "STUB: not implemented"
	return *new(exp.Expression)
}

func ConnectWithRetries(ctx context.Context, driverName, connStr string, retryConf *ConnRetryConf) (*sqlx.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
