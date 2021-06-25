function adjacentElementsProduct(inputArray: number[]): number {
  let max: number = null;
  for (let i = 1; i < inputArray.length; i++) {
    if (max === null || inputArray[i] * inputArray[i - 1] > max) {
      max = inputArray[i] * inputArray[i - 1];
    }
  }
  return max;
}
