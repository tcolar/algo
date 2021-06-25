function commonCharacterCount(s1: string, s2: string): number {
  const m1 = {};
  for (const s of s1) {
    m1[s] = (m1[s] || 0) + 1;
  }
  let count = 0;
  for (const s of s2) {
    if (s in m1 && m1[s] > 0) {
      count++;
      m1[s]--;
    }
  }
  return count;
}
