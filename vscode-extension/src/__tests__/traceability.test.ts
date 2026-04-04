// Copyright (c) 2026 Petar Djukic. All rights reserved.
// SPDX-License-Identifier: MIT

import { describe, it, expect } from "vitest";
import { SRD_PATTERN, UC_PATTERN } from "../traceability";

describe("SRD_PATTERN", () => {
  it("matches a srd comment line", () => {
    const match = "// srd: srd006-vscode-extension R8.1".match(SRD_PATTERN);
    expect(match).not.toBeNull();
    expect(match![1]).toBe("srd006-vscode-extension");
  });

  it("matches with extra spaces", () => {
    const match = "//  srd:  srd001-orchestrator-core R1".match(SRD_PATTERN);
    expect(match).not.toBeNull();
    expect(match![1]).toBe("srd001-orchestrator-core");
  });

  it("rejects uc comment lines", () => {
    const match = "// uc: rel02.0-uc006-specification-browser".match(
      SRD_PATTERN
    );
    expect(match).toBeNull();
  });

  it("rejects lines without //", () => {
    const match = "srd: srd006-vscode-extension R8.1".match(SRD_PATTERN);
    expect(match).toBeNull();
  });
});

describe("UC_PATTERN", () => {
  it("matches a uc comment line", () => {
    const match = "// uc: rel02.0-uc006-specification-browser".match(
      UC_PATTERN
    );
    expect(match).not.toBeNull();
    expect(match![1]).toBe("rel02.0-uc006-specification-browser");
  });

  it("matches with variable whitespace", () => {
    const match = "//  uc:  rel01.0-uc001-orchestrator-initialization".match(
      UC_PATTERN
    );
    expect(match).not.toBeNull();
    expect(match![1]).toBe("rel01.0-uc001-orchestrator-initialization");
  });

  it("rejects srd comment lines", () => {
    const match = "// srd: srd006-vscode-extension R8".match(UC_PATTERN);
    expect(match).toBeNull();
  });

  it("rejects lines without //", () => {
    const match = "uc: rel02.0-uc006-specification-browser".match(UC_PATTERN);
    expect(match).toBeNull();
  });
});
