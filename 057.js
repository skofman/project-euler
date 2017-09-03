function add(str1, str2) {
  let arr1 = str1.split('').reverse();
  let arr2 = str2.split('').reverse();
  let result = [];

  if (arr1.length > arr2.length) {
    for (let i = arr2.length; i < arr1.length; i++) {
      arr2.push('0');
    }
  }
  else if (arr2.length > arr1.length) {
    for (let i = arr1.length; i< arr2.length; i++) {
      arr1.push('0');
    }
  }

  let adder = 0;
  for (let i = 0; i < arr1.length; i++) {
    let temp = Number(arr1[i]) + Number(arr2[i]) + adder;
    if (temp > 9) {
      result.push(temp.toString().split('')[1]);
      adder = 1;
    }
    else {
      result.push(temp.toString());
      adder = 0;
    }
  }
  if (adder === 1) {
    result.push('1');
  }

  return result.reverse().join('');
}

function multiply(str1, str2) {
  let arr1 = str1.split('').reverse();
  let arr2 = str2.split('').reverse();
  let result = [];

  if (arr1.length > arr2.length) {
    for (let i = arr2.length; i < arr1.length; i++) {
      arr2.push('0');
    }
  }
  else if (arr2.length > arr1.length) {
    for (let i = arr1.length; i< arr2.length; i++) {
      arr1.push('0');
    }
  }

  let tempArr = [];
  let adder = 0;
  for (let i = 0; i < arr1.length; i++) {
    for (let k = 0; k < i; k++) {
      tempArr.push('0');
    }
    for (let j = 0; j < arr2.length; j++) {
      let temp = Number(arr2[j]) * Number(arr1[i]) + adder;
      if (temp > 9) {
        tempArr.push(temp.toString().split('')[1]);
        adder = Number(temp.toString().split('')[0]);
      }
      else {
        tempArr.push(temp.toString());
        adder = 0;
      }
    }
    if (adder > 0) {
      tempArr.push(adder);
      adder = 0;
    }

    result = add(result.reverse().join(''), tempArr.reverse().join('')).split('').reverse();
    tempArr = [];
  }

  result = result.reverse().join('');
  for (let i = 0; i < result.length; i++) {
    if (result[i] !== '0') {
      return result.substr(i);
    }
  }
}

function squareRoots() {
  let nom = ["1", "3"];
  let dem = ["1", "2"];

  let prevNom = "1";
  let currentNom = "3";
  let resultNom;
  let prevDem = "1";
  let currentDem = "2";
  let resultDem;
  for (let i = 0; i < 1000; i++) {
    resultNom = add(multiply(currentNom,"2"), prevNom);
    prevNom = currentNom;
    currentNom = resultNom;
    nom.push(resultNom);
    resultDem = add(multiply(currentDem, "2"), prevDem);
    prevDem = currentDem;
    currentDem = resultDem;
    dem.push(resultDem);
  }

  let result = 0;
  for (let i = 0; i < nom.length; i++) {
    if (nom[i].length > dem[i].length) {
      result++;
    }
  }

  return result;
}

console.log(squareRoots());