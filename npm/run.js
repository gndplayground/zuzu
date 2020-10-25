#!/usr/bin/env node

const { spawnSync } = require("child_process");
const { join } = require("path");
const { getBinaryPath } = require("./getBinary");

const binaryPath = getBinaryPath();

const [, , ...args] = process.argv;

const options = {
  cwd: process.cwd(),
  stdio: "inherit",
};

const result = spawnSync(join(__dirname, "../", binaryPath), args, options);

if (result.error) {
  console.error(result.error);
  process.exit(1);
}

process.exit(result.status);
