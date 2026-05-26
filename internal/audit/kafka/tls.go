// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package kafka

import (
	"context"
	"crypto/tls"
	"sync"
	"time"
)

func NewTLSConfig(ctx context.Context, reloadInterval time.Duration, insecureSkipVerify bool, caPath, certPath, keyPath string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G402

//nolint:gosec

func loadTLSCert(certPath, keyPath string) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type tlsReloader struct {
	cert           *tls.Certificate
	certPath       string
	keyPath        string
	mu             sync.RWMutex
	reloadInterval time.Duration
}

func newTLSReloader(ctx context.Context, reloadInterval time.Duration, certPath, keyPath string) (*tlsReloader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *tlsReloader) reload(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *tlsReloader) GetCertificateFunc() func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	_ = "STUB: not implemented"
	return nil
}
