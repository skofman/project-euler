function bigPandigital() {
  var str = "";
  var multi = 2;
  var max = 0;
  var result = 0;
  for (i = 1; i < 10000; i++) {
    if (i.toString().substr(0,1) === "9") {
      result = i;
      str = result.toString();
      while (result.toString().length < 9) {
        str = (i * multi).toString();
        result = Number(result.toString() + str);
        multi++;
      }
      multi = 2;
      if (result.toString().length === 9 && _.uniq(result.toString().split('')).length === 9 && result.toString().indexOf('0') === -1) {
        if (result > max) {
          max = result;
        }
      }
    }
  }
  return max;
}
