// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package stats

import (
	"os"
	"path/filepath"
	"testing"
)

// --- CollectStats ---

func TestCollectStats_CountsGoFiles(t *testing.T) {
	// Not parallel: uses os.Chdir.
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.go"), []byte("line 1\nline 2\nline 3\n"), 0644)
	os.WriteFile(filepath.Join(dir, "b.go"), []byte("line 1\nline 2\n"), 0644)
	os.WriteFile(filepath.Join(dir, "c_test.go"), []byte("line 1\nline 2\nline 3\nline 4\n"), 0644)
	os.WriteFile(filepath.Join(dir, "d_test.go"), []byte("line 1\n"), 0644)

	origDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(origDir) })

	deps := StatsDeps{BinaryDir: "bin", MagefilesDir: "magefiles"}
	rec, err := CollectStats(deps)
	if err != nil {
		t.Fatalf("CollectStats: %v", err)
	}
	if rec.GoProdLOC != 5 {
		t.Errorf("GoProdLOC = %d, want 5", rec.GoProdLOC)
	}
	if rec.GoTestLOC != 5 {
		t.Errorf("GoTestLOC = %d, want 5", rec.GoTestLOC)
	}
	if rec.GoLOC != 10 {
		t.Errorf("GoLOC = %d, want 10", rec.GoLOC)
	}
}

func TestCollectStats_SkipsVendorAndBinaryDir(t *testing.T) {
	// Not parallel: uses os.Chdir.
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "vendor"), 0755)
	os.WriteFile(filepath.Join(dir, "vendor", "pkg.go"), []byte("skip\nskip\n"), 0644)
	os.MkdirAll(filepath.Join(dir, "bin"), 0755)
	os.WriteFile(filepath.Join(dir, "bin", "main.go"), []byte("skip\nskip\n"), 0644)
	os.MkdirAll(filepath.Join(dir, "magefiles"), 0755)
	os.WriteFile(filepath.Join(dir, "magefiles", "build.go"), []byte("skip\nskip\n"), 0644)
	os.WriteFile(filepath.Join(dir, "real.go"), []byte("counted\n"), 0644)

	origDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(origDir) })

	deps := StatsDeps{BinaryDir: "bin", MagefilesDir: "magefiles"}
	rec, err := CollectStats(deps)
	if err != nil {
		t.Fatalf("CollectStats: %v", err)
	}
	if rec.GoProdLOC != 1 {
		t.Errorf("GoProdLOC = %d, want 1 (only real.go counted)", rec.GoProdLOC)
	}
	if rec.GoTestLOC != 0 {
		t.Errorf("GoTestLOC = %d, want 0", rec.GoTestLOC)
	}
}

// --- CountLines ---

func TestCountLines_MultipleLines(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "test.go")
	os.WriteFile(path, []byte("line 1\nline 2\nline 3\n"), 0644)

	got, err := CountLines(path)
	if err != nil {
		t.Fatalf("CountLines: %v", err)
	}
	if got != 3 {
		t.Errorf("CountLines = %d, want 3", got)
	}
}

func TestCountLines_EmptyFile(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.go")
	os.WriteFile(path, []byte(""), 0644)

	got, err := CountLines(path)
	if err != nil {
		t.Fatalf("CountLines: %v", err)
	}
	if got != 0 {
		t.Errorf("CountLines(empty) = %d, want 0", got)
	}
}

func TestCountLines_NoTrailingNewline(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "noeol.go")
	os.WriteFile(path, []byte("line 1\nline 2"), 0644)

	got, err := CountLines(path)
	if err != nil {
		t.Fatalf("CountLines: %v", err)
	}
	if got != 2 {
		t.Errorf("CountLines(no-eol) = %d, want 2", got)
	}
}

func TestCountLines_MissingFile(t *testing.T) {
	t.Parallel()
	_, err := CountLines("/nonexistent/file.go")
	if err == nil {
		t.Error("CountLines(missing) should return error")
	}
}

// --- CountWordsInFile ---

func TestCountWordsInFile_Basic(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "words.txt")
	os.WriteFile(path, []byte("hello world foo bar"), 0644)

	got, err := CountWordsInFile(path)
	if err != nil {
		t.Fatalf("CountWordsInFile: %v", err)
	}
	if got != 4 {
		t.Errorf("CountWordsInFile = %d, want 4", got)
	}
}

func TestCountWordsInFile_MultipleSpaces(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "spaces.txt")
	os.WriteFile(path, []byte("  hello   world  \n\n  foo  "), 0644)

	got, err := CountWordsInFile(path)
	if err != nil {
		t.Fatalf("CountWordsInFile: %v", err)
	}
	if got != 3 {
		t.Errorf("CountWordsInFile(multi-space) = %d, want 3", got)
	}
}

func TestCountWordsInFile_Empty(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.txt")
	os.WriteFile(path, []byte(""), 0644)

	got, err := CountWordsInFile(path)
	if err != nil {
		t.Fatalf("CountWordsInFile: %v", err)
	}
	if got != 0 {
		t.Errorf("CountWordsInFile(empty) = %d, want 0", got)
	}
}

func TestCountWordsInFile_MissingFile(t *testing.T) {
	t.Parallel()
	_, err := CountWordsInFile("/nonexistent/file.txt")
	if err == nil {
		t.Error("CountWordsInFile(missing) should return error")
	}
}
