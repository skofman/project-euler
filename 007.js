function isPrime(num) {
  for (var i = 2; i < Math.floor(num / 2); i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

function findPrime(index) {
  var arr = [2,3];
  var i = 5;
  while (arr.length < index) {
    if (isPrime(i)) {
      arr.push(i);
    }
    i++;
  }
  return arr[arr.length - 1];
}
