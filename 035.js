function isPrime(num) {
  if (num <= 1) {
    return false;
  }
  if (num <= 3) {
    return true;
  }
  for (var i = 2; i <= Math.floor(num / 2); i++) {
    if (num % i === 0) {
      return false;
    }
  }
  return true;
}

function circularPrimes(num) {
  var str;
  var arr;
  var count = 0;
  for (var i = 2; i < num; i++) {
    arr = i.toString().split('');
    if (primesConfirm(arr)) {
      count++;
    }
  }
  return count;
}

function primesConfirm(arr) {
  for (var i = 0; i < arr.length; i++) {
    if (!isPrime(Number(arr.join('')))) {
      return false;
    }
    arr.push(arr.shift());
  }
  return true;
}
