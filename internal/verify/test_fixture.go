// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package verify

import (
	"io/fs"

	"google.golang.org/protobuf/proto"

	enginev1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
)

type Principals struct {
	LoadError error
	Fixtures  map[string]*enginev1.Principal
	Groups    map[string][]string
	FilePath  string
}

type Resources struct {
	LoadError error
	Fixtures  map[string]*enginev1.Resource
	Groups    map[string][]string
	FilePath  string
}

type AuxData struct {
	LoadError error
	Fixtures  map[string]*enginev1.AuxData
	FilePath  string
}

type TestFixture struct {
	Principals *Principals
	Resources  *Resources
	AuxData    *AuxData
}

const (
	principalsFileName = "principals"
	resourcesFileName  = "resources"
)

var auxDataFileNames = []string{"auxdata", "auxData", "aux_data"}

func LoadTestFixture(fsys fs.FS, path string, continueOnError bool) (tf *TestFixture, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadResources(fsys fs.FS, path string) (*Resources, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadPrincipals(fsys fs.FS, path string) (*Principals, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadAuxData(fsys fs.FS, path string) (*AuxData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadFixtureElement(fsys fs.FS, path string, pb proto.Message, validate func(fs.FS, string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (tf *TestFixture) lookupPrincipal(name string) (*enginev1.Principal, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tf *TestFixture) lookupPrincipalGroup(name string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tf *TestFixture) lookupResource(name string) (*enginev1.Resource, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tf *TestFixture) lookupResourceGroup(name string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (tf *TestFixture) lookupAuxData(name string) (*enginev1.AuxData, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
