function sumOfTwo(a, b, v) {
  let sums = {};
  for (let x of a) {
    sums[v - x] = true;
  }
  for (let j of b) {
    if (sums[j]) return true;
  }
  return false;
}
