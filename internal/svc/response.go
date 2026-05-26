// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package svc

import (
	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	requestv1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	responsev1 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
)

type checkResourceSetResponseBuilder struct {
	*responsev1.CheckResourceSetResponse
	includeMeta bool
}

func newCheckResourceSetResponseBuilder(req *requestv1.CheckResourceSetRequest) *checkResourceSetResponseBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (resp *checkResourceSetResponseBuilder) addResult(resourceKey string, result *enginev1.CheckOutput) {
	_ = "STUB: not implemented"
	return
}

func (resp *checkResourceSetResponseBuilder) addResultMeta(resourceKey string, result *enginev1.CheckOutput) {
	_ = "STUB: not implemented"
	return
}

func (resp *checkResourceSetResponseBuilder) build() *responsev1.CheckResourceSetResponse {
	_ = "STUB: not implemented"
	return nil
}
