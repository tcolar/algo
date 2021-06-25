/*
In chess, queens can move any number of squares vertically, horizontally, or diagonally. The n-queens puzzle is the problem of placing 
n queens on an n × n chessboard so that no two queens can attack each other.

Given an integer n, print all possible distinct solutions to the n-queens puzzle. Each solution contains distinct board configurations of the placement of the n queens, 
where the solutions are arrays that contain permutations of [1, 2, 3, .. n]. The number in the ith position of the results array indicates that the ith column queen is 
placed in the row with that number. In your solution, the board configurations should be returned in lexicographical order.
*/

function nQueens(n: number): number[][] {
  return solutionsForCol(n, 1);
}

function solutionsForCol(n: number, col: number, queens: number[] = []): number[][] {
  const solutions: number[][] = [];
  for (let row = 1; row <= n; row++) {
    if (isOpen(col, row, queens)) {
      const next = [...queens, row];
      //console.log(col, row, next);
      if (next.length === n) {
        solutions.push(next);
      }
      solutions.push(...solutionsForCol(n, col + 1, next));
    }
  }
  return solutions;
}

function isOpen(newCol: number, newRow: number, queens: number[]): boolean {
  let col = 1;
  for (const row of queens) {
    if (newCol === col) {
      // same row
      return false;
    }
    if (newRow === row) {
      // same col
      return false;
    }
    if (Math.abs(newCol - col) === Math.abs(newRow - row)) {
      // same diagonal
      return false;
    }
    col++;
  }
  return true;
}

console.log(solutionsForCol(1, 0));
console.log(solutionsForCol(4, 0));
