// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

const (
	newPackageTime = 1
	testTimesPath  = "test-times.json"
)

type testTime struct {
	Package string  `json:"package" xml:"name,attr"`
	Time    float64 `json:"time" xml:"time,attr"`
}

type testTimes []testTime

func (tts testTimes) Len() int { _ = "STUB: not implemented"; return 0 }

func (tts testTimes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (tts testTimes) Swap(i, j int) { _ = "STUB: not implemented"; return }

type testTimesByKind map[string]testTimes
