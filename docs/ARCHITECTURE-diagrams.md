# Architecture Diagrams

Companion diagrams for [ARCHITECTURE.yaml](ARCHITECTURE.yaml).

|  |
|:--:|

```plantuml
@startuml
!theme plain
skinparam backgroundColor white

package "Consuming Project" {
  [Magefile] <<mage targets>>
}

package "orchestrator" {
  [Orchestrator] <<main struct>>
  [Generator] <<lifecycle>>
  [Cobbler] <<measure + stitch>>
  [Commands] <<git, beads, go wrappers>>
  [Stats] <<metrics>>
}

package "External Tools" {
  [Git]
  [Claude Code]
  [Beads (bd)]
  [Go Toolchain]
}

[Magefile] --> [Orchestrator]
[Orchestrator] --> [Generator]
[Orchestrator] --> [Cobbler]
[Orchestrator] --> [Stats]
[Generator] --> [Commands]
[Cobbler] --> [Commands]
[Cobbler] --> [Claude Code]
[Commands] --> [Git]
[Commands] --> [Beads (bd)]
[Commands] --> [Go Toolchain]

@enduml
```

|Figure 1 System context showing orchestrator components and external tools |
