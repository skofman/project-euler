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

const checkTwoNumbers = (a, b) => {
  return isPrime(+`${a}${b}`) && isPrime(+`${b}${a}`);
};

let sumFound = false;
let testPrime = 3;
let primes = [3];
let testingPrime = getNextPrime(3);

while (!sumFound) {
  if (primes.length === 1) {
    if (checkTwoNumbers(primes[0], testingPrime)) {
      primes.push(testingPrime);
    }
  } else if (primes.length === 2) {
    if (checkTwoNumbers(primes[0], testingPrime) && checkTwoNumbers(primes[1], testingPrime)) {
      primes.push(testingPrime);
    }
  } else if (primes.length === 3) {
    if (
      checkTwoNumbers(primes[0], testingPrime) &&
      checkTwoNumbers(primes[1], testingPrime) &&
      checkTwoNumbers(primes[2], testingPrime)
    ) {
      primes.push(testingPrime);
    }
  } else if (primes.length === 4) {
    if (
      checkTwoNumbers(primes[0], testingPrime) &&
      checkTwoNumbers(primes[1], testingPrime) &&
      checkTwoNumbers(primes[2], testingPrime) &&
      checkTwoNumbers(primes[3], testingPrime)
    ) {
      primes.push(testingPrime);
    }
  }

  if (primes.length === 5) {
    sumFound = true;
  } else if (primes.length === 0) {
    console.log('Increase the test');
    sumFound = true;
  }
  if (testingPrime > 10000) {
    if (primes.length === 1) {
      primes = [getNextPrime(primes[0])];
      testingPrime = getNextPrime(primes[0]);
    } else {
      testingPrime = getNextPrime(primes[primes.length - 1]);
      primes.splice(-1);
    }
  } else {
    testingPrime = getNextPrime(testingPrime);
  }
}

let sum = 0;
for (let i = 0; i < primes.length; i++) {
  sum += primes[i];
}

console.log(sum);
