function dividers(num) {
  var test = num;
  for (var i = num; i > 0; i--) {
    if (test % i != 0) {
      test += num;
      i = num;
    }
  }
  return test;
}
