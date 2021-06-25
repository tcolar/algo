function arrayMaxConsecutiveSum2(a: number[]): number {
  let curSum: number = null;
  let maxSum: number = null;
  a.forEach((value) => {
    if (curSum + value > value) curSum += value;
    else curSum = value;
    if (!maxSum || curSum > maxSum) {
      maxSum = curSum;
    }
    //console.log(curSum, maxSum)
  });
  return maxSum;
}
