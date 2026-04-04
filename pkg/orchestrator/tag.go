// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

package orchestrator

import (
	rel "github.com/mesh-intelligence/cobbler-scaffold/pkg/orchestrator/internal/release"
)

// Tag creates a documentation-only release tag (v0.YYYYMMDD.N) for the current
// state of the repository. The revision number increments for each tag created
// on the same date. Optionally updates the version file if configured.
//
// Tag convention:
//   - v0.* = documentation-only releases on main (manual)
//   - v1.* = Claude-generated code (created by GeneratorStop)
//
// Exposed as a mage target (e.g., mage tag).
func (r *Releaser) Tag() error {
	return rel.Tag(rel.TagParams{
		BaseBranch:   r.cfg.Cobbler.BaseBranch,
		DocTagPrefix: r.cfg.Cobbler.DocTagPrefix,
		VersionFile:  r.cfg.Project.VersionFile,
	})
}

