// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package kafka

import (
	"time"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/twmb/franz-go/pkg/kgo"
)

const confKey = audit.ConfKey + ".kafka"

const (
	defaultAcknowledgement    = AckAll
	defaultEncoding           = EncodingJSON
	defaultCloseTimeout       = 30 * time.Second
	defaultClientID           = "cerbos"
	defaultMaxBufferedRecords = 250
)

type Authentication struct {
	TLS *TLS `yaml:"tls"`
}

type TLS struct {
	// CAPath is the path to the CA certificate.
	CAPath string `yaml:"caPath" conf:"required,example=/path/to/ca.crt"`
	// CertPath is the path to the client certificate.
	CertPath string `yaml:"certPath" conf:",example=/path/to/tls.cert"`
	// KeyPath is the path to the client key.
	KeyPath string `yaml:"keyPath" conf:",example=/path/to/tls.key"`
	// ReloadInterval is the interval at which the TLS certificates are reloaded. The default is 0 (no reload).
	ReloadInterval time.Duration `yaml:"reloadInterval" conf:",example=5m"`
	// InsecureSkipVerify controls whether the server's certificate chain and host name are verified. Default is false.
	InsecureSkipVerify bool `yaml:"insecureSkipVerify" conf:",example=true"`
}

// Conf is optional configuration for kafka Audit.
type Conf struct {
	// Ack mode for producing messages. Valid values are "none", "leader" or "all" (default). Idempotency is disabled when mode is not "all".
	Ack string `yaml:"ack" conf:",example=all"`
	// Authentication
	Authentication Authentication `yaml:"authentication"`
	// Topic to write audit entries to.
	Topic string `yaml:"topic" conf:"required,example=cerbos.audit.log"`
	// Encoding format. Valid values are "json" (default) or "protobuf".
	Encoding Encoding `yaml:"encoding" conf:",example=json"`
	// ClientID reported in Kafka connections.
	ClientID string `yaml:"clientID" conf:",example=cerbos"`
	// Brokers list to seed the Kafka client.
	Brokers []string `yaml:"brokers" conf:"required,example=['localhost:9092']"`
	// Compression sets the compression algorithm to use in order of priority. Valid values are "none", "gzip", "snappy","lz4", "zstd". Default is ["snappy", "none"].
	Compression []string `yaml:"compression" conf:",example=['snappy', 'none']"`
	// CloseTimeout sets how long when closing the client to wait for any remaining messages to be flushed.
	CloseTimeout time.Duration `yaml:"closeTimeout" conf:",example=30s"`
	// MaxBufferedRecords sets the maximum number of records the client should buffer in memory in async mode.
	MaxBufferedRecords int `yaml:"maxBufferedRecords" conf:",example=1000"`
	// ProduceSync forces the client to produce messages to Kafka synchronously. This can have a significant impact on performance.
	ProduceSync bool `yaml:"produceSync" conf:",example=false"`
}

func (c *Conf) Key() string { _ = "STUB: not implemented"; return "" }

func (c *Conf) SetDefaults() { _ = "STUB: not implemented"; return }

func (c *Conf) Validate() error { _ = "STUB: not implemented"; return nil }

func formatAck(ack string) (kgo.Acks, error) { _ = "STUB: not implemented"; return *new(kgo.Acks), nil }

func formatCompression(compression []string) ([]kgo.CompressionCodec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
