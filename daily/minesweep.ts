function minesweeper(matrix: boolean[][]): number[][] {
  const ms: number[][] = [];
  matrix.forEach((row, i) => {
    ms.push(
      row.map((b, j) => {
        return count(matrix, i, j);
      })
    );
  });
  return ms;
}

function count(matrix: boolean[][], i: number, j: number): number {
  let total = 0;
  for (let k = -1; k <= 1; k++) {
    for (let l = -1; l <= 1; l++) {
      if (i + k < 0 || i + k >= matrix.length || j + l < 0 || j + l >= matrix[i].length) {
        continue;
      }
      if (k === 0 && l === 0) continue;
      if (matrix[i + k][j + l]) total++;
    }
  }
  return total;
}
