function doublePalindrome(num) {
  var sum = 0;
  var str = "";
  var binary = "";
  for (var i = 1; i < num; i += 2) {
    str = i.toString();
    binary = makeBinary(i);
    if (str === str.split('').reverse().join('') && binary === binary.split('').reverse().join('')) {
      sum += i;
    }
  }
  return sum;
}

function makeBinary(num) {
  var str = "";
  var i = 0;
  while (num > 0) {
    if (num % 2 === 0) {
      str += "0";
    }
    else {
      str += "1";
    }
    num = Math.floor(num / 2);
  }
  return str;
}
