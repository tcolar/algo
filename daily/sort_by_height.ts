function sortByHeight(a: number[]): number[] {
  let ppl: number[] = [];
  for (let i = 0; i < a.length; i++) {
    if (a[i] !== -1) ppl.push(a[i]);
  }
  ppl.sort((a, b) => (a < b ? -1 : 1));
  // console.log(ppl)
  for (let i = 0; i < a.length; i++) {
    if (a[i] !== -1) a[i] = ppl.shift();
  }
  return a;
}
