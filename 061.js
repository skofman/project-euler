const triangle = [];
const square = [];
const pentagonal = [];
const hexagonal = [];
const heptagonal = [];
const octagonal = [];

for (let i = 1; i <= 150; i++) {
  triangle.push((i * (i + 1)) / 2);
  square.push(i * i);
  pentagonal.push((i * (3 * i - 1)) / 2);
  hexagonal.push(i * (2 * i - 1));
  heptagonal.push((i * (5 * i - 3)) / 2);
  octagonal.push(i * (3 * i - 2));
}

const candidates = [];

for (let i = 0; i < triangle.length; i++) {
  if (triangle[i] >= 1000 && triangle[i] < 10000 && !candidates.includes(triangle[i])) {
    candidates.push(triangle[i]);
  }
}

for (let i = 0; i < square.length; i++) {
  if (square[i] >= 1000 && square[i] < 10000 && !candidates.includes(square[i])) {
    candidates.push(square[i]);
  }
}

for (let i = 0; i < pentagonal.length; i++) {
  if (pentagonal[i] >= 1000 && pentagonal[i] < 10000 && !candidates.includes(pentagonal[i])) {
    candidates.push(pentagonal[i]);
  }
}

for (let i = 0; i < hexagonal.length; i++) {
  if (hexagonal[i] >= 1000 && hexagonal[i] < 10000 && !candidates.includes(hexagonal[i])) {
    candidates.push(hexagonal[i]);
  }
}

for (let i = 0; i < heptagonal.length; i++) {
  if (heptagonal[i] >= 1000 && heptagonal[i] < 10000 && !candidates.includes(heptagonal[i])) {
    candidates.push(heptagonal[i]);
  }
}

for (let i = 0; i < octagonal.length; i++) {
  if (octagonal[i] >= 1000 && octagonal[i] < 10000 && !candidates.includes(octagonal[i])) {
    candidates.push(octagonal[i]);
  }
}

let chains = [];
for (let i = 0; i < candidates.length; i++) {
  chains.push([candidates[i]]);
}

let tempChains = [];
for (let i = 0; i < chains.length; i++) {
  const split = chains[i][0].toString().split('');
  const end = split[2] + split[3];
  for (let j = 0; j < candidates.length; j++) {
    const split = candidates[j].toString().split('');
    const start = split[0] + split[1];
    if (start === end) {
      tempChains.push([...chains[i], candidates[j]]);
    }
  }
}

chains = tempChains;
tempChains = [];

for (let i = 0; i < chains.length; i++) {
  const split = chains[i][1].toString().split('');
  const end = split[2] + split[3];
  for (let j = 0; j < candidates.length; j++) {
    const split = candidates[j].toString().split('');
    const start = split[0] + split[1];
    if (start === end) {
      tempChains.push([...chains[i], candidates[j]]);
    }
  }
}

chains = tempChains;
tempChains = [];

for (let i = 0; i < chains.length; i++) {
  const split = chains[i][2].toString().split('');
  const end = split[2] + split[3];
  for (let j = 0; j < candidates.length; j++) {
    const split = candidates[j].toString().split('');
    const start = split[0] + split[1];
    if (start === end) {
      tempChains.push([...chains[i], candidates[j]]);
    }
  }
}

chains = tempChains;
tempChains = [];

for (let i = 0; i < chains.length; i++) {
  const split = chains[i][3].toString().split('');
  const end = split[2] + split[3];
  for (let j = 0; j < candidates.length; j++) {
    const split = candidates[j].toString().split('');
    const start = split[0] + split[1];
    if (start === end) {
      tempChains.push([...chains[i], candidates[j]]);
    }
  }
}

chains = tempChains;
tempChains = [];

for (let i = 0; i < chains.length; i++) {
  const split = chains[i][4].toString().split('');
  const end = split[2] + split[3];
  const splitFirst = chains[i][0].toString().split('');
  const startFirst = splitFirst[0] + splitFirst[1];
  for (let j = 0; j < candidates.length; j++) {
    const split = candidates[j].toString().split('');
    const start = split[0] + split[1];
    const endLast = split[2] + split[3];
    if (start === end && endLast === startFirst) {
      tempChains.push([...chains[i], candidates[j]]);
    }
  }
}

chains = tempChains;
tempChains = [];

for (let i = 0; i < chains.length; i++) {
  let obj = {
    triangle: [],
    square: [],
    pentagonal: [],
    hexagonal: [],
    heptagonal: [],
    octagonal: []
  };
  for (let j = 0; j < chains[i].length; j++) {
    if (triangle.includes(chains[i][j])) {
      obj.triangle.push(j);
    }
    if (square.includes(chains[i][j])) {
      obj.square.push(j);
    }
    if (pentagonal.includes(chains[i][j])) {
      obj.pentagonal.push(j);
    }
    if (hexagonal.includes(chains[i][j])) {
      obj.hexagonal.push(j);
    }
    if (heptagonal.includes(chains[i][j])) {
      obj.heptagonal.push(j);
    }
    if (octagonal.includes(chains[i][j])) {
      obj.octagonal.push(j);
    }
  }
  let isEmpty = false;
  for (const key in obj) {
    if (obj[key].length === 0) {
      isEmpty = true;
    }
  }
  const index = [];
  if (!isEmpty) {
    for (const key in obj) {
      if (obj[key].length === 1) {
        if (!index.includes(obj[key][0])) {
          index.push(obj[key][0]);
        } else {
          isEmpty = true;
        }
      }
    }
  }
  if (!isEmpty) {
    tempChains.push(chains[i]);
  }
}

chains = tempChains;
tempChains = [];
const sums = [];

for (let i = 0; i < chains.length; i++) {
  let sum = 0;
  for (let j = 0; j < chains[i].length; j++) {
    sum += chains[i][j];
  }
  if (!sums.includes(sum)) {
    sums.push(sum);
  }
}

console.log(sums);
