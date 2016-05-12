function isPrime(num) {
  for (var i = 2; i < Math.floor(num / 2); i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

function largePrime(num) {
  for (var i = 1; i < num; i++) {
    if (num % i === 0) {
      var div = num / i;
      if (isPrime(div)) {
        return div;
      }
    }
  }
  return 1;
}
