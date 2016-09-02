function findTriangles() {
  var p;
  var max = 0;
  var c;
  var counter = 0;
  for (var i = 3; i <= 1000; i++) {
    for (var a = 1; a <= 1000; a++) {
      for (var b = a + 1; b <= 1000; b++) {
        c = i - b - a;
        if (a * a + b * b === c * c && c > 0) {
          counter++;
        }
      }
    }
    if (counter > max) {
      max = counter;
      p = i;
    }
    counter = 0;
  }
  return p;
}
