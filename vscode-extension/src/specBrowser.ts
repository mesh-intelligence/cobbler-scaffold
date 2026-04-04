// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

// srd: srd006-vscode-extension R8
// uc: rel02.0-uc006-specification-browser

import * as vscode from "vscode";
import * as path from "path";
import {
  SpecGraph,
  UseCase,
  Srd,
  TestSuite,
  Touchpoint,
  SourceRef,
} from "./specModel";

// ---- Tree item types ----

/** Discriminated union for all node types in the specification tree. */
export type SpecTreeItem =
  | CategoryItem
  | UseCaseItem
  | TouchpointItem
  | SrdItem
  | TestSuiteItem
  | SourceRefItem;

interface CategoryItem {
  kind: "category";
  label: string;
}

interface UseCaseItem {
  kind: "useCase";
  useCase: UseCase;
}

interface TouchpointItem {
  kind: "touchpoint";
  touchpoint: Touchpoint;
}

interface SrdItem {
  kind: "srd";
  srd: Srd;
}

interface TestSuiteItem {
  kind: "testSuite";
  testSuite: TestSuite;
}

interface SourceRefItem {
  kind: "sourceRef";
  ref: SourceRef;
  workspaceRoot: string;
}

// ---- Provider ----

/**
 * TreeDataProvider for the mageOrchestrator.specs view. Renders three
 * top-level categories (Use Cases, SRDs, Test Suites) backed by SpecGraph.
 * Expanding a use case reveals its touchpoints; expanding a touchpoint
 * with a SRD reference reveals Go source files that implement it.
 */
export class SpecBrowserProvider
  implements vscode.TreeDataProvider<SpecTreeItem>
{
  private _onDidChangeTreeData = new vscode.EventEmitter<
    SpecTreeItem | undefined | void
  >();
  readonly onDidChangeTreeData = this._onDidChangeTreeData.event;

  private graph: SpecGraph;
  private root: string;

  constructor(workspaceRoot: string) {
    this.root = workspaceRoot;
    this.graph = new SpecGraph(workspaceRoot);
  }

  /** Invalidates the SpecGraph cache and fires a tree refresh. */
  refresh(): void {
    this.graph.invalidate();
    this._onDidChangeTreeData.fire();
  }

  getTreeItem(element: SpecTreeItem): vscode.TreeItem {
    switch (element.kind) {
      case "category":
        return this.categoryTreeItem(element);
      case "useCase":
        return this.useCaseTreeItem(element);
      case "touchpoint":
        return this.touchpointTreeItem(element);
      case "srd":
        return this.srdTreeItem(element);
      case "testSuite":
        return this.testSuiteTreeItem(element);
      case "sourceRef":
        return this.sourceRefTreeItem(element);
    }
  }

  async getChildren(element?: SpecTreeItem): Promise<SpecTreeItem[]> {
    await this.graph.ensureBuilt();

    if (!element) {
      return [
        { kind: "category", label: "Use Cases" },
        { kind: "category", label: "SRDs" },
        { kind: "category", label: "Test Suites" },
      ];
    }

    switch (element.kind) {
      case "category":
        return this.categoryChildren(element.label);
      case "useCase":
        return element.useCase.touchpoints.map(
          (tp): TouchpointItem => ({ kind: "touchpoint", touchpoint: tp })
        );
      case "touchpoint":
        return this.touchpointChildren(element.touchpoint);
      default:
        return [];
    }
  }

  // ---- Category children ----

  private categoryChildren(label: string): SpecTreeItem[] {
    switch (label) {
      case "Use Cases":
        return this.graph
          .listUseCases()
          .map((uc): UseCaseItem => ({ kind: "useCase", useCase: uc }));
      case "SRDs":
        return this.graph
          .listSrds()
          .map((srd): SrdItem => ({ kind: "srd", srd }));
      case "Test Suites":
        return this.graph
          .listTestSuites()
          .map((ts): TestSuiteItem => ({ kind: "testSuite", testSuite: ts }));
      default:
        return [];
    }
  }

  // ---- Touchpoint children (source refs) ----

  private touchpointChildren(tp: Touchpoint): SpecTreeItem[] {
    if (!tp.srdId) {
      return [];
    }
    return this.graph
      .getSourceFiles(tp.srdId)
      .map(
        (ref): SourceRefItem => ({
          kind: "sourceRef",
          ref,
          workspaceRoot: this.root,
        })
      );
  }

  // ---- Tree item builders ----

  private categoryTreeItem(item: CategoryItem): vscode.TreeItem {
    const ti = new vscode.TreeItem(
      item.label,
      vscode.TreeItemCollapsibleState.Collapsed
    );
    ti.contextValue = "specCategory";
    return ti;
  }

  private useCaseTreeItem(item: UseCaseItem): vscode.TreeItem {
    const uc = item.useCase;
    const ti = new vscode.TreeItem(
      `${uc.id}: ${uc.title}`,
      vscode.TreeItemCollapsibleState.Collapsed
    );
    ti.tooltip = uc.summary;
    ti.contextValue = "specUseCase";
    ti.command = {
      command: "vscode.open",
      title: "Open Use Case",
      arguments: [vscode.Uri.file(uc.filePath)],
    };
    return ti;
  }

  private touchpointTreeItem(item: TouchpointItem): vscode.TreeItem {
    const tp = item.touchpoint;
    const label = tp.srdId
      ? `${tp.key}: ${tp.srdId} ${tp.requirementIds.join(", ")}`
      : `${tp.key}: ${tp.description}`;
    const expandable = tp.srdId !== undefined;
    const ti = new vscode.TreeItem(
      label,
      expandable
        ? vscode.TreeItemCollapsibleState.Collapsed
        : vscode.TreeItemCollapsibleState.None
    );
    ti.tooltip = tp.description;
    ti.contextValue = "specTouchpoint";
    if (tp.srdId) {
      const srd = this.graph.getSrd(tp.srdId);
      if (srd) {
        ti.command = {
          command: "vscode.open",
          title: "Open SRD",
          arguments: [vscode.Uri.file(srd.filePath)],
        };
      }
    }
    return ti;
  }

  private srdTreeItem(item: SrdItem): vscode.TreeItem {
    const srd = item.srd;
    const ti = new vscode.TreeItem(
      `${srd.id}: ${srd.title}`,
      vscode.TreeItemCollapsibleState.None
    );
    ti.contextValue = "specPrd";
    ti.command = {
      command: "vscode.open",
      title: "Open SRD",
      arguments: [vscode.Uri.file(srd.filePath)],
    };
    return ti;
  }

  private testSuiteTreeItem(item: TestSuiteItem): vscode.TreeItem {
    const ts = item.testSuite;
    const ti = new vscode.TreeItem(
      `${ts.id}: ${ts.title}`,
      vscode.TreeItemCollapsibleState.None
    );
    ti.tooltip = `Release: ${ts.release}`;
    ti.contextValue = "specTestSuite";
    ti.command = {
      command: "vscode.open",
      title: "Open Test Suite",
      arguments: [vscode.Uri.file(ts.filePath)],
    };
    return ti;
  }

  private sourceRefTreeItem(item: SourceRefItem): vscode.TreeItem {
    const ref = item.ref;
    const relativePath = path.relative(item.workspaceRoot, ref.file);
    const ti = new vscode.TreeItem(
      `${relativePath}:${ref.line}`,
      vscode.TreeItemCollapsibleState.None
    );
    ti.contextValue = "specSourceRef";
    const uri = vscode.Uri.file(
      path.isAbsolute(ref.file)
        ? ref.file
        : path.join(item.workspaceRoot, ref.file)
    );
    const line = Math.max(0, ref.line - 1);
    ti.command = {
      command: "vscode.open",
      title: "Open Source File",
      arguments: [
        uri,
        {
          selection: new vscode.Range(line, 0, line, 0),
        } as vscode.TextDocumentShowOptions,
      ],
    };
    return ti;
  }
}
