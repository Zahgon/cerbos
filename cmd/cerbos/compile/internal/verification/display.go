// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verification

import (
	"github.com/pterm/pterm"
	"google.golang.org/protobuf/proto"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/cmd/cerbos/compile/internal/flagset"
	"github.com/cerbos/cerbos/cmd/cerbos/compile/internal/verification/internal/traces"
	"github.com/cerbos/cerbos/internal/outputcolor"
	"github.com/cerbos/cerbos/internal/printer"
	"github.com/cerbos/cerbos/internal/printer/colored"
)

const (
	suiteLevel         = 0
	testCaseLevel      = 1
	principalLevel     = 2
	resourceLevel      = 3
	actionLevel        = 4
	resultLevel        = 5
	outputSrcLevel     = 6
	outputErrKindLevel = 7
	outputErrValLevel  = 8

	listIndent = 2
)

var (
	labelColors = map[policyv1.TestResults_Result]func(...any) string{
		policyv1.TestResults_RESULT_PASSED:  colored.PassedTest,
		policyv1.TestResults_RESULT_SKIPPED: colored.SkippedTest,
		policyv1.TestResults_RESULT_FAILED:  colored.FailedTest,
		policyv1.TestResults_RESULT_ERRORED: colored.ErroredTest,
	}

	labels = map[policyv1.TestResults_Result]string{
		policyv1.TestResults_RESULT_PASSED:  "OK",
		policyv1.TestResults_RESULT_SKIPPED: "SKIPPED",
		policyv1.TestResults_RESULT_FAILED:  "FAILED",
		policyv1.TestResults_RESULT_ERRORED: "ERROR",
	}
)

func Display(p *printer.Printer, results *policyv1.TestResults, output flagset.VerificationOutputFormat, verbose bool, colorLevel outputcolor.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func displayJUnit(p *printer.Printer, results *policyv1.TestResults, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

func displayTree(p *printer.Printer, tp pterm.TreePrinter, results *policyv1.TestResults, verbose bool) error {
	_ = "STUB: not implemented"
	return nil
}

type testOutput struct {
	traces  traces.Map
	tree    pterm.LeveledList
	verbose bool
}

func buildTestOutput(results *policyv1.TestResults, verbose bool) *testOutput {
	_ = "STUB: not implemented"
	return nil
}

func (o *testOutput) addSuite(suite *policyv1.TestResults_Suite) { _ = "STUB: not implemented"; return }

func (o *testOutput) shouldAddSuite(suite *policyv1.TestResults_Suite) bool {
	_ = "STUB: not implemented"
	return false
}

func (o *testOutput) allTestsSkippedDueToFilter(suite *policyv1.TestResults_Suite) bool {
	_ = "STUB: not implemented"
	return false
}

func (o *testOutput) addTestCase(suite *policyv1.TestResults_Suite, testCase *policyv1.TestResults_TestCase) {
	_ = "STUB: not implemented"
	return
}

func (o *testOutput) shouldAddTestCase(testCase *policyv1.TestResults_TestCase) bool {
	_ = "STUB: not implemented"
	return false
}

func (o *testOutput) addPrincipal(suite *policyv1.TestResults_Suite, principal *policyv1.TestResults_Principal) {
	_ = "STUB: not implemented"
	return
}

func (o *testOutput) shouldAddPrincipal(principal *policyv1.TestResults_Principal) bool {
	_ = "STUB: not implemented"
	return false
}

func (o *testOutput) addResource(suite *policyv1.TestResults_Suite, principal *policyv1.TestResults_Principal, resource *policyv1.TestResults_Resource) {
	_ = "STUB: not implemented"
	return
}

func (o *testOutput) shouldAddResource(resource *policyv1.TestResults_Resource) bool {
	_ = "STUB: not implemented"
	return false
}

func (o *testOutput) addAction(suite *policyv1.TestResults_Suite, principal *policyv1.TestResults_Principal, resource *policyv1.TestResults_Resource, action *policyv1.TestResults_Action) {
	_ = "STUB: not implemented"
	return
}

func (o *testOutput) shouldAddAction(action *policyv1.TestResults_Action) bool {
	_ = "STUB: not implemented"
	return false
}

func (o *testOutput) appendNode(level int, text string) { _ = "STUB: not implemented"; return }

func resultLabel(result policyv1.TestResults_Result) string { _ = "STUB: not implemented"; return "" }

func tallyLabel(tally *policyv1.TestResults_Tally) string { _ = "STUB: not implemented"; return "" }

func singleLineJSON(m proto.Message) string { _ = "STUB: not implemented"; return "" }
