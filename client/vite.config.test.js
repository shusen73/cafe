import { describe, it, expect } from "vitest";
import config from "./vite.config.js";

describe("vite build config", () => {
  it("outputs to server/public", () => {
    const outDir = config?.build?.outDir ?? "";
    expect(String(outDir)).toContain("server/public");
  });
});
