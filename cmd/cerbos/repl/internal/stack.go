// Copyright 2021-2026 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package internal

type runeStack []rune

func (s *runeStack) Push(r rune) { _ = "STUB: not implemented"; return }

func (s *runeStack) Pop() (rune, bool) { _ = "STUB: not implemented"; return 0, false }

func (s *runeStack) Peek() (rune, bool) { _ = "STUB: not implemented"; return 0, false }

func (s *runeStack) IsEmpty() bool { _ = "STUB: not implemented"; return false }
