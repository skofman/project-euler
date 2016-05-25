function sumDigits(num) {
  var arr = [];
  arr[0] = 1;
  var store = 0;
  for (var i = 0; i < num; i++) {
    for (var j = 0; j < arr.length; j++) {
      arr[j] = Number(arr[j]) * 2 + Number(store);
      store = 0;
      if (Number(arr[j] > 9)) {
        store = 1;
        arr[j] = Number(arr[j]) - 10;
      }
    }
    if (store === 1) {
      arr.push(store);
      store = 0;
    }
  }
  var sum = 0;
  for (var k = 0; k < arr.length; k++) {
    sum += Number(arr[k]);
  }
  return sum;
}
