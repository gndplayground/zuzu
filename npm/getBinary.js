const os = require("os");
const { join } = require("path");

function getPlatform() {
  const type = os.type();
  const arch = os.arch();

  if (type === "Windows_NT" && arch === "x64") return "window";
  if (type === "Windows_NT") return "win32";
  if (type === "Linux" && arch === "x64") return "linux";
  if (type === "Darwin" && arch === "x64") return "mac";
  if (type === "Darwin" && arch === "arm64") return "mac-silicon";

  throw new Error(`Unsupported platform: ${type} ${arch}`);
}

// Hard code for now
function getConfig() {
  const platform = getPlatform();
  const version = require("../package.json").version;
  const url = `https://github.com/gndplayground/zuzu/releases/download/v${version}/zuzu-${platform}.tar.gz`;
  const name = "zuzu";
  return {
    platform: platform,
    version,
    url,
    name,
    dir: "./bin",
  };
}

function getBinaryDir() {
  const config = getConfig();
  return config.dir;
}

function getBinaryPath() {
  const config = getConfig();
  return join(config.dir, `${config.name}-${config.platform}`);
}

module.exports = {
  getConfig,
  getPlatform,
  getBinaryPath,
  getBinaryDir,
};
