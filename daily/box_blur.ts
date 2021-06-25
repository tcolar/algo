function boxBlur(image: number[][]): number[][] {
  const result: number[][] = [];
  for (let i = 1; i < image.length - 1; i++) {
    const newRow: number[] = [];
    for (let j = 1; j < image[i].length - 1; j++) {
      let x = 0;
      for (let k = -1; k <= 1; k++) {
        for (let l = -1; l <= 1; l++) {
          x += image[i + k][j + l];
        }
      }
      newRow.push(Math.floor(x / 9));
    }
    result.push(newRow);
  }
  return result;
}
