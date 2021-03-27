/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

function findSum(list1, sumToFind) {
  const map = {};
  for (num of list1) {
    if (sumToFind - num in map) {
      console.log(`found sum with ${num} and ${sumToFind - num}`);
      return;
    }
    if (!(num in map)) {
      map[num] = true;
    }
  }
  console.log(`sum not found in ${JSON.stringify(map)}`);
}

findSum([10, 15, 3, -7, 7, 4, -3], 15);
