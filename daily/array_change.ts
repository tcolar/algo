function arrayChange(inputArray: number[]): number {
  let moves = 0;
  let nextNum: number = inputArray[0] + 1;
  for (let i = 1; i < inputArray.length; i++) {
    const num = inputArray[i];
    if (num < nextNum) {
      moves += nextNum - num;
    }
    nextNum = Math.max(nextNum, num) + 1;
    console.log(num, nextNum);
  }
  return moves;
}

console.log(arrayChange([-1, 1, 0, 1]));
