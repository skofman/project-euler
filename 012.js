function triangle() {
  var arr = [1];
  var divisors;
  while(true) {
    arr.push(arr[arr.length - 1] + arr.length + 1);
    if (Math.sqrt(arr[arr.length - 1]) * Math.sqrt(arr[arr.length - 1]) === arr[arr.length - 1]) {
      divisors = 1;
      for (var i = 1; i < Math.sqrt(arr[arr.length - 1]); i++) {
        if (arr[arr.length - 1] % i === 0) {
          divisors++;
          if (divisors === 251) {
            return arr[arr.length - 1];
          }
        }
      }
    }
  }
}
