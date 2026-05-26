// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/xml"
)

type combineCmd struct {
	Kinds []string
	Total int
}

func (cmd *combineCmd) Run() error { _ = "STUB: not implemented"; return nil }

//nolint:mnd

type junitReport struct {
	XMLName   xml.Name   `xml:"testsuites"`
	TestTimes []testTime `xml:"testsuite"`
}

func readReport(kind string, index int) (junitReport, error) {
	_ = "STUB: not implemented"
	return *new(junitReport), nil
}
