function numberPalindrome(num) {
  var rev = Number(num.toString().split('').reverse().join(''));
  if (num === rev) {
    return true;
  }
  else {
    return false;
  }
}

function product(num) {
  var j = 0;
  var pal = 100 * 100;
  for (var i = num; i > 0; i--) {
    j = i;
    if (i * j > pal) {
      while (i * j > pal) {
        if (numberPalindrome(i * j)) {
          pal = i * j;
        }
        j--;
      }
    }
    else {
      return pal;
    }
  }
}
