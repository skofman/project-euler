function amicable(num) {
  var sum = 0;
  var test = 0;
  for (var i = 1; i < num; i++) {
    test = findDivisorsSum(i);
    if (test != i && test != 1 && i === findDivisorsSum(test)) {
      sum += i;
    }
  }
  return sum;
}

function findDivisorsSum(num) {
  var sum = 0;
  for (var i = 1; i <= num / 2; i++) {
    if (num % i === 0) {
      sum += i;
    }
  }
  return sum;
}
