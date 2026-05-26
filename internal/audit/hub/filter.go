// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	logsv1 "github.com/cerbos/cloud-api/genpb/cerbos/cloud/logs/v1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	peerPart             = "peer"
	metadataPart         = "metadata"
	checkResourcesPrefix = "decisionLogEntry.checkResources"
	planResourcesPrefix  = "decisionLogEntry.planResources"
)

var entryKindPrefixes = []string{"accessLogEntry", "decisionLogEntry"}

type tokenType int8

const (
	tokenUnknown tokenType = iota
	tokenAccessor
	tokenIndex
	tokenWildcard
)

type Token struct {
	val      any
	children map[string]*Token
	typ      tokenType
}

func (t *Token) key() string { _ = "STUB: not implemented"; return "" }

//nolint:forcetypeassert

//nolint:forcetypeassert

type AuditLogFilter struct {
	astRoot *Token
}

func NewAuditLogFilter(conf MaskConf) (*AuditLogFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewOversizedLogFilter() (*AuditLogFilter, error) { _ = "STUB: not implemented"; return nil, nil }

func parseJSONPathExprs(conf MaskConf) (ast *Token, outErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type state int8

const (
	stateUnknown state = iota
	statePlainAccessor
	stateParenOpen
	stateWildcard
	stateNumberOpen
	stateStringOpen
	stateStringClosed
	stateClosed
)

type tokenBuilder struct {
	curToken *Token
	buf      string
	size     int
	t        tokenType
	s        state
}

func (tb *tokenBuilder) WriteRune(r rune) error { _ = "STUB: not implemented"; return nil }

func (tb *tokenBuilder) Flush() error { _ = "STUB: not implemented"; return nil }

//nolint:exhaustive

func Tokenize(root *Token, path string) error { _ = "STUB: not implemented"; return nil }

// handle and validate token boundaries

func (f *AuditLogFilter) Filter(entry *logsv1.IngestBatch_Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// We support a subset of JSONPath operations, as follows:
//
// - dot notation: `foo.bar.baz`
// - or bracket-notation: `['foo']['bar']['baz]`
// - or combinations thereof
//
// `bar` or `baz` above can be map keys, nested messages or structs.
//
// We support list indexing with Ints or wildcards:
// - foo.bar[0]
// - foo.bar[*]
//
// Wildcards can also operate on member names as a match-all. E.g `foo[*].baz`
// will match both `baz` values in the pseudo-object below:
//
//	{
//	  'foo': {
//	    'pow': {
//	        'baz',
//	    },
//	    'bosh': {
//	        'baz',
//	    },
//	  }
//	}
func visit(t *Token, m protoreflect.Message) { _ = "STUB: not implemented"; return }

//nolint:forcetypeassert

//nolint:exhaustive

//nolint:forcetypeassert

// apparently retrieving typed values each iteration causes no slow-down

// For array indexes, reach ahead to the next token.

//nolint:exhaustive

//nolint:forcetypeassert

func visitStructpb(t *Token, v *structpb.Value) { _ = "STUB: not implemented"; return }

//nolint:exhaustive
