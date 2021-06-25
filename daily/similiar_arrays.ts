function areSimilar(a: number[], b: number[]): boolean {
  if (a.length !== b.length) return false;
  let didSwap = false;
  let toSwap: number[] = [];
  for (let i = 0; i < a.length; i++) {
    if (a[i] === b[i]) continue;
    if (didSwap) return false; // one swap max
    if (toSwap.length) {
      // we have a potential swap pending, check if it matches
      if (a[i] !== toSwap[1] || b[i] !== toSwap[0]) return false;
      // it matched, but we are now out of swaps
      didSwap = true;
    } else {
      // first mismatch, store it so we can check for swap later
      toSwap = [a[i], b[i]];
    }
  }
  return true;
}
