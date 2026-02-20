// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package orchestrator

import (
	"os"
	"path/filepath"
	"testing"
)

// --- extractID ---

func TestExtractID(t *testing.T) {
	cases := []struct {
		path string
		want string
	}{
		{"docs/specs/product-requirements/prd001-feature.yaml", "prd001-feature"},
		{"docs/specs/use-cases/rel01.0-uc001-init.yaml", "rel01.0-uc001-init"},
		{"docs/specs/test-suites/test-rel01.0.yaml", "test-rel01.0"},
		{"simple.yaml", "simple"},
	}
	for _, tc := range cases {
		if got := extractID(tc.path); got != tc.want {
			t.Errorf("extractID(%q) = %q, want %q", tc.path, got, tc.want)
		}
	}
}

// --- extractPRDsFromTouchpoints ---

func TestExtractPRDsFromTouchpoints(t *testing.T) {
	tps := []string{
		"T1: Calculator component (prd001-core R1, R2)",
		"T2: Parser subsystem (prd002-parser)",
		"T3: No PRD reference here",
	}
	got := extractPRDsFromTouchpoints(tps)
	want := map[string]bool{"prd001-core": true, "prd002-parser": true}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for _, id := range got {
		if !want[id] {
			t.Errorf("unexpected PRD ID %q", id)
		}
	}
}

func TestExtractPRDsFromTouchpoints_Empty(t *testing.T) {
	got := extractPRDsFromTouchpoints(nil)
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

func TestExtractPRDsFromTouchpoints_NoPRDs(t *testing.T) {
	tps := []string{"T1: Some component", "T2: Another component"}
	got := extractPRDsFromTouchpoints(tps)
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

// --- extractUseCaseIDsFromTraces ---

func TestExtractUseCaseIDsFromTraces(t *testing.T) {
	traces := []string{
		"rel01.0-uc001-init",
		"rel01.0-uc002-lifecycle",
		"prd001-core R4",
	}
	got := extractUseCaseIDsFromTraces(traces)
	if len(got) != 2 {
		t.Fatalf("got %v, want 2 use case IDs", got)
	}
	want := map[string]bool{"rel01.0-uc001-init": true, "rel01.0-uc002-lifecycle": true}
	for _, id := range got {
		if !want[id] {
			t.Errorf("unexpected use case ID %q", id)
		}
	}
}

func TestExtractUseCaseIDsFromTraces_Empty(t *testing.T) {
	got := extractUseCaseIDsFromTraces(nil)
	if len(got) != 0 {
		t.Errorf("got %v, want empty", got)
	}
}

// --- loadUseCase ---

func TestLoadUseCase_ParsesIDAndTouchpoints(t *testing.T) {
	content := `id: rel01.0-uc001-init
title: Initialization
touchpoints:
  - T1: Core component (prd001-core R1)
  - T2: Config subsystem
`
	dir := t.TempDir()
	path := filepath.Join(dir, "rel01.0-uc001-init.yaml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	uc, err := loadUseCase(path)
	if err != nil {
		t.Fatalf("loadUseCase: %v", err)
	}
	if uc.ID != "rel01.0-uc001-init" {
		t.Errorf("ID: got %q, want %q", uc.ID, "rel01.0-uc001-init")
	}
	if len(uc.Touchpoints) != 2 {
		t.Errorf("Touchpoints: got %d, want 2", len(uc.Touchpoints))
	}
}

func TestLoadUseCase_MissingFile(t *testing.T) {
	_, err := loadUseCase("/nonexistent/uc.yaml")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}

// --- loadTestSuite ---

func TestLoadTestSuite_ParsesIDAndTraces(t *testing.T) {
	content := `id: test-rel01.0
title: Release 01.0 Tests
release: rel01.0
traces:
  - rel01.0-uc001-init
  - rel01.0-uc002-lifecycle
test_cases:
  - name: Init smoke test
    inputs:
      command: mage init
    expected:
      exit_code: 0
`
	dir := t.TempDir()
	path := filepath.Join(dir, "test-rel01.0.yaml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	ts, err := loadTestSuite(path)
	if err != nil {
		t.Fatalf("loadTestSuite: %v", err)
	}
	if ts.ID != "test-rel01.0" {
		t.Errorf("ID: got %q, want %q", ts.ID, "test-rel01.0")
	}
	if len(ts.Traces) != 2 {
		t.Errorf("Traces: got %d, want 2", len(ts.Traces))
	}
	if ts.Traces[0] != "rel01.0-uc001-init" {
		t.Errorf("Traces[0]: got %q, want %q", ts.Traces[0], "rel01.0-uc001-init")
	}
}

func TestLoadTestSuite_MissingFile(t *testing.T) {
	_, err := loadTestSuite("/nonexistent/test.yaml")
	if err == nil {
		t.Error("expected error for missing file, got nil")
	}
}
