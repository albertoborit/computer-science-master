const {parentPort} = require("worker_threads")
parentPort.on('message', (n)=>{
    let count = 0
    for(let i = 0; i < n; i++){
        count++
    }
    console.log(count)
    parentPort.postMessage('done')
})