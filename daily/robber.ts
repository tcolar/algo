/*
You are planning to rob houses on a specific street, and you know that every house on the street has a certain amount of money hidden. The only thing stopping you from robbing all of them in one night is that adjacent houses on the street have a connected security system. The system will automatically trigger an alarm if two adjacent houses are broken into on the same night.

Given a list of non-negative integers nums representing the amount of money hidden in each house, determine the maximum amount of money you can rob in one night without triggering an alarm.

Example

For nums = [1, 1, 1], the output should be
houseRobber(nums) = 2.

The optimal way to get the most money in one night is to rob the first and the third houses for a total of 2.
*/

function houseRobber(nums: number[]): number {
  if (!nums.length) return 0;
  if (nums.length === 1) return nums[0];
  if (nums.length === 2) return Math.max(...nums);
  let pre1 = nums[0];
  let pre2 = nums[0] > nums[1] ? nums[0] : nums[1];
  for (let i = 2; i < nums.length; i++) {
    const tmp = pre1 + nums[i] > pre2 ? pre1 + nums[i] : pre2;
    console.log(nums[i], pre1, pre2, '-', tmp);
    pre1 = pre2;
    pre2 = tmp;
  }
  return pre2;
}

const input = [3, 6, 9, 2, 6, 5, 1, 10];
console.log(input);
console.log(houseRobber(input));
