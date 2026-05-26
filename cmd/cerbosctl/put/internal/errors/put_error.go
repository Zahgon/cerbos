// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package errors

func NewPutError(path, message string) *PutError { _ = "STUB: not implemented"; return nil }

type PutError struct {
	Path    string
	Message string
}

func (pe *PutError) Error() string { _ = "STUB: not implemented"; return "" }
