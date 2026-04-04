// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

// generator_stats.go delegates generator statistics to the internal/stats
// sub-package.

package orchestrator

import (
	gh "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/github"
	st "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/stats"
)

// GeneratorStats prints a status report for the current generation run.
func (s *Stats) GeneratorStats() error {
	currentBranch, _ := s.git.CurrentBranch(".")
	return st.PrintGeneratorStats(st.GeneratorStatsDeps{
		Log:                    s.logf,
		ListGenerationBranches: s.listGenerationBranches,
		GenerationBranch:       s.cfg.Generation.Branch,
		CurrentBranch:          currentBranch,
		DetectGitHubRepo: func() (string, error) {
			return s.tracker.DetectGitHubRepo(".")
		},
		ListAllIssues: func(repo, generation string) ([]gh.CobblerIssue, error) {
			return s.tracker.ListAllCobblerIssues(repo, generation)
		},
		HistoryDir: s.historyDir(),
		CobblerDir: s.cfg.Cobbler.Dir,
		ReadBranchFile: func(branch, path string) ([]byte, error) {
			return s.git.ShowFileContent(branch, path, ".")
		},
	})
}
