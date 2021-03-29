/* 
closest entry in 3 sorted arrays

take 3 sorted arrays and return one entry from each array
such as the interval containing these 3 entries is as small
as possible.
example (5,10,15) (3,6,9,12,15) (8,16,24) would return 15,15,16
*/

// actually this is wrong because it does not guarantee it will be
// one from each array, could be all from the same array
function minInterval(arrays) {
  console.log(arrays);
  i1 = 0;
  i2 = 0;
  i3 = 0;
  all = [];
  while (i1 < arrays[0].length || i2 < arrays[1].length || i3 < arrays[2].length) {
    a = arrays[0][i1] || 99;
    b = arrays[1][i2] || 99;
    c = arrays[2][i3] || 99;
    if (a <= b && a <= c) {
      all.push(a);
      i1++;
    } else if (b <= a && b <= c) {
      all.push(b);
      i2++;
    } else {
      all.push(c);
      i3++;
    }
  }
  bestIndex = 1;
  var bestInterval;
  console.log(all);
  for (i = 1; i < all.length - 1; i++) {
    value = Math.abs(all[i] - all[i - 1]) + Math.abs(all[i] - all[i + 1]);
    console.log(i, value);
    if (!bestInterval || value < bestInterval) {
      bestInterval = value;
      bestIndex = i;
    }
  }
  console.log(bestIndex);
  return all.slice(bestIndex - 1, bestIndex + 2);
}

console.log(
  minInterval([
    [5, 10, 15],
    [3, 6, 9, 12, 15],
    [8, 16, 24],
  ])
);
