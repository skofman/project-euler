function findNumber() {
  var test = null;
  var found = true;

  for (var i = 125873; i < 10000000; i++) {
    for (var j = 2; j <= 6; j++) {
      test = i * j;
      var a = _.sortBy(i.toString().split(""));
      var b = _.sortBy(test.toString().split(""));
      for (var k = 0; k < a.length; k++) {
        if (a[k] != b[k]) {
          found = false;
          break;
        }
      }
    }

    if (found) {
      return i;
    }
    else {
      found = true;
    }

    if (i.toString().split("")[0] != "1") {
      var length = i.toString().length;
      i = Math.pow(10, length);
    }
  }
}
