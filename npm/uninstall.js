#!/usr/bin/env node

const rimraf = require("rimraf");

const { existsSync } = require("fs");

const { getBinaryDir } = require("./getBinary");

const binaryDirectory = getBinaryDir();

if (existsSync(binaryDirectory)) {
  rimraf.sync(binaryDirectory);
}
