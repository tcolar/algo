/*
ou have an array of integers nums and an array queries, where queries[i] is a pair of indices (0-based). Find the sum of the elements in nums from the indices at queries[i][0] to queries[i][1] (inclusive) for each query, then add all of the sums for all the queries together. Return that number modulo 109 + 7.

Example

For nums = [3, 0, -2, 6, -3, 2] and queries = [[0, 2], [2, 5], [0, 5]], the output should be
sumInRange(nums, queries) = 10.

The array of results for queries is [1, 3, 6], so the answer is 1 + 3 + 6 = 10.
*/

function sumInRange(nums: number[], queries: number[][]): number {
  let start = new Date();
  let sum = 0;
  // original approach, was too slow, this new approach works fast enough:
  // keep track of multiplication factor changes (by index),
  // as in, how many ranges are overlapping at a given index,
  // only track when that changes
  let changes = {};
  const mod = 10 ** 9 + 7;
  for (const q of queries) {
    // increase the factor at range start
    changes[q[0]] = (changes[q[0]] || 0) + 1;
    // decrease after range (since range is inclusive)
    changes[q[1] + 1] = (changes[q[1] + 1] || 0) - 1;
  }
  let factor = 0;
  for (let i = 0; i != nums.length; i++) {
    if (i in changes) {
      factor += changes[i];
    }
    if (factor === 0) continue;
    // js does not deal well with modulo of negative numbers, see here for info:
    // https://web.archive.org/web/20090717035140if_/javascript.about.com/od/problemsolving/a/modulobug.htm
    sum = (((sum + nums[i] * factor) % mod) + mod) % mod;
  }
  console.log('got', sum, 'in', new Date().getTime() - start.getTime(), 'ms');
  return sum;
}

sumInRange(
  [3, 0, -2, 6, -3, 2],
  [
    [0, 2],
    [2, 5],
    [0, 5],
  ]
);

sumInRange([-1000], [[0, 0]]);

const testNums: number[] = [];
const testQueries: number[][] = [];
for (let i = 0; i < 100000; i++) {
  testNums.push(Math.floor(Math.random() * 2000) - 1000);
}
for (let i = 0; i < 100000; i++) {
  const start = Math.floor(Math.random() * 100000);
  const end = Math.floor(Math.random() * 100000);
  if (start < end) {
    testQueries.push([start, end]);
  } else {
    testQueries.push([end, start]);
  }
}
sumInRange(testNums, testQueries);
