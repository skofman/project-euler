function spiral() {
  var sum = 1;
  var prevRight = 1;
  var prevLeft = 1;
  for (var i = 1; i <= 500; i++) {
    sum += (2 * i + 1) * (2 * i + 1);
    sum += (4 * i * i + 1);
    sum += (prevRight + 8 * i - 6);
    prevRight += (8 * i - 6);
    sum += (prevLeft + 8 * i - 2);
    prevLeft += (8 * i - 2);
  }
  return sum;
}
