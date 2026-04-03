// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

// stats.go delegates LOC and documentation word counting to the
// internal/stats sub-package.

package orchestrator

import (
	"github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/claude"
	ictx "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/context"
	gh "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/github"
	"github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/gitops"
	st "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/stats"
)

// StatsRecord holds collected LOC and documentation word counts.
type StatsRecord = st.StatsRecord

// Stats provides LOC, documentation, generator, release, run, and
// outcome statistics.
type Stats struct {
	cfg     Config
	logf    func(string, ...any)
	git     gitops.GitOps
	tracker gh.WorkTracker
}

// NewStats creates a Stats instance with explicit dependencies.
func NewStats(cfg Config, logf func(string, ...any), git gitops.GitOps, tracker gh.WorkTracker) *Stats {
	return &Stats{cfg: cfg, logf: logf, git: git, tracker: tracker}
}

// CollectStats gathers Go LOC and documentation word counts.
func (s *Stats) CollectStats() (StatsRecord, error) {
	return st.CollectStats(s.statsDeps())
}

// PrintStats prints Go lines of code and documentation word counts as YAML.
func (s *Stats) PrintStats() error {
	return st.PrintStats(s.statsDeps())
}

// statsDeps constructs the StatsDeps struct from stats state.
func (s *Stats) statsDeps() st.StatsDeps {
	return st.StatsDeps{
		BinaryDir:            s.cfg.Project.BinaryDir,
		MagefilesDir:         s.cfg.Project.MagefilesDir,
		ResolveStandardFiles: ictx.ResolveStandardFiles,
		ClassifyContextFile:  ictx.ClassifyContextFile,
	}
}

// listGenerationBranches returns git branches matching the generation prefix.
func (s *Stats) listGenerationBranches() []string {
	return s.git.ListBranches(s.cfg.Generation.Prefix+"*", ".")
}

// historyDir returns the path to the history directory.
func (s *Stats) historyDir() string {
	return claude.HistoryDir(s.cfg.Cobbler.Dir, s.cfg.Cobbler.HistoryDir)
}
