import { spawn } from "child_process";
import es from "event-stream";

export function psTree(pid) {
  var headers = null;

  if (typeof pid === "number") {
    pid = pid.toString();
  }

  var processLister;
  if (process.platform === "win32") {
    processLister = spawn("wmic.exe", [
      "PROCESS",
      "GET",
      "Name,ProcessId,ParentProcessId,Status",
    ]);
  } else {
    processLister = spawn("ps", ["-A", "-o", "ppid,pid,stat,comm"]);
  }

  return new Promise((res, rej) => {
    es.connect(
      processLister.stdout,
      es.split(),
      es.map(function (line, cb) {
        var columns = line.trim().split(/\s+/);
        if (!headers) {
          headers = columns;
          headers = headers.map(normalizeHeader);
          return cb();
        }

        var row = {};
        var h = headers.slice();
        while (h.length) {
          row[h.shift()] = h.length ? columns.shift() : columns.join(" ");
        }

        return cb(null, row);
      }),
      es.writeArray(function (_, ps) {
        var parents = {},
          children = [];

        parents[pid] = true;
        ps.forEach(function (proc) {
          if (parents[proc.PPID]) {
            parents[proc.PID] = true;
            children.push(proc);
          }
        });
        res(children);
      }),
    ).on("error", rej);
  });
}

/**
 * Normalizes the given header `str` from the Windows
 * title to the *nix title.
 *
 * @param {string} str Header string to normalize
 */
function normalizeHeader(str) {
  switch (str) {
    case "Name": // for win32
    case "COMM": // for darwin
      return "COMMAND";
    case "ParentProcessId":
      return "PPID";
    case "ProcessId":
      return "PID";
    case "Status":
      return "STAT";
    default:
      return str;
  }
}
