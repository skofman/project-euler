function sumFacts(num) {
  var store = 0;
  var result = [1];
  for (var i = 2; i <= num; i++) {
    for (var j = 0; j < result.length; j++) {
      result[j] *= i;
    }
    for (var k = 0; k < result.length; k++) {
      if (result[k] > 9) {
        store = result[k].toString().substr(0,result[k].toString().length - 1).split('').reverse().join('');
        result[k] = Number(result[k].toString().substr(-1,1));
      for (var l = 0; l < store.length; l++) {
        if ((l + k + 2) > result.length) {
          result.push(Number(store.charAt(l)));
        }
        else {
          result[k + l + 1] += Number(store.charAt(l));
        }
      }
    }
    }
  }
  var sum = 0;
  for (var m = 0; m < result.length; m++) {
    sum += result[m];
  }
  return sum;
}
