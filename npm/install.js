#!/usr/bin/env node

const mkdirp = require("mkdirp");
const axios = require("axios");
const tar = require("tar");
const rimraf = require("rimraf");

const { existsSync } = require("fs");

const { getConfig, getBinaryDir } = require("./getBinary");

const config = getConfig();

const binaryDirectory = getBinaryDir();

if (existsSync(binaryDirectory)) {
  rimraf.sync(binaryDirectory);
}

mkdirp.sync(binaryDirectory);

console.log("Downloading release", config.url);

return axios({
  url: config.url,
  responseType: "stream",
})
  .then((res) => {
    res.data.pipe(
      tar.x({
        strip: 1,
        C: binaryDirectory,
      })
    );
  })
  .then(() => {
    console.log(
      `${config.name ? config.name : "Your package"} has been installed!`
    );
  })
  .catch((e) => {
    console.error("Error fetching release", e.message);
    throw e.message;
  });
