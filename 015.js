function factorial(num) {
  var result = 1;
  for (var i = 1; i <= num; i++) {
    result *= i;
  }
  return result;
}

function path(num) {
  return factorial(2 * num) / (factorial(num) * factorial(num));
}
