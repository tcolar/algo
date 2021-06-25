function almostIncreasingSequence(sequence: number[]): boolean {
  let oooIndex = null; // outoforder item index
  for (let i = 0; i < sequence.length; i++) {
    if (sequence[i] <= sequence[i - 1]) {
      if (oooIndex !== null) {
        console.log('more than one bad index, not fixable', oooIndex, i);
        return false;
      }
      oooIndex = i;
    }
  }
  console.log(oooIndex);
  if (!oooIndex) return true; // everything was in order
  if (oooIndex === 1) {
    console.log('we can drop the head');
    return true;
  }
  if (oooIndex === sequence.length - 1) {
    console.log('we can drop the tail');
    return true;
  }
  // would removing the item at the bad index fix it ?
  if (sequence[oooIndex - 1] < sequence[oooIndex + 1]) {
    console.log('rm cur', sequence[oooIndex]);
    return true;
  }
  // would removing the item at the previous index fix it ?
  if (oooIndex > 1 && sequence[oooIndex - 2] < sequence[oooIndex]) {
    console.log('rm prev', sequence[oooIndex - 1]);
    return true;
  }
  return false;
}
