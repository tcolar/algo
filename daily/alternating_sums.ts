function alternatingSums(a: number[]): number[] {
  const sums: number[] = [0, 0];
  let cpt = 0;
  for (const nb of a) {
    const team = cpt % 2;
    sums[team] += nb;
    cpt++;
  }
  return sums;
}
