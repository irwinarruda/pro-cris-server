#!/usr/bin/env node
import fs from "node:fs/promises";
import { spawn } from "node:child_process";
import { psTree } from "./ps-tree.js";

function createCmd(command) {
  const cmdProcess = spawn(command[0], command.slice(1), {
    stdio: "inherit",
    env: process.env,
    detached: false,
    killSignal: "SIGINT",
  });
  return cmdProcess;
}

async function main(args) {
  const command = args;
  if (!command) {
    console.log("No callback provided.");
    return;
  }

  let debounce = null;
  let cmdProcess = createCmd(command);

  const watch = fs.watch("./", { recursive: true });
  let time = Date.now();
  for await (const _ of watch) {
    if (debounce) {
      clearTimeout(debounce);
      debounce = null;
      time = Date.now();
    }
    debounce = setTimeout(async () => {
      if (process.env.OS_TYPE === "macos") {
        const children = await psTree(cmdProcess.pid);
        spawn("kill", ["-9"].concat(children.map((p) => p.PID)));
      }
      cmdProcess = createCmd(command);
      console.clear();
      console.log(
        `===================================== Restarted in ${Date.now() - time}ms =====================================`,
      );
      debounce = null;
      time = Date.now();
    }, 60);
  }
}

main(process.argv.slice(2));
