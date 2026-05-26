// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package junit

import (
	"encoding/xml"

	"google.golang.org/protobuf/proto"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
)

const (
	skipTestCaseMessage  = "This test was skipped"
	skipTestSuiteMessage = "This test suite was skipped"
)

func Build(results *policyv1.TestResults, verbose bool) (*TestSuites, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processTestCases(s *policyv1.TestResults_Suite) ([]testCase, Summary, error) {
	_ = "STUB: not implemented"
	return nil, *new(Summary), nil
}

func renderValue(v proto.Message) string { _ = "STUB: not implemented"; return "" }

type TestSuites struct {
	XMLName xml.Name `xml:"testsuites"`
	Suites  []testSuite
	Summary
	Tests int `xml:"tests,attr"`
}

type testSuite struct {
	XMLName     xml.Name   `xml:"testsuite"`
	Description string     `xml:"description,attr,omitempty"`
	Name        string     `xml:"name,attr"`
	File        string     `xml:"file,attr"`
	Failure     *failure   `xml:"failure,omitempty"`
	Error       *testError `xml:"error,omitempty"`
	Skip        *skipped   `xml:"skipped,omitempty"`
	Properties  []property `xml:"properties>property,omitempty"`
	TestCases   []testCase `xml:"testCases,omitempty"`
	Summary
	Tests int `xml:"tests,attr"`
}

type testCase struct {
	XMLName    xml.Name   `xml:"testcase"`
	Skipped    *skipped   `xml:"skipped,omitempty"`
	Failure    *failure   `xml:"failure,omitempty"`
	Error      *testError `xml:"error,omitempty"`
	Success    *success   `xml:"success,omitempty"`
	File       string     `xml:"file,attr"`
	Classname  string     `xml:"classname,attr"`
	Name       string     `xml:"name,attr"`
	Properties []property `xml:"properties>property,omitempty"`
}

type testError struct {
	XMLName xml.Name `xml:"error"`
	Type    string   `xml:"type,attr,omitempty"`
	Value   string   `xml:",chardata"` //nolint:tagliatelle
}

type success struct {
	resultSuccess
	XMLName xml.Name `xml:"success"`
	Type    string   `xml:"type,attr,omitempty"`
}

type resultSuccess struct {
	Outputs  *[]output `xml:"outputs>output,omitempty"`
	Actual   string    `xml:"actual,omitempty"`
	Expected string    `xml:"expected,omitempty"`
}

type failure struct {
	resultFailed
	XMLName xml.Name `xml:"failure"`
	Type    string   `xml:"type,attr,omitempty"`
	Message string   `xml:"message,attr"`
}

type resultFailed struct {
	Outputs  *[]output `xml:"outputs>output,omitempty"`
	Actual   string    `xml:"actual,omitempty"`
	Expected string    `xml:"expected,omitempty"`
}

type output struct {
	XMLName  xml.Name    `xml:"output"`
	Src      string      `xml:"src,attr"`
	Expected outputValue `xml:"expected"`
	Actual   outputValue `xml:"actual"`
}

type outputValue struct {
	Value string `xml:",cdata"` //nolint:tagliatelle
}

type skipped struct {
	XMLName xml.Name `xml:"skipped"`
	Message string   `xml:"message,attr"`
}

type property struct {
	XMLName xml.Name `xml:"property"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"` //nolint:tagliatelle
}

type Summary struct {
	Errors   int `xml:"errors,attr"`
	Failures int `xml:"failures,attr"`
	Skipped  int `xml:"skipped,attr"`
}
