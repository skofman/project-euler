function fibThree() {
  var arr = [];
  arr[0] = [1];
  arr[1] = [1];
  var store = 0;
  var i = 2;
  var sum = 0;
  while (arr[arr.length - 1].length < 1000) {
    arr[i] = [];
    if (arr[i - 1].length > arr[i - 2].length) {
      arr[i - 2].push(0);
    }
    for (var j = 0; j < arr[i - 1].length; j++) {
      sum += arr[i - 1][j] + arr[i - 2][j];
      if (sum <= 9) {
        arr[i].push(sum);
        sum = 0;
      }
      else {
        arr[i].push(Number(sum.toString().substr(-1,1)));
        if (j === arr[i - 1].length - 1) {
          arr[i].push(1);
          sum = 0;
        }
        else {
          sum = 1;
        }
      }
    }
    i++;
  }
  return arr.length;
}
