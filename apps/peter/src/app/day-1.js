const input = [
  21,
  1814,
  1826,
  1667,
  1334,
  1504,
  1420,
  1164,
  1414,
  1934,
  1823,
  1507,
  1195,
  21,
  1752,
  1472,
  1196,
  1558,
  1322,
  1927,
  1556,
  1922,
  277,
  1828,
  1883,
  1280,
  1947,
  1231,
  1915,
  1235,
  1961,
  1494,
  1324,
  2009,
  1367,
  1545,
  1736,
  1575,
  1214,
  1704,
  1833,
  1663,
  1474,
  1894,
  1754,
  1564,
  1321,
  1119,
  1975,
  1987,
  1873,
  1834,
  1686,
  1574,
  1505,
  1656,
  1688,
  1896,
  1982,
  1554,
  1990,
  1902,
  1859,
  1293,
  1739,
  1282,
  1889,
  1981,
  1283,
  1687,
  1220,
  1443,
  1409,
  1252,
  1506,
  1742,
  1319,
  1882,
  951,
  1849,
];

function dayOneFirstChallenge(input, match = 2020) {
  input.sort((a, b) => a - b);
  for (curr of input) {
    const rest = match - curr;
    if (input.indexOf(rest) !== -1) {
      return curr * rest;
    }
  }
}

dayOneFirstChallenge(input);

function dayOneSecondChallenge(input, match = 2020) {
  input.sort((a, b) => a - b);
  for (let i = 0; i < input.length; i++) {
    const curr = input[i];
    const rest = match - curr;
    const searchIn = [...input].splice(i).filter((val) => val < rest);
    const result = curr * dayOneFirstChallenge(searchIn, rest);

    if (result) {
      return result;
    }
  }
}

dayOneSecondChallenge(input);
