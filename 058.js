function isPrime(num) {
  if (num <= 1) {
    return false;
  }
  for (var i = 2; i < Math.floor(num / 2); i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

function spiralPrimes() {
  let walls = 1;
  let ratio = 1;
  let diagonalPrimes= [];
  let nonPrimes = [];
  let max = 1;
  let min = 1;
  while(ratio > 0.1) {
    walls += 2;
    min = max;
    for (let i = max + walls - 1; i <= max + (walls - 1) * 4; i += (walls - 1)) {
      if (isPrime(i)) {
        diagonalPrimes.push(i);
      }
      else {
        nonPrimes.push(i);
      }
    }
    max = walls * walls;
    ratio = diagonalPrimes.length / (nonPrimes.length + diagonalPrimes.length);
  }

  console.log(walls);
  console.log(diagonalPrimes);
  console.log(diagonalPrimes.length);
  console.log(diagonalPrimes.length + nonPrimes.length);
  console.log(ratio);
}

spiralPrimes();