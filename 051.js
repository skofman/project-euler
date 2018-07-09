const isPrime = num => {
  if (num === 2) {
    return true;
  }
  if (num % 2 === 0) {
    return false;
  }
  for (let i = 3; i <= Math.sqrt(num); i += 2) {
    if (num % i === 0) {
      return false;
    }
  }

  return true;
};

const getNextPrime = num => {
  let i = num + 1;
  if (isPrime(i)) {
    return i;
  } else {
    return getNextPrime(i);
  }
};

const factor = num => {
  if (num === 0 || num === 1) {
    return 1;
  }
  return num * factor(num - 1);
};

let solutionFound = false;
let prime = 101;
let primes = 0;
let nonPrimes = 0;
let primeList = [];

while (!solutionFound) {
  const length = prime.toString().length - 1;
  const end = Math.pow(2, length);
  const num = prime.toString().split('');

  for (let j = 1; j < end; j++) {
    let test = j.toString(2);
    switch (length - test.length) {
      case 1:
        test = '0' + test;
        break;
      case 2:
        test = '00' + test;
        break;
      case 3:
        test = '000' + test;
        break;
      case 4:
        test = '0000' + test;
        break;
    }
    for (let i = 0; i < 10; i++) {
      const testNumber = [...num];
      for (let k = 0; k < test.length; k++) {
        if (test[k] === '1') {
          testNumber[k] = i;
        }
      }
      if (isPrime(+testNumber.join('')) && testNumber[0] !== 0) {
        primes++;
        primeList.push(testNumber.join(''));
      } else {
        nonPrimes++;
      }
    }
    if (primes === 8) {
      console.log(primeList);
      solutionFound = true;
      break;
    } else {
      primes = 0;
      nonPrimes = 0;
      primeList = [];
    }
  }
  prime = getNextPrime(prime);
}
