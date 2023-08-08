class Queue {
    constructor(){
        this.arr = []
    }
    getLength(){
        return this.arr.length
    }
    queue(el){
        this.arr.push(el)
    }
    unqueue(){
        return this.arr.shift()
    }
}
const arr = [1,23,4,5,45,645,645,7,0]
function main(arr){
    let q = new Queue()
    arr.forEach(e=>{
        q.queue(e)
    })
    console.log(q.getLength())
    while(q.getLength()!==0){
        q.unqueue()
        console.log(q.arr)
    }
    console.log(q.arr)
}
main(arr)