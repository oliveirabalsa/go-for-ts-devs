const { parentPort } = require('worker_threads');

function fibonacci(n) {
  return n < 2 ? n : fibonacci(n - 1) + fibonacci(n - 2);
}

parentPort.on('message', (n) => {
  parentPort.postMessage(fibonacci(n));
})
