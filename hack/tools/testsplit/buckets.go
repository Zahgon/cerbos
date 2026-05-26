// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

type testBucket struct {
	Packages  []string
	TotalTime float64
}

func (tb *testBucket) Add(time testTime) { _ = "STUB: not implemented"; return }

type testBuckets []testBucket

func (tbs testBuckets) LeastFull() *testBucket { _ = "STUB: not implemented"; return nil }
