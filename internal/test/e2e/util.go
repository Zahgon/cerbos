// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

//go:build e2e

package e2e

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randSrc = rand.NewSource(time.Now().UnixNano())

// RandomStr generates a string containing random alphabetic characters of the given length
// https://stackoverflow.com/a/31832326/7364928
func RandomStr(n int) string {
	_ = "STUB: not implemented"

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	return ""
}
