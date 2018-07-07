const cube = num => {
  return num * num * num;
};

const isPermutation = (a, b) => {
  return (
    a
      .toString()
      .split('')
      .sort()
      .join('') ===
    b
      .toString()
      .split('')
      .sort()
      .join('')
  );
};

const max = 208064;

let i = 1;
let numbers = [];
while (i < max) {
  const cubic = cube(i);
  for (let j = i + 1; j < max; j++) {
    const testNumber = cube(j);
    if (testNumber.toString().length > cubic.toString().length) {
      break;
    }
    if (isPermutation(cubic, testNumber)) {
      numbers.push(testNumber);
      if (numbers.length === 4) {
        console.log(cubic);
        i = max + 1;
        break;
      }
    }
  }
  i++;
  numbers = [];
}
