const { performance } = require("perf_hooks");

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
  for (let i = 0; i < 5; i++) {
    block1Second(i);
  }

  await sleep(2000);
}

main();
