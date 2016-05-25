function convert(num) {
  var arr = num.toString().split('');
  var string = "";
  if (arr.length === 1) {
    switch(arr[0]) {
      case "1":
      string += "one";
      break;
      case "2":
      string += "two";
      break;
      case "3":
      string += "three";
      break;
      case "4":
      string += "four";
      break;
      case "5":
      string += "five";
      break;
      case "6":
      string += "six";
      break;
      case "7":
      string += "seven";
      break;
      case "8":
      string += "eight";
      break;
      case "9":
      string += "nine";
      break;
    }
  }
  if (arr.length === 2) {
    if (arr[0] === "1") {
      switch(arr[1]) {
        case "0":
        string += "ten";
        break;
        case "1":
        string += "eleven";
        break;
        case "2":
        string += "twelve";
        break;
        case "3":
        string += "thirteen";
        break;
        case "4":
        string += "fourteen";
        break;
        case "5":
        string += "fifteen";
        break;
        case "6":
        string += "sixteen";
        break;
        case "7":
        string += "seventeen";
        break;
        case "8":
        string += "eighteen";
        break;
        case "9":
        string += "nineteen";
        break;
      }
    }
    else {
      switch(arr[1]) {
        case "1":
        string += "one";
        break;
        case "2":
        string += "two";
        break;
        case "3":
        string += "three";
        break;
        case "4":
        string += "four";
        break;
        case "5":
        string += "five";
        break;
        case "6":
        string += "six";
        break;
        case "7":
        string += "seven";
        break;
        case "8":
        string += "eight";
        break;
        case "9":
        string += "nine";
        break;
      }
      switch(arr[0]) {
        case "2":
        string += "twenty";
        break;
        case "3":
        string += "thirty";
        break;
        case "4":
        string += "forty";
        break;
        case "5":
        string += "fifty";
        break;
        case "6":
        string += "sixty";
        break;
        case "7":
        string += "seventy";
        break;
        case "8":
        string += "eighty";
        break;
        case "9":
        string += "ninety";
        break;
      }
    }
  }
  if (arr.length === 3) {
    switch(arr[0]) {
      case "1":
      string += "onehundred";
      break;
      case "2":
      string += "twohundred";
      break;
      case "3":
      string += "threehundred";
      break;
      case "4":
      string += "fourhundred";
      break;
      case "5":
      string += "fivehundred";
      break;
      case "6":
      string += "sixhundred";
      break;
      case "7":
      string += "sevenhundred";
      break;
      case "8":
      string += "eighthundred";
      break;
      case "9":
      string += "ninehundred";
      break;
    }
    switch(arr[1]) {
      case "2":
      string += "andtwenty";
      break;
      case "3":
      string += "andthirty";
      break;
      case "4":
      string += "andforty";
      break;
      case "5":
      string += "andfifty";
      break;
      case "6":
      string += "andsixty";
      break;
      case "7":
      string += "andseventy";
      break;
      case "8":
      string += "andeighty";
      break;
      case "9":
      string += "andninety";
      break;
    }
    if (arr[1] === "1") {
      switch(arr[2]) {
        case "0":
        string += "andten";
        break;
        case "1":
        string += "andeleven";
        break;
        case "2":
        string += "andtwelve";
        break;
        case "3":
        string += "andthirteen";
        break;
        case "4":
        string += "andfourteen";
        break;
        case "5":
        string += "andfifteen";
        break;
        case "6":
        string += "andsixteen";
        break;
        case "7":
        string += "andseventeen";
        break;
        case "8":
        string += "andeighteen";
        break;
        case "9":
        string += "andnineteen";
        break;
      }
    }
    else if (arr[1] === "0") {
      switch(arr[2]) {
        case "1":
        string += "andone";
        break;
        case "2":
        string += "andtwo";
        break;
        case "3":
        string += "andthree";
        break;
        case "4":
        string += "andfour";
        break;
        case "5":
        string += "andfive";
        break;
        case "6":
        string += "andsix";
        break;
        case "7":
        string += "andseven";
        break;
        case "8":
        string += "andeight";
        break;
        case "9":
        string += "andnine";
        break;
      }
    }
    else {
      switch(arr[2]) {
        case "1":
        string += "one";
        break;
        case "2":
        string += "two";
        break;
        case "3":
        string += "three";
        break;
        case "4":
        string += "four";
        break;
        case "5":
        string += "five";
        break;
        case "6":
        string += "six";
        break;
        case "7":
        string += "seven";
        break;
        case "8":
        string += "eight";
        break;
        case "9":
        string += "nine";
        break;
      }
    }
  }
console.log(string);
return string.length;
}

function letterCount(num) {
  sum = 0;
  for (var i = 1; i < num; i++) {
    sum += convert(i);
  }
  sum += "onethousand".length;
  return sum;
}
