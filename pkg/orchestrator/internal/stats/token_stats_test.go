// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package stats

import "testing"

// --- SortedKeys ---

func TestSortedKeys_Empty(t *testing.T) {
	t.Parallel()
	got := SortedKeys(map[string]int{})
	if len(got) != 0 {
		t.Errorf("SortedKeys(empty) = %v, want []", got)
	}
}

func TestSortedKeys_Sorted(t *testing.T) {
	t.Parallel()
	got := SortedKeys(map[string]int{"c": 3, "a": 1, "b": 2})
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("SortedKeys len = %d, want %d", len(got), len(want))
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("SortedKeys[%d] = %q, want %q", i, got[i], w)
		}
	}
}

func TestSortedKeys_SingleKey(t *testing.T) {
	t.Parallel()
	got := SortedKeys(map[string]int{"only": 42})
	if len(got) != 1 || got[0] != "only" {
		t.Errorf("SortedKeys(single) = %v, want [only]", got)
	}
}
