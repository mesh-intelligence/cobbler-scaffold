// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

// Command orchestrator is a CLI for the mage-claude-orchestrator library.
// It exposes the same targets as the magefile as cobra subcommands so the
// library can be exercised without mage.
package main

import (
	"fmt"
	"os"

	"github.com/mesh-intelligence/mage-claude-orchestrator/pkg/orchestrator"
	"github.com/spf13/cobra"
)

var configFile string

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "orchestrator",
	Short: "mage-claude-orchestrator CLI",
	Long:  "CLI wrapper around the mage-claude-orchestrator library for direct invocation without mage.",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", orchestrator.DefaultConfigFile, "path to configuration YAML")

	rootCmd.AddCommand(
		initCmd,
		resetCmd,
		statsCmd,
		buildCmd,
		lintCmd,
		installCmd,
		cleanCmd,
		credentialsCmd,
		measureCmd,
		stitchCmd,
		generatorCmd,
		beadsCmd,
		cobblerCmd,
	)
}

func newOrch() (*orchestrator.Orchestrator, error) {
	cfg, err := orchestrator.LoadConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("loading config %s: %w", configFile, err)
	}
	return orchestrator.New(cfg), nil
}

func run(fn func(*orchestrator.Orchestrator) error) error {
	o, err := newOrch()
	if err != nil {
		return err
	}
	return fn(o)
}

// --- top-level commands ---

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project (beads)",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Init) },
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Full reset: cobbler, generator, beads",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).FullReset) },
}

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Print Go lines of code and documentation word counts",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Stats) },
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile the project binary",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Build) },
}

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Run golangci-lint",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Lint) },
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Run go install for the main package",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Install) },
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove build artifacts",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Clean) },
}

var credentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Extract Claude credentials from macOS Keychain",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).ExtractCredentials) },
}

// --- measure ---

var measureCmd = &cobra.Command{
	Use:   "measure",
	Short: "Assess project state and propose tasks via Claude",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Measure) },
}

var measurePromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Print the measure prompt that would be sent to Claude",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).MeasurePrompt) },
}

func init() {
	measureCmd.AddCommand(measurePromptCmd)
}

// --- stitch ---

var stitchCmd = &cobra.Command{
	Use:   "stitch",
	Short: "Pick ready tasks and invoke Claude to execute them",
	RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).Stitch) },
}

// --- generator ---

var generatorCmd = &cobra.Command{
	Use:     "generator",
	Aliases: []string{"gen"},
	Short:   "Code-generation trail lifecycle commands",
}

func init() {
	generatorCmd.AddCommand(
		&cobra.Command{
			Use:   "start",
			Short: "Begin a new generation trail",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).GeneratorStart) },
		},
		&cobra.Command{
			Use:   "run",
			Short: "Execute N cycles of measure + stitch within the current generation",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).GeneratorRun) },
		},
		&cobra.Command{
			Use:   "resume",
			Short: "Recover from an interrupted run and continue",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).GeneratorResume) },
		},
		&cobra.Command{
			Use:   "stop",
			Short: "Complete a generation trail and merge into main",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).GeneratorStop) },
		},
		&cobra.Command{
			Use:   "list",
			Short: "Show active branches and past generations",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).GeneratorList) },
		},
		&cobra.Command{
			Use:   "switch",
			Short: "Commit current work and check out another generation branch",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).GeneratorSwitch) },
		},
		&cobra.Command{
			Use:   "reset",
			Short: "Destroy generation branches, worktrees, and Go source directories",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).GeneratorReset) },
		},
	)
}

// --- beads ---

var beadsCmd = &cobra.Command{
	Use:   "beads",
	Short: "Issue-tracker lifecycle commands",
}

func init() {
	beadsCmd.AddCommand(
		&cobra.Command{
			Use:   "init",
			Short: "Initialize the beads issue tracker",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).BeadsInit) },
		},
		&cobra.Command{
			Use:   "reset",
			Short: "Clear beads issue history",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).BeadsReset) },
		},
	)
}

// --- cobbler ---

var cobblerCmd = &cobra.Command{
	Use:   "cobbler",
	Short: "Cobbler scratch directory commands",
}

func init() {
	cobblerCmd.AddCommand(
		&cobra.Command{
			Use:   "reset",
			Short: "Remove the cobbler scratch directory",
			RunE:  func(cmd *cobra.Command, args []string) error { return run((*orchestrator.Orchestrator).CobblerReset) },
		},
	)
}
