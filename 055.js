function findLychrel() {
  let count = 0;
  let palindrome = false;
  for (let i = 5; i <= 10000; i++) {
    let number = i.toString();
    for (let j = 0; j < 50; j++) {
      number = add(number);
      if (isPalindrome(number)) {
        palindrome = true;
      }
    }
    if (!palindrome) {
      count++;
    }
    palindrome = false;
  }
  console.log(count);
}

function add(str) {
  const arr1 = str.split('');
  const arr2 = str.split('').reverse();
  let result = [];
  let adder = 0;
  for (let i = 0; i < arr1.length; i++) {
    const temp = Number(arr1[i]) + Number(arr2[i]) + adder;
    if (temp > 9) {
      result.push(temp.toString().split('')[1]);
      adder = 1;
    }
    else {
      result.push(temp.toString());
      adder = 0;
    }
  }
  if (adder  === 1) {
    result.push('1');
  }
  return result.reverse().join('');
}

function isPalindrome(str) {
  const arr = str.split('');
  return arr.join('') === arr.reverse().join('');
}

findLychrel();