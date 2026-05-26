// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package crosspath

import (
	"errors"
)

const noOfPartsInUNCPaths = 3

var ErrUnsupportedWin32Path = errors.New("unsupported Win32 path")

// Encoded is an encoded path.
//
// UNIX paths are encoded as-is,
// UNC absolute paths such as `\\host\share\path\to\dir` encoded as `/host/share/path/to/dir`.
// Win32 absolute paths such as `C:\path\to\dir` encoded as `/C:/path/to/dir`,
// Win32 relative paths such as `path\to\dir` encoded as `path/to/dir`.
type Encoded struct {
	value string
	kind  Kind
	win32 bool
	root  bool
}

// Encode encodes a UNIX or Win32 path as Encoded.
func Encode(path string) (Encoded, error) { _ = "STUB: not implemented"; return *new(Encoded), nil }

// The path has a value such as `\\host\share` which means it is root.
//nolint:mnd

// deny paths such as `D:.`, `D:foo`

// The path has a value such as `D:` or `D:/` which means it is root.

// Decode decodes an Encoded path in its original encoding.
func Decode(encoded Encoded) string { _ = "STUB: not implemented"; return "" }

// Add the leading `\` we have removed while encoding the path (`\host\share\dir` -> `\\host\share\dir`).

// Trim the leading `/` we have added while encoding the path (`\c:\path\to\dir` -> `c:\path\to\dir`).

// Base returns the last element of path.
func Base(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Dir returns all but the last element of path, typically the path's directory.
func Dir(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Ext returns the file name extension used by path.
func Ext(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Join joins any number of path elements into a single path.
func Join(paths ...string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Match reports whether name matches the shell file name pattern.
func Match(path, pattern string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Rel returns a relative path that is lexically equivalent to targetPath when
// joined to basePath with an intervening separator. That is,
// [Join](basePath, Rel(basePath, targPath)) is equivalent to targPath itself.
func Rel(basePath, targetPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// VolumeName returns leading volume name for Windows paths.
// Given a Windows path
//   - "C:\foo\bar" it returns "C:".
//   - "\\host\share\foo" it returns "\\host\share".
//
// Otherwise, it returns empty string.
func VolumeName(path string) string { _ = "STUB: not implemented"; return "" }

type Kind uint32

const (
	KindUnknown Kind = iota
	KindDrive
	KindUNC
)

func isUNCPath(path string) bool { _ = "STUB: not implemented"; return false }

func isDrivePath(path string) bool { _ = "STUB: not implemented"; return false }

func validDriveLetter(letter uint8) bool { _ = "STUB: not implemented"; return false }

func toSlash(path string) string { _ = "STUB: not implemented"; return "" }

func toBackslash(path string) string { _ = "STUB: not implemented"; return "" }
