function pyth() {
  var c = 0;
  for (var a = 1; a < 1000; a++) {
    for (var b = a + 1; b < 1000; b++) {
      c = 1000 - b - a;
      if (c * c === a * a + b * b) {
        return a * b * c;
      }
    }
  }
}
