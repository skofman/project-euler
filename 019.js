function sundays() {
  var sunday = 0;
  var days = -6;
  for (var i = 1901; i <= 2000; i++) {
    for (var j = 1; j <= 12; j++) {
      switch(j) {
        case 1:
        case 3:
        case 5:
        case 7:
        case 8:
        case 10:
        case 12:
          for (var k = 1; k <= 31; k++) {
            days++;
            if (k === 1 && days % 7 === 0) {
              sunday++;
            }
          }
          break;
        case 4:
        case 6:
        case 9:
        case 11:
          for (var k = 1; k <= 30; k++) {
            days++;
            if (k === 1 && days % 7 === 0) {
              sunday++;
            }
          }
          break;
        case 2:
          if (i % 4 === 0) {
            for (var k = 1; k <= 29; k++) {
              days++;
              if (k === 1 && days % 7 === 0) {
                sunday++;
              }
            }
          }
          else {
            for (var k = 1; k <= 28; k++) {
              days++;
              if (k === 1 && days % 7 === 0) {
                sunday++;
              }
            }
          }
        }
      }
    }
    return sunday;
  }
