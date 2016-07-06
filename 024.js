function factorial(num) {
  var result = 1;
  for (var i = 1; i <= num; i++) {
    result *= i;
  }
  return result;
}

function orderNumbers() {
  var arr = [0,1,2,3,4,5,6,7,8,9];
  var start = 1000000;
  var result = [];
  var cycle = 0;
  var cycles = 0;
  while (arr.length > 0) {
    cycle = factorial(arr.length - 1);
    cycles = Math.floor(start / cycle);
    result.push(arr[cycles]);
    start = Math.abs(start - cycles * cycle);
    arr.splice(cycles,1);
  }
  return result;
}
