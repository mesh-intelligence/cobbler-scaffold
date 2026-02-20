// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package orchestrator

import "testing"

// --- parseBranchList ---

func TestParseBranchList_StripsMarkers(t *testing.T) {
	input := "  main\n* current\n+ other\n"
	got := parseBranchList(input)
	want := []string{"main", "current", "other"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i, w := range want {
		if got[i] != w {
			t.Errorf("index %d: got %q, want %q", i, got[i], w)
		}
	}
}

func TestParseBranchList_EmptyInput(t *testing.T) {
	got := parseBranchList("")
	if len(got) != 0 {
		t.Errorf("got %v, want empty slice", got)
	}
}

func TestParseBranchList_SkipsBlankLines(t *testing.T) {
	got := parseBranchList("main\n\n  \nfeature\n")
	if len(got) != 2 || got[0] != "main" || got[1] != "feature" {
		t.Errorf("got %v, want [main feature]", got)
	}
}

func TestParseBranchList_GenerationPattern(t *testing.T) {
	input := "  generation-20260214.0\n  generation-20260215.1\n"
	got := parseBranchList(input)
	if len(got) != 2 {
		t.Fatalf("got %v, want 2 entries", got)
	}
	if got[0] != "generation-20260214.0" || got[1] != "generation-20260215.1" {
		t.Errorf("got %v", got)
	}
}

// --- parseDiffShortstat ---

func TestParseDiffShortstat_FullOutput(t *testing.T) {
	s := " 5 files changed, 100 insertions(+), 20 deletions(-)\n"
	ds := parseDiffShortstat(s)
	if ds.FilesChanged != 5 {
		t.Errorf("FilesChanged: got %d, want 5", ds.FilesChanged)
	}
	if ds.Insertions != 100 {
		t.Errorf("Insertions: got %d, want 100", ds.Insertions)
	}
	if ds.Deletions != 20 {
		t.Errorf("Deletions: got %d, want 20", ds.Deletions)
	}
}

func TestParseDiffShortstat_InsertionsOnly(t *testing.T) {
	s := " 3 files changed, 42 insertions(+)\n"
	ds := parseDiffShortstat(s)
	if ds.FilesChanged != 3 {
		t.Errorf("FilesChanged: got %d, want 3", ds.FilesChanged)
	}
	if ds.Insertions != 42 {
		t.Errorf("Insertions: got %d, want 42", ds.Insertions)
	}
	if ds.Deletions != 0 {
		t.Errorf("Deletions: got %d, want 0", ds.Deletions)
	}
}

func TestParseDiffShortstat_Empty(t *testing.T) {
	ds := parseDiffShortstat("")
	if ds.FilesChanged != 0 || ds.Insertions != 0 || ds.Deletions != 0 {
		t.Errorf("empty input: got %+v, want all zeros", ds)
	}
}

func TestParseDiffShortstat_SingleFile(t *testing.T) {
	s := " 1 file changed, 1 insertion(+), 1 deletion(-)\n"
	ds := parseDiffShortstat(s)
	if ds.FilesChanged != 1 {
		t.Errorf("FilesChanged: got %d, want 1", ds.FilesChanged)
	}
	if ds.Insertions != 1 {
		t.Errorf("Insertions: got %d, want 1", ds.Insertions)
	}
	if ds.Deletions != 1 {
		t.Errorf("Deletions: got %d, want 1", ds.Deletions)
	}
}
