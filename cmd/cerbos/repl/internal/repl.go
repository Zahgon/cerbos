// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"context"
	_ "embed"
	"errors"
	"io"
	"reflect"
	"regexp"

	"github.com/alecthomas/participle/v2"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/decls"
	celtypes "github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter"
	"github.com/peterh/liner"
	"github.com/pterm/pterm"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	runtimev1 "github.com/cerbos/cerbos/api/genpb/cerbos/runtime/v1"
	"github.com/cerbos/cerbos/internal/conditions/types"
	"github.com/cerbos/cerbos/internal/outputcolor"
	"github.com/cerbos/cerbos/internal/printer"
)

var (
	//go:embed banner.txt
	banner string
	//go:embed help.txt
	helpText string

	errExit   = errors.New("exit")
	errSilent = errors.New("") // returned when an error has occurred but feedback has already been provided to the user

	listType = reflect.TypeFor[[]any]()
	mapType  = reflect.TypeFor[map[string]any]()

	oppositeChars = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
)

const (
	commentPrefix   = '#'
	directivePrefix = ':'
	prompt          = "-> "
	rulePrefix      = "#"
	secondaryPrompt = "> "
	yamlIndent      = 2
)

type policyHolder struct {
	key       string
	variables map[string]string
	rules     []proto.Message
}

type REPL struct {
	output       Output
	vars         variables
	decls        map[string]*decls.VariableDecl
	reader       *liner.State
	parser       *participle.Parser[REPLDirective]
	toRefVal     func(any) ref.Val
	policy       *policyHolder
	varC         types.VariablesMap
	varV         types.VariablesMap
	constExports map[string]map[string]*structpb.Value
	varExports   map[string]map[string]string
}

func NewREPL(reader *liner.State, output Output) (*REPL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *REPL) Loop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *REPL) readInput() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *REPL) handleInput(ctx context.Context, input string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *REPL) processDirective(ctx context.Context, line string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *REPL) reset() error { _ = "STUB: not implemented"; return nil }

func (r *REPL) help() error { _ = "STUB: not implemented"; return nil }

func (r *REPL) showVars() error { _ = "STUB: not implemented"; return nil }

func (r *REPL) showRules() error { _ = "STUB: not implemented"; return nil }

func (r *REPL) setSpecialVar(ctx context.Context, name, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *REPL) addToVarC(c map[string]any) { _ = "STUB: not implemented"; return }

func (r *REPL) addToVarV(v map[string]any) { _ = "STUB: not implemented"; return }

func (r *REPL) processExpr(ctx context.Context, name, expr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *REPL) evalExpr(ctx context.Context, expr string) (ref.Val, *celtypes.Type, error) {
	_ = "STUB: not implemented"
	return *new(ref.Val), nil, nil
}

func (r *REPL) loadPolicy(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

func (r *REPL) printLoadError(err error) { _ = "STUB: not implemented"; return }

//nolint:errorlint

func (r *REPL) mergeConstantDefinitions(policyKey string, policyConstants *policyv1.Constants) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeConstantDefinitions(merged map[string]any, sources map[string][]string, values map[string]*structpb.Value, source string) {
	_ = "STUB: not implemented"
	return
}

func (r *REPL) mergeVariableDefinitions(policyKey string, policyVariables *policyv1.Variables, deprecatedTopLevel map[string]string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeVariableDefinitions(merged map[string]string, sources map[string][]string, values map[string]string, source string) {
	_ = "STUB: not implemented"
	return
}

func (r *REPL) evalPolicyVariables(ctx context.Context) { _ = "STUB: not implemented"; return }

func (r *REPL) execRule(ctx context.Context, id int) error { _ = "STUB: not implemented"; return nil }

func (r *REPL) evalCondition(ctx context.Context, id int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *REPL) doEvalCondition(ctx context.Context, condition *runtimev1.Condition) *eval {
	_ = "STUB: not implemented"
	return nil
}

func (r *REPL) mkEnv() (*cel.Env, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *REPL) Complete(line string) []string { _ = "STUB: not implemented"; return nil }

// This isn't ideal, but the directives are baked into tags
// making this trickier to implement neatly

var directives = []string{
	":help",
	":let",
	":load",
	":vars",
	":rules",
	":exec",
	":reset",
	":quit",
}

func (r *REPL) completeCmd(prefix string) (c []string) { _ = "STUB: not implemented"; return nil }

var (
	skippableDirs  = regexp.MustCompile(`^(_schemas|testdata|derived_roles)$`)
	skippableFiles = regexp.MustCompile(`(^\.)|(_test\.(yaml|yml|json)$)`)
	matchingFiles  = regexp.MustCompile(`\.(yaml|yml|json)$`)
)

func (r *REPL) completeFile(prefix string) (c []string) { _ = "STUB: not implemented"; return nil }

// This is subtle, but the first path may be a complete directory the user explicitly
// passed, so should not be skipped. This is also needed for the tests to pass (since
// we skip testdata

// TODO(tcm): this should probably limit the depth of the search

func (r *REPL) completeVar(prefix string) (c []string) { _ = "STUB: not implemented"; return nil }

func (r *REPL) completeRule(prefix string) (c []string) { _ = "STUB: not implemented"; return nil }

func isTerminated(line string, stack *runeStack) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// variables is a type that provides the interpreter.Activation interface.
type variables map[string]ref.Val

func (v variables) ResolveName(name string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (v variables) Parent() interpreter.Activation {
	_ = "STUB: not implemented"
	return *new(interpreter.Activation)
}

type Output interface {
	Print(string, ...any)
	Println(...any)
	PrintResult(string, ref.Val)
	PrintRule(int, proto.Message) error
	PrintJSON(any)
	PrintYAML(proto.Message, int)
	PrintTree(pterm.LeveledList) error
	PrintErr(string, error)
}

type PrinterOutput struct {
	*printer.Printer
	treePrinter *pterm.TreePrinter
	level       outputcolor.Level
}

func NewPrinterOutput(stdout, stderr io.Writer) *PrinterOutput {
	_ = "STUB: not implemented"
	return nil
}

func (po *PrinterOutput) Print(format string, args ...any) { _ = "STUB: not implemented"; return }

func (po *PrinterOutput) PrintRule(id int, rule proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (po *PrinterOutput) PrintResult(name string, value ref.Val) { _ = "STUB: not implemented"; return }

func (po *PrinterOutput) PrintJSON(obj any) { _ = "STUB: not implemented"; return }

func (po *PrinterOutput) PrintYAML(obj proto.Message, indent int) {
	_ = "STUB: not implemented"
	return
}

func (po *PrinterOutput) PrintTree(tree pterm.LeveledList) error {
	_ = "STUB: not implemented"
	return nil
}

func (po *PrinterOutput) PrintErr(msg string, err error) { _ = "STUB: not implemented"; return }
