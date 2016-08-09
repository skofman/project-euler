function cycles(num) {
  var arr = [];
  var div = "10";
  var score;
  for (var i = 1; i < num; i++) {
    arr[i - 1] = [0,0,0,0];
    if (div.length < i.toString().length) {
      div += "0";
    }
    while (arr[i - 1].length < 10000) {
      score = Math.floor(Number(div) / i);
      arr[i - 1].push(score);
      div = ((Number(div) - score * i)).toString() + "0";
    }
    div = "10";
  }
  var str;
  var newArr = [];
  for (var i = 0; i < arr.length; i++) {
    str = arr[i].join('');
    newArr[i] = checkRepeat(str);
  }
  return newArr.indexOf(Math.max(...newArr)) + 1;

}

function checkRepeat(str) {
  while(str.length > 1) {
    if (str.charAt(0) != 0 && str.charAt(1) != 0) {
      for (var i = 1; i < str.length / 2; i++) {
        if (str.substr(0, i) === str.substr(i, i) && str.charAt(0) != "0") {
          return i;
        }
      }
    }
    str = str.slice(1);
  }
  return false;
}
