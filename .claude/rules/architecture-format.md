<!-- Copyright (c) 2026 Petar Djukic. All rights reserved. SPDX-License-Identifier: MIT -->

# Architecture Document Format

An architecture document describes **how the system is built**: components, interfaces, protocols, data flow, and design decisions. It is the bridge between vision (what and why) and PRDs (numbered requirements). It does not list requirements; it describes structure, contracts, and rationale so implementers and PRD authors stay aligned.

## File and Naming

- **Location**: `docs/ARCHITECTURE.yaml`
- **Format**: YAML

## Required Fields

Every architecture document has these top-level fields in this order.

```yaml
id: architecture-[project-name]
title: Human-readable title
overview:
  summary: |
    What the system does in one or two paragraphs. Core insight.
  lifecycle: |
    Main state machines or lifecycles. Link to PRD for full spec.
  coordination_pattern: |
    How components interact. Sub-sections as needed.
interfaces:
  - name: Interface Name
    summary: |
      The contract between major parts of the system.
    data_structures:
      - "TypeA: role"
    operations:
      - "Op1: purpose (prd001-feature R2)"
    announcements:
      - "EventX: what is broadcast and when"
components:
  - name: Component Name
    responsibility: |
      One or two sentences describing what this component does.
    capabilities:
      - Capability one
      - Capability two
    references:
      - prd001-feature
design_decisions:
  - id: 1
    title: Interface Design
    decision: |
      Short statement of the decision.
    benefits:
      - Benefit one
    alternatives_rejected:
      - "Alternative A: reason rejected"
technology_choices:
  - component: Component or layer
    technology: Technology name
    purpose: One-line purpose
project_structure:
  - path: cmd/
    role: Entry points. Minimal wiring.
  - path: internal/
    role: Private implementation. Not importable outside the module.
  - path: pkg/
    role: Shared public types and interfaces. No implementation.
implementation_status:
  current_focus: Phase or theme description
  phases:
    - name: Phase 1
      focus: What this phase delivers
      deliverables: Key outputs
  progress:
    - done: Item completed
    - in_progress: Item underway
related_documents:
  - doc: VISION.yaml
    purpose: Goals and boundaries.
  - doc: docs/specs/product-requirements/prd001-feature.yaml
    purpose: Requirements for the core interface.
references:
  - External reference or link
```

### id

The architecture document identifier. Lowercase, hyphenated.

```yaml
id: architecture-cupboard
```

### title

A short human-readable name.

```yaml
title: Cupboard Storage Architecture
```

### overview

Describes what the system does and how it coordinates. Contains three sub-fields.

- **summary** — What the system does in one or two paragraphs. State the core insight.
- **lifecycle** — Main state machines or lifecycles. Link to PRD for the full spec.
- **coordination_pattern** — How components interact (e.g., pull-based workers, event-driven). Add sub-keys as needed.

Optional sub-fields: `branch_model`, `deployment_workflow`.

### interfaces

A list of the contracts between major parts of the system. Each interface entry contains:

| Field | Required | Description |
|-------|----------|-------------|
| name | yes | Interface name |
| summary | yes | What this interface does |
| data_structures | no | Short list of types and roles |
| operations | no | List of operations with purpose; reference PRD requirement IDs |
| announcements | no | Events that are broadcast, when, and payload shape |

### components

A list of major components. Each entry contains:

| Field | Required | Description |
|-------|----------|-------------|
| name | yes | Component name |
| responsibility | yes | One or two sentences on what it does |
| capabilities | no | Bullet list of what it can do |
| references | no | PRDs or use cases where details live |

### design_decisions

A numbered list of architectural decisions. Each entry contains:

| Field | Required | Description |
|-------|----------|-------------|
| id | yes | Sequential integer |
| title | yes | Short label |
| decision | yes | Statement of the decision |
| benefits | no | List of benefits |
| alternatives_rejected | no | Alternatives and why they were rejected |

### technology_choices

A list of technology selections. Each entry contains `component`, `technology`, and `purpose`. Link to a Technology Stack PRD for schema, API definitions, and config.

### project_structure

A list of directory paths and their roles. Each entry contains `path` and `role`. Clarify what is shared (pkg) versus internal.

### implementation_status

Describes the current focus, optional phases, and a progress checklist. Contains:

- **current_focus** — Phase or theme name.
- **phases** — Optional list of phases with `name`, `focus`, and `deliverables`.
- **progress** — Optional checklist with `done` and `in_progress` items.

Link to VISION and PRDs for context.

### related_documents

A list of related documents. Each entry contains `doc` (file path or name) and `purpose` (one sentence).

### references

Optional. External references or links. If references live in linked docs, write `"See PRDs"`.

## Optional Fields

These fields follow `related_documents` when present.

### figures

Inline PlantUML diagrams (per documentation-standards). Since ARCHITECTURE is now a YAML file, diagrams belong in a companion markdown file or in an `overview` document that renders them. Reference the diagram file path.

```yaml
figures:
  - path: docs/ARCHITECTURE-diagrams.md
    caption: "Figure 1 System context and component relationships"
```

## Writing Guidelines

- **Audience**: Implementers, PRD authors, and reviewers. Assume readers need to build or extend the system and to trace decisions to requirements.
- **Tone**: Use "we" in active voice. Follow documentation-standards: concise, active voice, no forbidden terms.
- **Scope**: Describe structure and contracts; do not duplicate PRD-level requirements. Point to PRDs for field specs, operation signatures, acceptance criteria, and state machine details.
- **Location**: `docs/ARCHITECTURE.yaml`.

## Completeness Checklist

- [ ] id matches the project name
- [ ] title describes the architecture
- [ ] overview.summary states what the system does and the core insight
- [ ] overview.lifecycle and overview.coordination_pattern are present
- [ ] interfaces list data structures, operations, and announcements; link to PRD for full spec
- [ ] components list each major component with responsibility and link to PRD or use case
- [ ] design_decisions are numbered with decision statement and benefits
- [ ] technology_choices table covers each major layer
- [ ] project_structure shows directory paths and roles
- [ ] implementation_status reflects current focus
- [ ] related_documents lists VISION, PRDs, use cases, test suites, engineering guidelines
- [ ] Style follows documentation-standards (no forbidden terms)
- [ ] File saved as `ARCHITECTURE.yaml` in `docs/`

## Relationship to Other Docs

- **VISION** — What we are and why; success criteria and phases. Architecture implements the vision.
- **ARCHITECTURE** — Components, interfaces, protocols, design decisions. How the system is built.
- **PRDs** — Numbered requirements, field specs, operation contracts. Architecture points to PRDs for detail.
- **Use cases** — Tracer bullets and flows. Architecture describes the components and interfaces those flows use.
- **Test suites** — Test cases with inputs and expected outputs. Validate use case success criteria.
- **Engineering guidelines** — Conventions and practices above the code layer.

Code and PRDs should be traceable to ARCHITECTURE (components, interfaces) and VISION (goals).
