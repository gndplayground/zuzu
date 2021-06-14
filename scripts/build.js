const { spawnSync } = require("child_process");
const { existsSync } = require("fs");

const rimraf = require("rimraf");

if (existsSync("build")) {
  rimraf.sync("build");
}

function checkResult(result) {
  if (result.error) {
    console.error(result.error);
    process.exit(1);
  }
}

const buildVersions = [
  { os: "darwin", arch: "amd64", name: "mac" },
  { os: "linux", arch: "amd64", name: "linux" },
  { os: "windows", arch: "386", name: "window" },
  { os: "darwin", arch: "arm64", name: "mac-silicon" },
];

// Might need to use worker if the build is long
buildVersions.forEach((v) => {
  const options = {
    cwd: process.cwd(),
    stdio: "inherit",
    env: {
      ...process.env,
      GOOS: v.os,
      GOARCH: v.arch,
    },
  };

  const buildPath = `build/zuzu-${v.name}`;

  let result = spawnSync("go", ["build", "-o", buildPath], options);

  checkResult(result);

  result = spawnSync("tar", ["-czvf", `${buildPath}.tar.gz`, buildPath], {
    cwd: process.cwd(),
    stdio: "inherit",
  });

  checkResult(result);
});
