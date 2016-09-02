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

function primePandigital() {
  for (var i = 7654321; i > 0; i--) {
    if (isPanidigital(i)) {
      if (isPrime(i)) {
        return i;
      }
    }
  }
}

function isPanidigital(num) {
  var arr = num.toString().split('');
  var max = 0;
  for (var i = 0; i < arr.length; i++) {
    if (Number(arr[i]) > max) {
      max = Number(arr[i]);
    }
  }

  if (num.toString().length === _.uniq(num.toString().split('')).length && num.toString().indexOf("0") === -1 && max === num.toString().length) {
    return true;
  }
  return false;
}
