function curiousFractions() {
  total = 1;
  var numer = 0;
  var denom = 0;
  for (var i = 10; i <= 99; i++) {
    for (var j = i + 1; j <= 99; j++) {
      numer = Number(i.toString().substr(0,1));
      denom = Number(j.toString().substr(1,1));
      if (numer / denom === i / j && i.toString().substr(1,1) === j.toString().substr(0,1)) {
        total *= numer;
        total /= denom;
      }
    }
  }
  return total;
}
