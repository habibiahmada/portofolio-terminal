#!/usr/bin/env node

"use strict";

const { execFileSync } = require("child_process");
const crypto = require("crypto");
const path = require("path");
const fs = require("fs");

/**
 * Detect the current platform and architecture.
 * Maps Node.js process.platform/arch to Go binary naming.
 */
function detectPlatform() {
  const platform = process.platform;
  const arch = process.arch;

  let os;
  switch (platform) {
    case "darwin":
      os = "darwin";
      break;
    case "linux":
      os = "linux";
      break;
    case "win32":
      os = "win";
      break;
    default:
      throw new Error(`Unsupported platform: ${platform}`);
  }

  let goArch;
  switch (arch) {
    case "x64":
      goArch = "x64";
      break;
    case "arm64":
      goArch = "arm64";
      break;
    default:
      throw new Error(`Unsupported architecture: ${arch}`);
  }

  return { os, goArch };
}

/**
 * Get the binary filename for the current platform.
 */
function getBinaryName() {
  const { os, goArch } = detectPlatform();
  const ext = os === "win" ? ".exe" : "";
  return `habibiahmada-${os}-${goArch}${ext}`;
}

/**
 * Resolve the binary path and ensure it stays inside the package bin/ directory.
 */
function resolveBinaryPath() {
  const binDir = path.resolve(__dirname, "bin");
  const binaryPath = path.resolve(binDir, getBinaryName());
  const relative = path.relative(binDir, binaryPath);
  if (relative.startsWith("..") || path.isAbsolute(relative)) {
    throw new Error("Invalid binary path");
  }
  return binaryPath;
}

/**
 * Verify SHA-256 checksum when checksums.json is present (published packages).
 */
function verifyBinaryIntegrity(binaryPath) {
  const checksumsPath = path.join(__dirname, "checksums.json");
  if (!fs.existsSync(checksumsPath)) {
    return;
  }

  const checksums = JSON.parse(fs.readFileSync(checksumsPath, "utf8"));
  const binaryName = path.basename(binaryPath);
  const expected = checksums[binaryName];
  if (!expected) {
    // Dev installs and partial platforms may omit checksums.
    return;
  }

  const hash = crypto
    .createHash("sha256")
    .update(fs.readFileSync(binaryPath))
    .digest("hex");

  if (hash !== expected) {
    throw new Error(`Binary integrity check failed for ${binaryName}`);
  }
}

/**
 * Main entry point.
 */
function main() {
  // Handle postinstall hook.
  if (process.argv.includes("--postinstall")) {
    if (process.platform !== "win32") {
      const binaryPath = resolveBinaryPath();
      if (fs.existsSync(binaryPath)) {
        fs.chmodSync(binaryPath, 0o755);
      }
    }
    return;
  }

  const binaryPath = resolveBinaryPath();

  if (!fs.existsSync(binaryPath)) {
    console.error(`Binary not found: ${binaryPath}`);
    console.error(
      "Please install the package again or report this issue."
    );
    process.exit(1);
  }

  verifyBinaryIntegrity(binaryPath);

  if (process.platform !== "win32") {
    try {
      fs.accessSync(binaryPath, fs.constants.X_OK);
    } catch {
      fs.chmodSync(binaryPath, 0o755);
    }
  }

  // execFileSync avoids shell interpretation of arguments (no command injection).
  try {
    execFileSync(binaryPath, process.argv.slice(2), {
      stdio: "inherit",
    });
  } catch (err) {
    if (err.status !== undefined) {
      process.exit(err.status);
    }
    console.error(err.message);
    process.exit(1);
  }
}

main();
