let count = 0;
let oldCount = 0;

let solutionFound = false;
let testNumber = 1;

for (let i = 1; i <= 21; i++) {
  for (let j = 1; j <= 9; j++) {
    const result = Math.pow(j, i);
    if (result.toString().length === i) {
      count++;
    }
  }
}

console.log(count);
