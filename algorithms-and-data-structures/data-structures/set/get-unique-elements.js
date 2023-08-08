function main(arr){
	const set = new Set(arr)
  return [...set]
}

const arr = [1,1,2,3,4,5,5,7,7,7,7,0,8]
console.log(main(arr))