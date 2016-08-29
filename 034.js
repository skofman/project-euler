function factorial(num) {
  var result = 1;
  for (var i = 1; i <= num; i++) {
    result *= i;
  }
  return result;
}

function digitFacts() {
  var arr = [];
  var sum = 0;
  var result = 0;
  for (var i = 10; i < 362880; i++) {
    arr = i.toString().split('');
    for (var j = 0; j < arr.length; j++) {
      sum += factorial(Number(arr[j]));
    }
    if (sum === i) {
      result += i;
      console.log(i);
    }
    sum = 0;
  }
  return result;
}
