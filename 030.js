function powerFive() {
  var arr;
  var sum = 0;
  var power = 0;
  for (var i = 10; i <= 1000000; i++) {
    arr = i.toString().split('');
    for (var j = 0; j < arr.length; j++) {
      power += Math.pow(Number(arr[j]), 5);
    }
    if (power === i) {
      sum += i;
    }
    power = 0;
  }
  return sum;
}
