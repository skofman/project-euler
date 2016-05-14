function findDiff(num) {
  var sum = 0;
  var squares = 0;
  for (var i = 1; i <= num; i++) {
    sum += i;
    squares += i * i;
  }
  return sum * sum - squares;
}
