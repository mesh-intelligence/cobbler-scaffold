// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package orchestrator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Build compiles the project binary. If MainPackage is empty, the
// target is skipped.
func (o *Orchestrator) Build() error {
	if o.cfg.Project.MainPackage == "" {
		logf("build: skipping (no main_package configured)")
		return nil
	}
	outPath := filepath.Join(o.cfg.Project.BinaryDir, o.cfg.Project.BinaryName)
	logf("build: go build -o %s %s", outPath, o.cfg.Project.MainPackage)
	if err := os.MkdirAll(o.cfg.Project.BinaryDir, 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}
	cmd := exec.Command(binGo, "build", "-o", outPath, o.cfg.Project.MainPackage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go build: %w", err)
	}
	logf("build: done")
	return nil
}

// BuildAll compiles all cmd/ sub-packages to BinaryDir when MainPackage is
// empty. It discovers every cmd/*/main.go package and builds each to
// bin/<name> using go build -o bin/<name> ./cmd/<name>/. If no cmd/
// directory exists the target is skipped. prd003 B1.1.
func (o *Orchestrator) BuildAll() error {
	if o.cfg.Project.MainPackage != "" {
		// Single-package project — delegate to Build.
		return o.Build()
	}

	pkgs, err := discoverCmdPackages(".")
	if err != nil {
		return fmt.Errorf("discovering cmd packages: %w", err)
	}
	if len(pkgs) == 0 {
		logf("build:all: no cmd/ packages found, skipping")
		return nil
	}

	if err := os.MkdirAll(o.cfg.Project.BinaryDir, 0o755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}

	for _, pkg := range pkgs {
		name := filepath.Base(pkg)
		outPath := filepath.Join(o.cfg.Project.BinaryDir, name)
		logf("build:all: go build -o %s %s", outPath, pkg)
		cmd := exec.Command(binGo, "build", "-o", outPath, pkg)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("go build %s: %w", pkg, err)
		}
	}

	logf("build:all: built %d package(s) to %s", len(pkgs), o.cfg.Project.BinaryDir)
	return nil
}

// discoverCmdPackages returns the import paths of all packages under cmd/
// that contain a main.go file, relative to root.
func discoverCmdPackages(root string) ([]string, error) {
	cmdDir := filepath.Join(root, "cmd")
	entries, err := os.ReadDir(cmdDir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading cmd/: %w", err)
	}

	var pkgs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		mainGo := filepath.Join(cmdDir, e.Name(), "main.go")
		if _, err := os.Stat(mainGo); err == nil {
			pkgs = append(pkgs, "./cmd/"+e.Name()+"/")
		}
	}
	return pkgs, nil
}

// Lint runs golangci-lint on the project.
func (o *Orchestrator) Lint() error {
	logf("lint: running golangci-lint")
	cmd := exec.Command(binLint, "run", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("golangci-lint: %w", err)
	}
	logf("lint: done")
	return nil
}

// Install runs go install for the main package. If MainPackage
// is empty, the target is skipped.
func (o *Orchestrator) Install() error {
	if o.cfg.Project.MainPackage == "" {
		logf("install: skipping (no main_package configured)")
		return nil
	}
	logf("install: go install %s", o.cfg.Project.MainPackage)
	cmd := exec.Command(binGo, "install", o.cfg.Project.MainPackage)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go install: %w", err)
	}
	logf("install: done")
	return nil
}

// Clean removes the build artifact directory.
func (o *Orchestrator) Clean() error {
	logf("clean: removing %s", o.cfg.Project.BinaryDir)
	if err := os.RemoveAll(o.cfg.Project.BinaryDir); err != nil {
		return fmt.Errorf("removing %s: %w", o.cfg.Project.BinaryDir, err)
	}
	logf("clean: done")
	return nil
}

// DumpMeasurePrompt assembles and prints the measure prompt to stdout.
func (o *Orchestrator) DumpMeasurePrompt() error {
	prompt, err := o.buildMeasurePrompt("", "[]", 1)
	if err != nil {
		return fmt.Errorf("building measure prompt: %w", err)
	}
	fmt.Print(prompt)
	return nil
}

// DumpStitchPrompt assembles and prints the stitch prompt to stdout.
// Uses a placeholder task so the template structure is visible.
func (o *Orchestrator) DumpStitchPrompt() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}
	prompt, err := o.buildStitchPrompt(stitchTask{
		worktreeDir: cwd,
		id:          "EXAMPLE-001",
		title:       "Example task",
		description: "Placeholder task description for prompt preview.",
		issueType:   "task",
	})
	if err != nil {
		return fmt.Errorf("building stitch prompt: %w", err)
	}
	fmt.Print(prompt)
	return nil
}

// ExtractCredentials reads Claude credentials from the macOS Keychain
// and writes them to SecretsDir/TokenFile.
func (o *Orchestrator) ExtractCredentials() error {
	outPath := filepath.Join(o.cfg.Claude.SecretsDir, o.cfg.EffectiveTokenFile())
	logf("credentials: extracting to %s", outPath)
	if err := os.MkdirAll(o.cfg.Claude.SecretsDir, 0o700); err != nil {
		return fmt.Errorf("creating secrets directory: %w", err)
	}
	out, err := exec.Command(binSecurity, "find-generic-password",
		"-s", "Claude Code-credentials", "-w").Output()
	if err != nil {
		return fmt.Errorf("extracting credentials from keychain: %w", err)
	}
	if err := os.WriteFile(outPath, out, 0o600); err != nil {
		return fmt.Errorf("writing credentials: %w", err)
	}
	logf("credentials: written to %s", outPath)
	return nil
}
