// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/config"
)

const Backend = "kafka"

const (
	AckNone   = "none"
	AckAll    = "all"
	AckLeader = "leader"

	HeaderKeyEncoding = "cerbos.audit.encoding"
	HeaderKeyKind     = "cerbos.audit.kind"

	CompressionNone   = "none"
	CompressionGzip   = "gzip"
	CompressionSnappy = "snappy"
	CompressionLZ4    = "lz4"
	CompressionZstd   = "zstd"
)

type Encoding string

const (
	EncodingJSON     Encoding = "json"
	EncodingProtobuf Encoding = "protobuf"
)

type Kind []byte

var (
	// reallocate once ahead of time to avoid allocations in the hot path.
	KindAccess   Kind = []byte(audit.KindAccess)
	KindDecision Kind = []byte(audit.KindDecision)
)

func init() {
	audit.RegisterBackend(Backend, func(ctx context.Context, confW *config.Wrapper, decisionFilter audit.DecisionLogEntryFilter) (audit.Log, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read kafka audit log configuration: %w", err)
		}

		return NewPublisher(ctx, conf, decisionFilter)
	})
}

type Client interface {
	Close()
	Flush(context.Context) error
	TryProduce(context.Context, *kgo.Record, func(*kgo.Record, error))
	ProduceSync(context.Context, ...*kgo.Record) kgo.ProduceResults
}

type Publisher struct {
	Client         Client
	decisionFilter audit.DecisionLogEntryFilter
	marshaller     recordMarshaller
	sync           bool
	closeTimeout   time.Duration
}

func NewPublisher(ctx context.Context, conf *Conf, decisionFilter audit.DecisionLogEntryFilter) (*Publisher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Publisher) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Publisher) Backend() string { _ = "STUB: not implemented"; return "" }

func (p *Publisher) Enabled() bool { _ = "STUB: not implemented"; return false }

func (p *Publisher) WriteAccessLogEntry(ctx context.Context, record audit.AccessLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Publisher) WriteDecisionLogEntry(ctx context.Context, record audit.DecisionLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Publisher) write(ctx context.Context, msg *kgo.Record) error {
	_ = "STUB: not implemented"
	return nil
}

// detach the context from the caller so the request can return
// without cancelling any async kafka operations

// TODO: Currently have to duplicate logWrapper as it does not support async audit publishing.

// Due to async nature of this callback, we need to pull the `kind` out of the header

func newMarshaller(enc Encoding) recordMarshaller {
	_ = "STUB: not implemented"
	return *new(recordMarshaller)
}

type recordMarshaller struct {
	encoding    Encoding
	encodingKey []byte
}

type auditEntry interface {
	proto.Message
	GetCallId() string
	MarshalVT() ([]byte, error)
}

func (m recordMarshaller) Marshal(entry auditEntry, kind Kind) (*kgo.Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
