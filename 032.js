function pandigital() {
  var sum = 0;
  var multiply = 0;
  var str = "";
  var arr = [];
  var result = [];
  for  (var i = 1; i < 5000; i++) {
    for (var j = i; j < 5000; j++) {
      multiply = i * j;
      str = i.toString() + j.toString() + multiply.toString();
      arr = _.union(str.split(''));
      if (arr.length === 9 && str.length === 9 && arr.indexOf("0") === -1) {
        result.push(multiply);
        console.log(i + " x " + j + " = " + multiply);
      }
    }
  }
  result = _.union(result);
  for (var i = 0; i < result.length; i++) {
    sum += result[i];
  }
  return sum;
}
