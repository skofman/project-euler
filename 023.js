function findDivisorsSum(num) {
  var sum = 0;
  for (var i = 1; i <= num / 2; i++) {
    if (num % i === 0) {
      sum += i;
    }
  }
  return sum;
}

function abundant() {
  var arr = [];

  for (var i = 2; i <=28123; i++) {
    if (findDivisorsSum(i) > i) {
      arr.push(i);
    }
  }
  var sum = 0;
  for (var i = 1; i <= 28123; i++) {
    if (checkedOut(i, arr)) {
      sum += i;
    }
  }
  return sum;
}

function checkedOut(num, arr) {
  var test = 0;
  for (var i = 0; i < arr.length; i++) {
    if (num < arr[i]) {
      return true;
    }
    test = num - arr[i];
    if (arr.includes(test)) {
      return false;
    }
  }
  return true;
}
