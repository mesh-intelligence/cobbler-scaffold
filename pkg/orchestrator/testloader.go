// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

// testloader.go re-exports comparison types from the internal/compare
// sub-package for backward compatibility within the orchestrator package.
// prd: prd004-differential-comparison R4, R5, R6

package orchestrator

import (
	"github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/compare"
)

// CompareTestCase is re-exported from the internal compare package.
type CompareTestCase = compare.CompareTestCase

// CompareExpected is re-exported from the internal compare package.
type CompareExpected = compare.CompareExpected

// TestResult is re-exported from the internal compare package.
type TestResult = compare.TestResult

// BinaryResolver is re-exported from the internal compare package.
type BinaryResolver = compare.BinaryResolver

// LoadCompareTestCases delegates to the internal compare package.
func LoadCompareTestCases(specsDir string) ([]CompareTestCase, error) {
	return compare.LoadCompareTestCases(specsDir)
}

// FilterByUtility delegates to the internal compare package.
func FilterByUtility(cases []CompareTestCase, utility string) []CompareTestCase {
	return compare.FilterByUtility(cases, utility)
}

// CompareUtility delegates to the internal compare package.
func CompareUtility(binA, binB string, cases []CompareTestCase) []TestResult {
	return compare.CompareUtility(binA, binB, cases)
}

// FormatResults delegates to the internal compare package.
func FormatResults(results []TestResult) string {
	return compare.FormatResults(results)
}

// ResolverFromArg delegates to the internal compare package.
func ResolverFromArg(arg string) BinaryResolver {
	deps := compare.Deps{
		Log:            logf,
		GitBin:         binGit,
		GoBin:          binGo,
		RemoveWorktree: defaultGitOps.WorktreeRemove,
	}
	return compare.ResolverFromArg(arg, deps)
}
