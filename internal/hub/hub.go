// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package hub

import (
	"sync"

	"github.com/cerbos/cloud-api/hub"
)

var Get = sync.OnceValues(getInstance)

func getInstance() (*hub.Hub, error) { _ = "STUB: not implemented"; return nil, nil }
