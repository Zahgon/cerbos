// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package auxdata

import (
	"context"
	"errors"
	"time"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"

	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
)

var (
	errNilLocalKeySet   = errors.New("nil local keyset")
	errNoKeySetToVerify = errors.New("cannot determine keyset to use for validating the JWT")
)

type jwtHelper struct {
	keySets        map[string]keySet
	verify         bool
	acceptableSkew time.Duration
}

func newJWTHelper(ctx context.Context, conf *JWTConf) *jwtHelper {
	_ = "STUB: not implemented"
	return nil
}

func newJWKCache(ctx context.Context, log *zap.Logger) *jwk.Cache {
	_ = "STUB: not implemented"
	return nil
}

// this should never happen; jwk.NewCache only returns an error if you pass it an httprc.Client that has already been started.

type jwkErrSink struct {
	log *zap.Logger
}

func (j jwkErrSink) Put(_ context.Context, err error) { _ = "STUB: not implemented"; return }

func (j *jwtHelper) extract(ctx context.Context, auxJWT *requestv1.AuxData_JWT) (map[string]*structpb.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *jwtHelper) parseOptions(ctx context.Context, auxJWT *requestv1.AuxData_JWT) (opts []jwt.ParseOption, _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if keyset ID is not provided and we only have one keyset configured, use that as the default.

// use the keyset specified in the request

func (j *jwtHelper) doExtract(ctx context.Context, auxJWT *requestv1.AuxData_JWT, parseOpts []jwt.ParseOption) (map[string]*structpb.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type keySet interface {
	keySet(context.Context) (jwk.Set, []any, error)
}

// remoteKeySet holds an auto-refreshing remote keyset.
type remoteKeySet struct {
	*jwk.Cache
	url     string
	options []any
}

func newRemoteKeySet(ctx context.Context, cache *jwk.Cache, src *RemoteSource, options []any) *remoteKeySet {
	_ = "STUB: not implemented"
	return nil
}

func (rks *remoteKeySet) keySet(ctx context.Context) (jwk.Set, []any, error) {
	_ = "STUB: not implemented"
	return *new(jwk.Set), nil, nil
}

// localKeySet represents a keyset defined manually through the configuration.
type localKeySet func(context.Context) (jwk.Set, []any, error)

func newLocalKeySet(src *LocalSource, options []any) localKeySet {
	_ = "STUB: not implemented"
	return *new(localKeySet)
}

func (lks localKeySet) keySet(ctx context.Context) (jwk.Set, []any, error) {
	_ = "STUB: not implemented"
	return *new(jwk.Set), nil, nil
}
