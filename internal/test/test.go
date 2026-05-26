// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"io"
	"io/fs"
	"os"
	"testing"

	"google.golang.org/protobuf/proto"

	policyv1 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	"github.com/cerbos/cerbos/internal/observability/logging"
	"github.com/cerbos/cerbos/internal/policy"
	"github.com/cerbos/cerbos/internal/storage"
)

func init() {
	if logLevel := os.Getenv("CERBOS_TEST_LOG_LEVEL"); logLevel != "" {
		logging.InitLogging(context.Background(), logLevel, nil)
	}
}

func LoadPolicy(t *testing.T, path string) *policyv1.Policy { _ = "STUB: not implemented"; return nil }

func AddSchemasToStore(t *testing.T, dir string, ms storage.MutableStore) {
	_ = "STUB: not implemented"
	return
}

func ReadSchemaFromFile(t *testing.T, path string) []byte { _ = "STUB: not implemented"; return nil }

func ReadSchemaFromFS(t *testing.T, fsys fs.FS, path string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func mkReadCloser(t *testing.T, file string) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func PathToDir(tb testing.TB, dir string) string { _ = "STUB: not implemented"; return "" }

func DataFS() fs.FS { _ = "STUB: not implemented"; return *new(fs.FS) }

type Case struct {
	Want       map[string][]byte
	Name       string
	SourceFile string
	Input      []byte
}

// LoadTestCases loads groups of test files from the given path.
// Consider a directory containing the following set of files:
// |- test01.yaml
// |- test01.yaml.err
// |- test01.yaml.out
//
// The above files will be converted to a Case object as follows:
//
//	Case {
//	  Name: "test01",
//	  Input: <contents_of_test01.yaml>,
//	  Want: map[string][]byte{
//	    "err": <contents_of_test01.yaml.err>,
//	    "out": <contents_of_test01.yaml.out>,
//	  }
//	}.
func LoadTestCases(tb testing.TB, subDir string) []Case { _ = "STUB: not implemented"; return nil }

func readFileContents(tb testing.TB, filePath string) []byte { _ = "STUB: not implemented"; return nil }

func SkipIfGHActions(t *testing.T) { _ = "STUB: not implemented"; return }

func FindPolicyFiles(t *testing.T, dir string, callback func(string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func FilterPolicies[P *policyv1.Policy | policy.Wrapper](t *testing.T, policies []P, params storage.ListPolicyIDsParams) []P {
	_ = "STUB: not implemented"
	return nil
}

func WriteGoldenFile(t *testing.T, path string, contents proto.Message) {
	_ = "STUB: not implemented"
	return
}

//nolint:mnd,gosec
