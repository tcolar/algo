function matrixElementsSum(matrix: number[][]): number {
  let skipCols = [];
  var sum = 0;
  for (var i = 0; i < matrix.length; i++) {
    for (var j = 0; j < matrix[i].length; j++) {
      const cell = matrix[i][j];
      if (cell === 0) {
        skipCols.push(j);
      }
      if (skipCols.includes(j)) continue;
      sum += cell;
    }
  }
  return sum;
}
