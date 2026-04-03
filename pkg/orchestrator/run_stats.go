// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package orchestrator

import (
	st "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/stats"
)

// RunStats prints aggregate statistics for a completed generation run.
// When name is empty, it lists available generations.
func (o *Orchestrator) RunStats(name string) error {
	return st.PrintRunStats(name, st.RunStatsDeps{
		Log: logf,
		ListTags: func(pattern string) []string {
			return defaultGitOps.ListTags(pattern, ".")
		},
		ShowFile: func(ref, path string) ([]byte, error) {
			return defaultGitOps.ShowFileContent(ref, path, ".")
		},
		GenerationPrefix: o.cfg.Generation.Prefix,
		CobblerDir:       o.cfg.Cobbler.Dir,
		HistorySubdir:    o.cfg.Cobbler.HistoryDir,
	})
}
