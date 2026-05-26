// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"
	"sync"
	"time"

	statev1 "github.com/cerbos/cerbos/api/genpb/cerbos/state/v1"
	telemetryv1 "github.com/cerbos/cerbos/api/genpb/cerbos/telemetry/v1"
	"github.com/cerbos/cerbos/internal/storage"
	analytics "github.com/rudderlabs/analytics-go"
	"github.com/spf13/afero"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const (
	doNotTrackEnvVar  = "DO_NOT_TRACK"
	noTelemetryEnvVar = "CERBOS_NO_TELEMETRY"

	eventBufferSize = 8
	stateFile       = "cerbos.telemetry.json"
)

var (
	WriteKey     string
	DataPlaneURL string

	reporter  Reporter = nopReporter{}
	startTime          = time.Now()
)

type Reporter interface {
	Report(*telemetryv1.Event) bool
	Intercept() Interceptors
	Stop() error
}

type nopReporter struct{}

func (nopReporter) Report(_ *telemetryv1.Event) bool { _ = "STUB: not implemented"; return false }

func (nopReporter) Intercept() Interceptors { _ = "STUB: not implemented"; return *new(Interceptors) }

func (nopReporter) Stop() error { _ = "STUB: not implemented"; return nil }

func Start(ctx context.Context, store storage.Store) { _ = "STUB: not implemented"; return }

func isEnabled(conf *Conf) bool { _ = "STUB: not implemented"; return false }

func disabledByEnvVar(name string) bool { _ = "STUB: not implemented"; return false }

// if the var is not defined, assume consent.

// err on the side of caution and assume no consent.

func startReporter(_ context.Context, conf *Conf, store storage.Store, logger *zap.Logger) *analyticsReporter {
	_ = "STUB: not implemented"
	return nil
}

func initStateFS(dir string) afero.Fs { _ = "STUB: not implemented"; return *new(afero.Fs) }

//nolint:mnd

func Stop() { _ = "STUB: not implemented"; return }

func Report(event *telemetryv1.Event) bool { _ = "STUB: not implemented"; return false }

func Intercept() Interceptors { _ = "STUB: not implemented"; return *new(Interceptors) }

type analyticsReporter struct {
	state          *statev1.TelemetryState
	fsys           afero.Fs
	store          storage.Store
	eventChan      chan *telemetryv1.Event
	client         analytics.Client
	logger         *zap.Logger
	shutdownChan   chan struct{}
	reportInterval time.Duration
	closeOnce      sync.Once
}

func newAnalyticsReporter(conf *Conf, store storage.Store, fsys afero.Fs, logger *zap.Logger) (*analyticsReporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAnalyticsReporterWithClient(client analytics.Client, conf *Conf, store storage.Store, fsys afero.Fs, logger *zap.Logger) *analyticsReporter {
	_ = "STUB: not implemented"
	return nil
}

func readState(fsys afero.Fs) *statev1.TelemetryState { _ = "STUB: not implemented"; return nil }

func newState() *statev1.TelemetryState { _ = "STUB: not implemented"; return nil }

func (r *analyticsReporter) start() {
	_ = "STUB: not implemented"

	// don't let a panic in this goroutine crash the whole app.
	return
}

func (r *analyticsReporter) Intercept() Interceptors {
	_ = "STUB: not implemented"
	return *new(Interceptors)
}

func (r *analyticsReporter) reportServerLaunch() { _ = "STUB: not implemented"; return }

func (r *analyticsReporter) Report(event *telemetryv1.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *analyticsReporter) reportAPIActivity(event *telemetryv1.Event_ApiActivity) {
	_ = "STUB: not implemented"
	return
}

func (r *analyticsReporter) Stop() error { _ = "STUB: not implemented"; return nil }

func (r *analyticsReporter) reportServerStop() { _ = "STUB: not implemented"; return }

func (r *analyticsReporter) writeState() error { _ = "STUB: not implemented"; return nil }

//nolint:mnd

func (r *analyticsReporter) send(kind string, event proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func mkProps(event proto.Message) (analytics.Properties, error) {
	_ = "STUB: not implemented"
	return *new(analytics.Properties), nil
}

type zapLogWrapper struct {
	logger *zap.SugaredLogger
}

func (zlw zapLogWrapper) Logf(fmt string, args ...any) { _ = "STUB: not implemented"; return }

func (zlw zapLogWrapper) Errorf(fmt string, args ...any) { _ = "STUB: not implemented"; return }
