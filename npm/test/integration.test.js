#!/usr/bin/env node

"use strict";

const { execFileSync } = require("child_process");
const path = require("path");
const fs = require("fs");
const os = require("os");
const assert = require("assert");

const ROOT = path.resolve(__dirname, "..");
const INDEX = path.join(ROOT, "index.js");
const BIN = path.join(ROOT, "bin");

function runWrapper(args, opts) {
  return execFileSync(process.execPath, [INDEX, ...args], {
    stdio: "pipe",
    ...opts,
  }).toString();
}

const SKIP = process.platform === "win32";

// Clean up any test binary from previous runs.
function cleanup() {
  try {
    fs.unlinkSync(path.join(BIN, "habibiahmada-linux-amd64"));
  } catch {}
}

async function main() {
  if (SKIP) {
    console.log("skipped (windows runner)");
    return;
  }

  fs.mkdirSync(BIN, { recursive: true });

  // Ensure a clean state: no stub binary present.
  cleanup();

  // 1. postinstall must succeed without a binary present.
  runWrapper(["--postinstall"]);

  // 2. Running without a binary must produce a clear error.
  try {
    runWrapper([]);
    assert.fail("expected wrapper to fail when no binary exists");
  } catch (err) {
    const output = (err.stdout || "") + (err.stderr || "");
    assert.match(output, /Binary not found/, "expected 'Binary not found' error");
  }

  // 3. With a stub binary present, args must be forwarded and exit 0.
  // Temporarily clear checksums — the stub is not the release binary.
  cleanup();
  const checksumsPath = path.join(ROOT, "checksums.json");
  let savedChecksums = null;
  if (fs.existsSync(checksumsPath)) {
    savedChecksums = fs.readFileSync(checksumsPath);
    fs.writeFileSync(checksumsPath, "{}\n");
  }
  const stub = path.join(BIN, "habibiahmada-linux-amd64");
  fs.writeFileSync(
    stub,
    "#!/bin/sh\necho wrapper-ok:$1\n",
    { mode: 0o755 },
  );
  const out = runWrapper(["--stub-flag"]);
  assert.match(out, /wrapper-ok:--stub-flag/, "expected stub to forward args");
  if (savedChecksums) {
    fs.writeFileSync(checksumsPath, savedChecksums);
  }
  cleanup();

  console.log("all integration checks passed");
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});