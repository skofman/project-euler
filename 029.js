function powers(num) {
  var arr = [];
  for (var a = 2; a <= num; a++) {
    for (var b = 2; b <= num; b++) {
      arr.push(Math.pow(a, b));
    }
  }
  return _.union(arr).length;
}
