const fs = require('fs');

function pokerHands() {
  fs.readFile('./p054_poker.txt', 'utf8', (err, data) => {
    const hands = data.split('\n');
    let player1 = [];
    let player2 = [];
    for (let i = 0; i < hands.length - 1; i++) {
      const cards = hands[i].split(' ');
      player1.push([cards[0], cards[1], cards[2], cards[3], cards[4]]);
      player2.push([cards[5], cards[6], cards[7], cards[8], cards[9]])
    }
    // console.log(player1);
    // console.log(player2);
    let results1 = [];
    let results2 = [];
    for (let i = 0; i < player1.length; i++) {
      results1.push(evaluateHand(player1[i]));
      results2.push(evaluateHand(player2[i]));
    }
    // console.log(results1);
    // console.log(results2);
    let win1 = 0;
    let win2 = 0;
    for (let i = 0; i < results1.length; i++) {
      if (results1[i].value > results2[i].value) {
        win1++;
      }
      else if (results2[i].value > results1[i].value) {
        win2++;
      }
      else if (results1[i].value === results2[i].value) {
        for (let j = 0; j < results1[i].kickers.length; j++) {
          if (results1[i].kickers[j] > results2[i].kickers[j]) {
            win1++;
            break;
          }
          else if (results2[i].kickers[j] > results1[i].kickers[j]) {
            win2++;
            break;
          }
        }
      }
    }
    console.log(win1);
    console.log(win2);
  });
}

function evaluateHand(hand) {
  let result;
  result = royalFlush(hand);
  if (result.value === 9) {
    return result;
  }
  result = straightFlush(hand);
  if (result.value === 8) {
    return result;
  }
  result = fourOfAKind(hand);
  if (result.value === 7) {
    return result;
  }
  result = fullHouse(hand);
  if (result.value === 6) {
    return result;
  }
  result = flush(hand);
  if (result.value === 5) {
    return result;
  }
  result = straight(hand);
  if (result.value === 4) {
    return result;
  }
  result = trips(hand);
  if (result.value === 3) {
    return result;
  }
  result = twoPair(hand);
  if (result.value === 2) {
    return result;
  }
  result = pair(hand);
  if (result.value === 1) {
    return result;
  }
  return highCard(hand);
}

function royalFlush(hand) {
  const colors = 'DHSC';
  for (let i = 0; i < colors.length; i++) {
    if (hand.includes(`A${colors[i]}`) && hand.includes(`K${colors[i]}`) && hand.includes(`Q${colors[i]}`) && hand.includes(`J${colors[i]}`) && hand.includes(`${colors[i]}`)) {
      return {
        value: 9,
        kickers: [],
        name: "Royal Flush"
      }
    }
  }
  return {
    value: 0
  }
}

function straightFlush(hand) {
  const ranks = '23456789TJQKA';
  for (let i = 0; i < hand.length; i++) {
    const split = hand[i].split("");
    const index = ranks.indexOf(split[0]);
    if (index > 3) {
      if (hand.includes(`${ranks.charAt(index - 1)}${split[1]}`) && hand.includes(`${ranks.charAt(index - 2)}${split[1]}`) && hand.includes(`${ranks.charAt(index - 3)}${split[1]}`) && hand.includes(`${ranks.charAt(index - 4)}${split[1]}`)) {
        return {
          value: 8,
          kickers: [index],
          name: `Straight Flush ${split[0]}`
        }
      }
    }
    else if (index === 3) {
      if (hand.includes(`${ranks.charAt(index - 1)}${split[1]}`) && hand.includes(`${ranks.charAt(index - 2)}${split[1]}`) && hand.includes(`${ranks.charAt(index - 3)}${split[1]}`) && hand.includes(`A${split[1]}`)) {
        return {
          value: 8,
          kickers: [index],
          name: `Straight Flush ${split[0]}`
        }
      }
    }
  }
  return {
    value: 0
  }
}

function fourOfAKind(hand) {
  const ranks = "23456789TJQKA";
  let values = "";
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split("")[0];
  }
  for (let i = 0; i < ranks.length; i++) {
    const reg = new RegExp(ranks.charAt(i), 'g');
    const result = values.match(reg);
    if (result !== null && result.length === 4) {
      return {
        value: 7,
        kickers: [i]
      }
    }
  }
  return {
    value: 0
  }
}

function fullHouse(hand) {
  let values = "";
  const ranks = "23456789TJQKA";
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split("")[0];
  }
  for (let i = 12; i >= 0; i--) {
    const reg1 = new RegExp(ranks[i], 'g');
    const threes = values.match(reg1);
    if (threes !== null && threes.length === 3) {
      for (let j = 12; j >= 0; j--) {
        const reg2 = new RegExp(ranks[j], 'g');
        const twos = values.match(reg2);
        if (twos !== null && twos.length === 2) {
          // console.log('full');
          // console.log(hand);
          return {
            value: 6,
            kickers: [ranks[i], ranks[j]]
          }
        }
      }
    }
  }
  return {
    value: 0
  }
}

function flush(hand) {
  const colors = "DHSC";
  const ranks = "23456789TJQKA";
  let values = "";
  let kickers = [];
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split('')[1];
    kickers.push(hand[i].split('')[0]);
  }
  for (let i = 0; i < colors.length; i++) {
    const reg = new RegExp(colors[i], 'g');
    const result = values.match(reg);
    if (result !== null && result.length === 5) {
      let resultKickers = [];
      for (let j = 12; j >= 0; j--) {
        if (kickers.includes(ranks.charAt(j))) {
          resultKickers.push(j);
        }
      }
      // console.log('flush');
      // console.log(hand);
      return {
        value: 5,
        kickers: resultKickers
      }
    }
  }
  return {
    value: 0
  }
}

function straight(hand) {
  const ranks = "23456789TJQKA";
  let values = "";
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split('')[0];
  }
  for (let i = 12; i >= 4; i--) {
    if (values.includes(ranks.charAt(i)) && values.includes(ranks.charAt(i - 1)) && values.includes(ranks.charAt(i - 2)) && values.includes(ranks.charAt(i - 3)) && values.includes(ranks.charAt(i - 4))) {
      // console.log('straight');
      // console.log(hand);
      return {
        value: 4,
        kickers: [i]
      }
    }
  }
  if (values.includes(ranks.charAt(3)) && values.includes(ranks.charAt(2)) && values.includes(ranks.charAt(1)) && values.includes(ranks.charAt(0)) && values.includes(ranks.charAt(12))) {
    // console.log('straight');
    // console.log(hand);
    return {
      value: 4,
      kickers: [3]
    }
  }
  return {
    value: 0
  }
}

function trips(hand) {
  let values = "";
  const ranks = "23456789TJQKA";
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split('')[0];
  }
  for (let i = 12; i >= 0; i--) {
    const reg = new RegExp(ranks[i], 'g');
    const result = values.match(reg);
    if (result !== null && result.length === 3) {
      let kickers = [i];
      for (let j = 12; j >= 0; j--) {
        if (values.includes(ranks[j])) {
          kickers.push(j);
        }
      }
      // console.log('trips');
      // console.log(hand);
      return {
        value: 3,
        kickers: kickers
      }
    }
  }
  return {
    value: 0
  }
}

function twoPair(hand) {
  const ranks = "23456789TJQKA";
  let values = "";
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split('')[0];
  }
  for (let i = 12; i >= 0; i--) {
    const reg1 = new RegExp(ranks[i], 'g');
    const result1 = values.match(reg1);
    if (result1 !== null && result1.length === 2) {
      for (let j = i - 1; j >= 0; j--) {
        const reg2 = new RegExp(ranks[j], 'g');
        const result2 = values.match(reg2);
        if (result2 !== null && result2.length === 2) {
          let kickers = [i, j];
          for (let k = 12; k >= 0; k--) {
            if (values.includes(ranks[k]) && k !== i && k !== j) {
              kickers.push(k);
              return {
                value: 2,
                kickers: kickers
              }
            }
          }
        }
      }
    }
  }
  return {
    value: 0
  }
}

function pair(hand) {
  const ranks = "23456789TJQKA";
  let values = '';
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split('')[0];
  }
  for (let i = 12; i >= 0; i--) {
    const reg = new RegExp(ranks[i], 'g');
    const result = values.match(reg);
    if (result !== null && result.length === 2) {
      let kickers = [i];
      for (let j = 12; j >= 0; j--) {
        if (values.includes(ranks[j]) && j !== i) {
          kickers.push(j);
        }
      }
      return {
        value: 1,
        kickers: kickers
      }
    }
  }
  return { value: 0 }
}

function highCard(hand) {
  const ranks = "23456789TJQKA";
  let values = "";
  for (let i = 0; i < hand.length; i++) {
    values += hand[i].split('')[0];
  }
  let kickers = [];
  for (let i = 12; i >= 0; i--) {
    if (values.includes(ranks[i])) {
      kickers.push(i);
    }
  }
  return {
    value: 0,
    kickers: kickers
  }
}

pokerHands();