// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package sqlite3

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"embed"
	"fmt"

	"github.com/doug-martin/goqu/v9"

	// import sqlite3 dialect.
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	gosqlite3 "modernc.org/sqlite"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/storage/db/internal"
	"github.com/cerbos/cerbos/internal/util"
)

const DriverName = "sqlite3"

//go:embed schema.sql
var schema string

//go:embed migrations/*.sql
var migrationsFS embed.FS

var (
	_ storage.SourceStore  = (*Store)(nil)
	_ storage.MutableStore = (*Store)(nil)
	_ storage.Subscribable = (*Store)(nil)
)

const nRegexpFnArgs = 2

var nameRegexpCache util.RegexpCache

func init() {
	nameRegexpCache = *util.NewRegexpCache()

	gosqlite3.MustRegisterDeterministicScalarFunction("regexp", nRegexpFnArgs, func(_ *gosqlite3.FunctionContext, args []driver.Value) (driver.Value, error) {
		if args[0] == nil || args[1] == nil {
			return nil, nil
		}

		re, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("arg[0] should be of type: string")
		}

		s, ok := args[1].(string)
		if !ok {
			return nil, fmt.Errorf("arg[1] should be of type: string")
		}

		r, err := nameRegexpCache.GetCompiledExpr(re)
		if err != nil {
			return nil, err
		}

		b := r.MatchString(s)
		return b, nil
	})

	storage.RegisterDriver(DriverName, func(ctx context.Context, confW *config.Wrapper) (storage.Store, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, err
		}

		return NewStore(ctx, conf)
	})
}

func NewStore(ctx context.Context, conf *Conf) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func runMigrations(db *sql.DB) error { _ = "STUB: not implemented"; return nil }

func upsertPolicy(ctx context.Context, tx *goqu.TxDatabase, p policy.Wrapper) error {
	_ = "STUB: not implemented"
	return nil
}

type Store struct {
	internal.DBStorage
	source *auditv1.PolicySource
}

func (s *Store) Driver() string { _ = "STUB: not implemented"; return "" }

func (s *Store) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }
