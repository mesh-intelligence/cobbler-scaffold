// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package orchestrator

import (
	st "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/stats"
)

// RunStats prints aggregate statistics for a completed generation run.
// When name is empty, it lists available generations.
func (s *Stats) RunStats(name string) error {
	return st.PrintRunStats(name, s.runStatsDeps())
}

// CompareRunStats prints a side-by-side comparison of two generation runs.
func (s *Stats) CompareRunStats(name1, name2 string) error {
	return st.PrintCompareStats(name1, name2, s.runStatsDeps())
}

func (s *Stats) runStatsDeps() st.RunStatsDeps {
	return st.RunStatsDeps{
		Log: s.logf,
		ListTags: func(pattern string) []string {
			return s.git.ListTags(pattern, ".")
		},
		ShowFile: func(ref, path string) ([]byte, error) {
			return s.git.ShowFileContent(ref, path, ".")
		},
		GenerationPrefix: s.cfg.Generation.Prefix,
		CobblerDir:       s.cfg.Cobbler.Dir,
		HistorySubdir:    s.cfg.Cobbler.HistoryDir,
	}
}
