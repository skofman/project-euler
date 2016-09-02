function truncatedPrimes() {
  var count = 11;
  var sum = 0;
  var i = 20;
  while (count > 0) {
    if (checkTruncation(i)) {
      sum += i;
      count--;
    }
    i++;
  }
  return sum;
}

function checkTruncation(num) {
  var length = num.toString().length;
  var arr = [];
  for (var i = 1; i < length; i++) {
    arr.push(num.toString().substr(0, i));
    arr.push(num.toString().substr(length - i, i));
  }
  arr.push(num);
  for (var j = 0; j < arr.length; j++) {
    if (!isPrime(Number(arr[j]))) {
      return false;
    }
  }
  return true;
}

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
