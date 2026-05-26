// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

const (
	FeedbackMsg  = "It looks like an unexpected error happened. Please use the link below to report it to Cerbos developers."
	baseURL      = "https://github.com/cerbos/cerbos/issues/new"
	FeedbackLink = baseURL + "/choose"
)

func GenerateFeedbackLink(header, version, commitSHA string, stack []byte) string {
	_ = "STUB: not implemented"
	return ""
}

// this should never happen
