// Copyright (c) 2017, Fatih Arslan

//go:build confdocs
// +build confdocs

package structtag

import (
	"errors"
)

var (
	errTagSyntax      = errors.New("bad syntax for struct tag pair")
	errTagKeySyntax   = errors.New("bad syntax for struct tag key")
	errTagValueSyntax = errors.New("bad syntax for struct tag value")
	errTagNotExist    = errors.New("tag does not exist")
)

type Tags struct {
	tags []*Tag
}

type Tag struct {
	Key     string
	Name    string
	Options []string
}

func Parse(tag string) (*Tags, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *Tags) Get(key string) (*Tag, error) { _ = "STUB: not implemented"; return nil, nil }
