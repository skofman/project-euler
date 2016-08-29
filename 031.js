function coins(num) {
  var arr = [100,50,20,10,5,2,1];
  var total = 1;
  var sum = 0;
  for (var i = 0; i <= num / 100; i++) {
    for (var j = 0; j <= num / 50; j++) {
      for (var k = 0; k <= num / 20; k++) {
        for (var l = 0; l <= num / 10; l++) {
          for (var m = 0; m <= num / 5; m++) {
            for (var n = 0; n <= num / 2; n++) {
              for (var o = 0; o <= num; o++) {
                sum = arr[0] * i + arr[1] * j + arr[2] * k + arr[3] * l + arr[4] * m + arr[5] * n + arr[6] * o;
                if (sum === num) {
                  total++;
                }
               }
            }
          }
        }
      }
    }
  }
  return total;
}
