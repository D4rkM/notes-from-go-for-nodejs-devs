// Node can use worker_threads to do a similar job as go routines
// BUT it does not work and hasn`t been planned to work the same way
// Go routines has concurrency based in it`s working format
// it was minded to solve problems of io
// Node Workers has paralelism

const { Worker, isMainThread, workerData } = require("worker_threads");

function block1Second(unit) {
  console.log(`starting unit ${unit}`);
  const end = performance.now() + 1000;
  while (performance.now() < end) {
    // busy wait for 1 second
  }
  console.log(`finished unit ${unit}`);
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

async function main() {
  await sleep(10000);
  if (isMainThread) {
    for (let i = 0; i < 5; i++) {
      const worker = new Worker(__filename, {
        workerData: { unit: i },
      });
    }

    await sleep(2000);
  } else {
    const { unit } = workerData;
    block1Second(unit);
  }
}

main();
