function champernowne() {
  var result = 1;
  var d = 1;
  var counter = 0;
  var str = "";
  var i = 1;
  while (d < 10000000) {
    str = i.toString();
    counter += str.length;
    if (counter >= d) {
      var slice = d - counter - 1;
      result *= Number(str.substr(slice, 1));
      d *= 10;
    }
    i++;
  }
  return result;
}
