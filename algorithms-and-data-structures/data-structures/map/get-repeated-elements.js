function main(arr){
	return getDuplicates(arr).filter(e=>e===e).flat(Infinity)
}

function getDuplicates(arrWithDuplicates){
	return arrWithDuplicates.map(e=>{
  	return countDuplicates(arrWithDuplicates,e).split("")
  })
}

function countDuplicates(arr, el){
	let counter = 0
  for(let i = 0; i<arr.length; i++){
  	if(arr.includes(el)) counter++
    delete arr[arr.indexOf(el)]
  }
  return Array(counter).join(el)
}

const arr = [1,1,2,3,4,5,5,7,7,7,7,0,8]
console.log(main(arr))