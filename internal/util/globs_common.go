// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"github.com/cerbos/cerbos/internal/cache"
	"github.com/gobwas/glob"
)

var globs = &globCache{cache: cache.New[string, glob.Glob]("glob", 1024)} //nolint:mnd

type globCache struct {
	cache *cache.Cache[string, glob.Glob]
}

func (gc *globCache) matches(globExpr, val string) bool { _ = "STUB: not implemented"; return false }

func (gc *globCache) getOrCompile(globExpr string) glob.Glob {
	_ = "STUB: not implemented"
	return *new(glob.Glob)
}

// MatchesGlob returns true if the given glob expression matches the given string.
func MatchesGlob(g, val string) bool { _ = "STUB: not implemented"; return false }

// FilterGlob returns the set of values that match the given glob.
func FilterGlob(g string, values []string) []string { _ = "STUB: not implemented"; return nil }

// FilterGlobNotMatches returns the set of values that do not match the given glob.
func FilterGlobNotMatches(g string, values []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func fixGlob(g string) string {
	_ = "STUB: not implemented"
	// for backward compatibility, consider single * as **
	return ""
}

// GetOrCompileGlob returns a compiled glob for the given expression, using the global cache.
// Returns nil if the glob expression is invalid.
func GetOrCompileGlob(globExpr string) glob.Glob { _ = "STUB: not implemented"; return *new(glob.Glob) }
