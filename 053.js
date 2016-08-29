function factorial(num) {
  if (num === 0) {
    return 1;
  }

  return num * factorial(num - 1);
}

function findValues() {
  var count = 0;

  for (var i = 1; i <= 100; i++) {
    for (var j = 1; j <= i; j++) {
      if (factorial(i) / (factorial(j) * factorial(i - j)) > 1000000) {
        count++;
      }
    }
  }

  return count;
}
