#!/usr/bin/env node

"use strict";

const { execFileSync } = require("child_process");
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
 * Get the binary path for the current platform.
 */
function getBinaryPath() {
  const { os, goArch } = detectPlatform();
  const ext = os === "win" ? ".exe" : "";
  const binaryName = `habibiahmada-${os}-${goArch}${ext}`;
  return path.join(__dirname, "bin", binaryName);
}

/**
 * Main entry point.
 */
function main() {
  // Handle postinstall hook.
  if (process.argv.includes("--postinstall")) {
    // Ensure binary is executable on Unix systems.
    if (process.platform !== "win32") {
      const binaryPath = getBinaryPath();
      if (fs.existsSync(binaryPath)) {
        fs.chmodSync(binaryPath, 0o755);
      }
    }
    return;
  }

  const binaryPath = getBinaryPath();

  // Check if binary exists.
  if (!fs.existsSync(binaryPath)) {
    console.error(`Binary not found: ${binaryPath}`);
    console.error(
      "Please install the package again or report this issue."
    );
    process.exit(1);
  }

  // Ensure binary is executable on Unix systems.
  if (process.platform !== "win32") {
    try {
      fs.accessSync(binaryPath, fs.constants.X_OK);
    } catch {
      fs.chmodSync(binaryPath, 0o755);
    }
  }

  // Run the Go binary, forwarding all arguments.
  try {
    execFileSync(binaryPath, process.argv.slice(2), {
      stdio: "inherit",
    });
  } catch (err) {
    // execFileSync throws when the child exits with non-zero.
    if (err.status !== undefined) {
      process.exit(err.status);
    }
    console.error(err.message);
    process.exit(1);
  }
}

main();
