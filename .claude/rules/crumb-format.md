<!-- Copyright (c) 2026 Petar Djukic. All rights reserved. SPDX-License-Identifier: MIT -->

# Issue Format

Cupboard issues fall into two deliverable types: **documentation** (files under `docs/`) and **code** (implementation). Issue descriptions are written in YAML. The description must make the deliverable type and output location explicit so agents know what to produce and where.

## Common Structure (All Issues)

Every issue description is a YAML document with these fields:

```yaml
deliverable_type: documentation  # or: code
format_rule: prd-format           # the rule file to follow (documentation issues only)

required_reading:
  - docs/ARCHITECTURE.yaml (components section)
  - docs/specs/product-requirements/prd001-feature.yaml

files:
  - path: docs/specs/product-requirements/prd-feature-name.yaml
    action: create

requirements:
  - What needs to be built or written

design_decisions:
  - Technical or structural choices to follow

acceptance_criteria:
  - Checkable outcome one
  - Checkable outcome two
```

Table 1 Common issue fields

| Field | Required | Description |
| ----- | -------- | ----------- |
| deliverable_type | yes | `documentation` or `code` |
| format_rule | doc issues only | Rule file that governs the output format |
| required_reading | yes | Files the agent must read before starting |
| files | yes | Files to create or modify with action (`create` or `modify`) |
| requirements | yes | What needs to be built or written |
| design_decisions | no | Architecture, patterns, or constraints to follow |
| acceptance_criteria | yes | Checkable outcomes |

For epics, the description can be higher level; child tasks carry the detailed structure.

## Documentation Issues

Documentation issues produce files under `docs/`. The issue must specify the output location and which format rule applies.

### Output Location and Format Rule

Table 2 Documentation deliverable types

| Deliverable type | Output location | Format rule | When to use |
| ----------------- | ---------------- | ----------- | ------------ |
| **ARCHITECTURE** | `docs/ARCHITECTURE.yaml` | architecture-format | Updating system overview, components, design decisions |
| **PRD** | `docs/specs/product-requirements/prd[NNN]-[feature-name].yaml` | prd-format | New or updated product requirements |
| **Use case** | `docs/specs/use-cases/rel[NN].[N]-uc[NNN]-[short-name].yaml` | use-case-format | Tracer-bullet flows, actor/trigger, demo criteria |
| **Test suite** | `docs/specs/test-suites/test-[use-case-id].yaml` | test-case-format | Test cases with inputs and expected outputs |
| **Engineering guideline** | `docs/engineering/eng[NN]-[short-name].md` | engineering-guideline-format | Conventions and practices |
| **Specification** | `docs/SPECIFICATIONS.md` | specification-format | Summary of PRDs, use cases, test suites, roadmap |

### Example (PRD issue)

```yaml
deliverable_type: documentation
format_rule: prd-format

required_reading:
  - docs/ARCHITECTURE.yaml (components section)
  - docs/specs/product-requirements/prd001-cupboard-core.yaml

files:
  - path: docs/specs/product-requirements/prd-feature-name.yaml
    action: create

required_sections:
  - "Problem: explain the problem and why it matters"
  - "Goals: G1 ..., G2 ..."
  - "Requirements: R1.1 ..., R1.2 ..."
  - "Non-Goals: what is out of scope"
  - "Acceptance Criteria: checkable outcomes"

acceptance_criteria:
  - All required sections present
  - File saved as prd-feature-name.yaml
  - Requirements numbered and specific
```

## Code Issues

Code issues produce or change implementation (Go, config, tests) outside of `docs/`. Do not include PRD-style Problem/Goals/Non-Goals in code issues.

### Example (code issue)

```yaml
deliverable_type: code

required_reading:
  - docs/specs/product-requirements/prd003-crumbs-interface.yaml
  - pkg/types/cupboard.go

files:
  - path: pkg/types/crumb.go
    action: create
    note: Crumb struct, Filter type
  - path: internal/sqlite/crumbs.go
    action: create
    note: CrumbTable implementation
  - path: internal/sqlite/crumbs_test.go
    action: create
    note: tests

requirements:
  - Implement CrumbTable interface per prd003-crumbs-interface
  - Add, Get, Archive, Purge, Fetch operations
  - Property operations (Set/Get/Clear)

design_decisions:
  - Use table accessor pattern from prd001-cupboard-core
  - Filter as map[string]any per PRD

acceptance_criteria:
  - All CrumbTable operations implemented
  - Tests pass for each operation
  - Errors match PRD error types
```

### Go Layout (Recommended)

- **pkg/** – Shared public API: types and interfaces. No implementation; importable by other modules.
- **internal/** – Private implementation details. Not importable outside the module.
- **cmd/** – Entry points and executables.

When proposing or implementing code issues, keep implementation in **internal/** not **pkg/**.

## Quick Reference

Table 3 Issue type quick reference

| Issue type | Output | Key fields |
| ---------- | ------ | ---------- |
| Documentation (ARCHITECTURE) | `docs/ARCHITECTURE.yaml` | required_reading, files, requirements, acceptance_criteria; format_rule: architecture-format |
| Documentation (PRD) | `docs/specs/product-requirements/prd*.yaml` | required_reading, files, required_sections, acceptance_criteria; format_rule: prd-format |
| Documentation (use case) | `docs/specs/use-cases/rel*-uc*-*.yaml` | required_reading, files, required_sections, acceptance_criteria; format_rule: use-case-format |
| Documentation (test suite) | `docs/specs/test-suites/test*.yaml` | required_reading, files, required_sections, acceptance_criteria; format_rule: test-case-format |
| Documentation (engineering guideline) | `docs/engineering/eng*.md` | required_reading, files, requirements, acceptance_criteria; format_rule: engineering-guideline-format |
| Documentation (specification) | `docs/SPECIFICATIONS.md` | required_reading, files, requirements, acceptance_criteria; format_rule: specification-format |
| Code | `pkg/`, `internal/`, `cmd/` | required_reading, files, requirements, design_decisions, acceptance_criteria |

## When Creating or Editing Issues

1. Set **deliverable_type**: `documentation` or `code`.
2. Set **format_rule** for documentation issues (the rule file that governs the output).
3. List **required_reading**: files the agent must read before starting.
4. List **files**: explicit paths and actions (`create` or `modify`) for all outputs.
5. For documentation issues: list **required_sections** from the format rule.
6. Include **requirements** and **acceptance_criteria** in every issue.
