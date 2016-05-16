function isPrime(num) {
  for (var i = 2; i < Math.floor(num / 2); i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

function primeSum(num) {
  var sum = 5;
  for (var i = 5; i < num; i++) {
    if (isPrime(i)) {
      sum += i;
    }
  }
  return sum;
}
