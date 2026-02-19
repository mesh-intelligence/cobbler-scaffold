<!-- Copyright (c) 2026 Petar Djukic. All rights reserved. SPDX-License-Identifier: MIT -->

# Vision Document Format

A vision document states **what the project is**, **why it exists**, **how success is measured**, and **what it is not**. It orients stakeholders, new contributors, and downstream docs (ARCHITECTURE, PRDs). It is not a PRD; it does not list numbered requirements. It sets direction and boundaries.

## File and Naming

- **Location**: `docs/VISION.yaml`
- **Format**: YAML

## Required Fields

Every vision document has these top-level fields in this order.

```yaml
id: vision-[project-name]
title: Human-readable title

executive_summary: |
  One or two short paragraphs. What the project is. What it is not in one line.

problem: |
  Context and problem. Why this exists, why current solutions fall short,
  what this project does differently.

what_this_does: |
  How the system works. The core approach in concrete terms.

why_we_build_this: |
  Why this org or team is doing it. Competitive advantage, fit with strategy,
  domain expertise, integration with existing products.

related_projects:
  - project: Project Name
    role: One-sentence description of the relationship

success_criteria:
  dimension_name: |
    What success looks like along this dimension.

implementation_phases:
  - phase: "01.0"
    focus: Phase focus area
    deliverables: Key outputs of this phase

risks:
  - risk: Short risk description
    impact: What breaks if this happens
    likelihood: High / Medium / Low
    mitigation: How we address it

not:
  - We are not building X.
  - We are not building Y.
```

### id

The vision document identifier. Lowercase, hyphenated.

```yaml
id: vision-mage-claude-orchestrator
```

### title

A short human-readable name.

```yaml
title: Mage Claude Orchestrator Vision
```

### executive_summary

One or two short paragraphs. State what the project is and what it is not. Elevator pitch only.

### problem

Context and problem. Include organizational or research context, why this exists, why current solutions fall short, and what this project does differently.

### what_this_does

How the system works. The core approach in concrete terms.

### why_we_build_this

Why this org or team is doing it. Include competitive advantage, fit with strategy, and relationship to existing products. Use the `related_projects` field for structured relationships.

### related_projects

A list of related projects. Each entry contains `project` (name) and `role` (one-sentence description of the relationship).

### success_criteria

A map of named success dimensions. Each key is a dimension name (snake_case) and each value is a prose description of what success looks like.

### implementation_phases

An ordered list of phases. Each entry contains `phase` (version string), `focus` (theme), and `deliverables` (key outputs).

### not

A list of explicit boundaries. Each entry is a complete sentence starting with "We are not".

### risks

A list of risks. Each entry contains `risk`, `impact`, `likelihood`, and `mitigation`.

## Optional Fields

### references

External references or links to related documentation.

```yaml
references:
  - ARCHITECTURE.yaml
  - External reference
```

## Writing Guidelines

- **Audience**: Stakeholders, new team members, and downstream doc authors. Avoid unexplained jargon; define domain-specific terms when first used.
- **Tone**: Use "we" in active voice. Follow documentation-standards: concise, active voice, no forbidden terms.
- **Scope**: Vision sets direction and boundaries. Do not duplicate PRD-level requirements; point to PRDs for detailed requirements and acceptance criteria.
- **Location**: `docs/VISION.yaml`.

## Completeness Checklist

- [ ] id matches the project name
- [ ] executive_summary states what the project is and what it is not
- [ ] problem covers context, why current solutions fall short, and what this does differently
- [ ] what_this_does describes the core approach concretely
- [ ] why_we_build_this explains the rationale
- [ ] success_criteria covers meaningful dimensions
- [ ] implementation_phases lists phases with focus and deliverables
- [ ] risks lists known risks with impact, likelihood, and mitigation
- [ ] not lists explicit boundaries
- [ ] Style follows documentation-standards (no forbidden terms)
- [ ] File saved as `VISION.yaml` in `docs/`

## Relationship to Other Docs

- **VISION** — What we are, why we exist, what done looks like, what we are not.
- **ARCHITECTURE** — Components, interfaces, protocols. How the system is built.
- **PRDs** — Numbered requirements, acceptance criteria. What the system must do.
- **Use cases** — Tracer bullets and demos. How we validate the path.
- **Test suites** — Test cases with inputs and expected outputs. How we verify the path.
- **Engineering guidelines** — Conventions and practices. How we work with the system.

Code and PRDs should be traceable to VISION (goals, boundaries) and ARCHITECTURE (design).
