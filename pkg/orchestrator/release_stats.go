// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

// release_stats.go delegates release statistics to the internal/stats
// sub-package.

package orchestrator

import (
	st "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/stats"
)

// releaseRow type alias for backward compatibility.
type releaseRow = st.ReleaseRow

// ReleaseStats prints a table of roadmap releases with per-release PRD and
// requirement counts.
func (o *Orchestrator) ReleaseStats() error {
	return st.PrintReleaseStats()
}

// buildReleaseRows delegates to the internal/stats package.
func buildReleaseRows() ([]releaseRow, error) {
	return st.BuildReleaseRows()
}

// releaseAllUCsDone delegates to the internal/stats package.
func releaseAllUCsDone(statuses []string) bool {
	return st.ReleaseAllUCsDone(statuses)
}

// releaseAnyUCDone delegates to the internal/stats package.
func releaseAnyUCDone(statuses []string) bool {
	return st.ReleaseAnyUCDone(statuses)
}
