const http = require('http')
const url = require('url')
const cluster = require('cluster')
const { Worker } = require('worker_threads')
const os = require('os')

if (cluster.isMaster) {
  const numCPUs = os.cpus().length

  for (let i = 0; i < numCPUs; i++) {
    cluster.fork()
  }
} else {
  http.createServer((req, res) => {
    const { pathname } = url.parse(req.url, true)
    const match = pathname.match(/^\/fibonacci\/(\d+)$/)

    if (match) {
      const num = parseInt(match[1], 10)
      const worker = new Worker('./worker.js')

      worker.on('message', (result) => {
        res.writeHead(200, { 'Content-Type': 'application/json' })
        res.end(JSON.stringify({ result }))
      })

      worker.postMessage(num)
    } else {
      res.writeHead(404)
      res.end()
    }
  }).listen(3000)
}
