function collatz(num) {
  var max = 0;
  var result = 0;
  var chain = 0;
  var test = 0;
  for (var i = num - 1; i > num / 2; i--) {
    test = i;
    while (test != 1) {
      chain++;
      if (test % 2 === 0) {
        test /= 2;
      }
      else {
        test = 3 * test + 1;
      }
    }
    if (chain > max) {
      max = chain;
      result = i;
    }
    chain = 0;
  }
  return result;
}
