// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package file

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/cerbos/cerbos/internal/audit"
	"github.com/cerbos/cerbos/internal/config"
)

const Backend = "file"

func init() {
	audit.RegisterBackend(Backend, func(_ context.Context, confW *config.Wrapper, decisionFilter audit.DecisionLogEntryFilter) (audit.Log, error) {
		conf := new(Conf)
		if err := confW.GetSection(conf); err != nil {
			return nil, fmt.Errorf("failed to read local audit log configuration: %w", err)
		}

		return NewLog(conf, decisionFilter)
	})
}

type Log struct {
	accessLog      *zap.Logger
	decisionLog    *zap.Logger
	decisionFilter audit.DecisionLogEntryFilter
}

func NewLog(conf *Conf, decisionFilter audit.DecisionLogEntryFilter) (*Log, error) {
	_ = "STUB: not implemented"
	// remove level, time and message because they are not useful in this context
	return nil, nil
}

func (l *Log) Backend() string { _ = "STUB: not implemented"; return "" }

func (l *Log) Enabled() bool { _ = "STUB: not implemented"; return false }

func (l *Log) WriteAccessLogEntry(_ context.Context, record audit.AccessLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Log) WriteDecisionLogEntry(_ context.Context, record audit.DecisionLogEntryMaker) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Log) Close() error { _ = "STUB: not implemented"; return nil }

type protoMsg struct {
	msg proto.Message
}

func (pm protoMsg) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeSingular(enc zapcore.ObjectEncoder, fieldName string, fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	_ = "STUB: not implemented"
	return
}

// output readbale timestamps and values

// do nothing

type protoMap struct {
	m       protoreflect.Map
	valueFD protoreflect.FieldDescriptor
}

func (pm protoMap) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type protoList struct {
	l       protoreflect.List
	valueFD protoreflect.FieldDescriptor
}

func (pl protoList) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

type syncErrIgnorer struct {
	zapcore.WriteSyncer
}

func (s syncErrIgnorer) Sync() error {
	_ = "STUB: not implemented"
	// https://github.com/uber-go/zap/issues/328
	return nil
}
