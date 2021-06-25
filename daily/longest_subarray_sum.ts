function findLongestSubarrayBySum(s: number, a: number[]): number[] {
  let start = 0;
  let end = 0;
  let sum = 0;
  let result: number[] = [-1];
  a.forEach((n, i) => {
    console.log(start, end, sum, n);
    if (sum === s) {
      if (n === 0 && i !== a.length - 1) {
        end = i;
        return;
      }
      console.log('match at', start, end);
      if (result.length === 1 || result[1] - result[0] > end - start) result = [start + 1, end + 1];
      return;
    }
    if (sum < s) {
      sum += n;
      end = i;
      return;
    }
    // sum > n
    while (sum > s && start <= i) {
      sum -= a[start];
      console.log('- ', sum, start, sum);
      start++;
      if (sum === s) {
        if (result.length === 1 || result[1] - result[0] > end - start) result = [start + 1, end + 1];
      }
    }
  });
  console.log('result', result);
  return result;
}

findLongestSubarrayBySum(12, [1, 2, 3, 7, 5, 3]); // [2, 4]
findLongestSubarrayBySum(
  // [20, 37]
  1562,
  [
    28, 68, 142, 130, 41, 14, 175, 2, 78, 16, 84, 14, 193, 25, 2, 193, 160, 71, 29, 28, 85, 76, 187, 99, 171, 88, 48, 5,
    104, 22, 64, 107, 164, 11, 172, 90, 41, 165, 143, 20, 114, 192, 105, 19, 33, 151, 6, 176, 140, 104, 23, 99, 48, 185,
    49, 172, 65,
  ]
);

findLongestSubarrayBySum(
  // [45, 62]
  1433,
  [
    137, 181, 76, 108, 185, 75, 120, 191, 77, 42, 152, 130, 91, 175, 73, 192, 190, 188, 91, 40, 94, 54, 64, 82, 104, 6,
    177, 80, 96, 140, 69, 199, 84, 40, 116, 22, 194, 27, 123, 39, 29, 182, 200, 179, 92, 24, 144, 35, 44, 150, 103, 108,
    103, 142, 88, 147, 92, 38, 14, 1, 17, 91, 163, 136, 127, 11, 78, 183, 61, 190, 106, 75, 64, 123, 196, 167, 161, 39,
    189, 12, 46, 68, 26, 68, 190, 143, 64, 148, 103, 18, 100, 24, 27, 4,
  ]
);
