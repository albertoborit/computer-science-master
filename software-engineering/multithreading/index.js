const { Worker } = require("worker_threads")
const {
    performance
  } = require('perf_hooks');

const totalArraySize = 100000000
const numberOfCPUs = 8

const chunkify = (n) => {
    const chunks = []
    for(i=0;i<n;i++){
        chunks.push([totalArraySize/n])
    }
    return chunks
}
const run = (concurrentWorkers)=>{
    let completedWorkers = 0
    const chunks = chunkify(concurrentWorkers)
    const start = performance.now()
    chunks.forEach((a)=>{
        const worker = new Worker("./worker.js")
        worker.postMessage(a)
        worker.on('message', ()=>{
            console.log(`worker [${worker.threadId}] done`)
            completedWorkers++
            if(completedWorkers===concurrentWorkers){
                console.log(`workers done in ${performance.now() - start}ms`)
                process.exit()
            }
        })
    })
}
run(numberOfCPUs)