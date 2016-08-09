function isPrime(num) {
  if (num <= 1) {
    return false;
  }
  if (num === 2) {
    return true;
  }
  for (var i = 2; i < Math.floor(num / 2); i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

function findPrimes() {
  var n = 0;
  var max = 0;
  var result;
  var test;
  for (var a = -999; a < 1000; a++) {
    for (var b = -999; b < 1000; b++) {
      while (isPrime(n * n + a * n + b)) {
        n++;
      }
      if (n > max) {
        max = n;
        result = a * b;
      }
      n = 0;
    }
  }
  return result;
}
