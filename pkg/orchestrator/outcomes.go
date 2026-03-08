// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

// outcomes.go delegates outcome reporting to the internal/stats sub-package.

package orchestrator

import (
	st "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/stats"
)

// OutcomeRecord holds parsed outcome trailer data from a single task commit.
type OutcomeRecord = st.OutcomeRecord

// outcomeSep delimits commit blocks in the git log output used by Outcomes.
const outcomeSep = st.OutcomeSep

// Outcomes scans all git branches for commits that carry outcome trailers
// and prints a summary table to stdout.
func (o *Orchestrator) Outcomes() error {
	return st.PrintOutcomes(st.OutcomesDeps{
		Log:    logf,
		GitBin: binGit,
	})
}

// parseOutcomeRecords delegates to the internal/stats package.
func parseOutcomeRecords(logOutput string) []OutcomeRecord {
	return st.ParseOutcomeRecords(logOutput)
}

// parseOneOutcomeBlock delegates to the internal/stats package.
func parseOneOutcomeBlock(block string) *OutcomeRecord {
	return st.ParseOneOutcomeBlock(block)
}

// extractBranchFromRefs delegates to the internal/stats package.
func extractBranchFromRefs(refs string) string {
	return st.ExtractBranchFromRefs(refs)
}

// formatDuration delegates to the internal/stats package.
func formatDuration(seconds int) string {
	return st.FormatDuration(seconds)
}
