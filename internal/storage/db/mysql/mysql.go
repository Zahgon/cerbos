// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build !js && !wasm

package mysql

import (
	"context"

	"github.com/doug-martin/goqu/v9"

	// Import the MySQL dialect.
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"

	auditv1 "github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	"github.com/cerbos/cerbos/internal/config"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
	"github.com/cerbos/cerbos/internal/storage/db/internal"
)

const (
	DriverName                 = "mysql"
	urlToSchemaDocs            = "https://docs.cerbos.dev/cerbos/latest/configuration/storage.html#mysql-schema"
	constraintViolationErrCode = 1062
)

var (
	_ storage.SourceStore  = (*Store)(nil)
	_ storage.MutableStore = (*Store)(nil)
	_ storage.Subscribable = (*Store)(nil)
)

func init() {
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

func buildDSN(conf *Conf) (string, error) { _ = "STUB: not implemented"; return "", nil }

func registerTLS(conf *Conf) error { _ = "STUB: not implemented"; return nil }

// Most MySQL versions have not caught up to modern TLS settings so we can't use utils.DefaultTLSConfig here.
//nolint:gosec

func registerServerPubKeys(conf *Conf) error { _ = "STUB: not implemented"; return nil }

func upsertPolicy(ctx context.Context, tx *goqu.TxDatabase, p policy.Wrapper) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if the existing policy name matches the name of the policy we are trying to insert.
// The reason for not doing an UPDATE WHERE and checking the number of affected rows is because MySQL
// returns 0 if the update did not change any of the columns as well.

// attempt update

type Store struct {
	internal.DBStorage
	source *auditv1.PolicySource
}

func (s *Store) Driver() string { _ = "STUB: not implemented"; return "" }

func (s *Store) Source() *auditv1.PolicySource { _ = "STUB: not implemented"; return nil }
