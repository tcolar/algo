function firstNotRepeatingCharacter(s: string): string {
  const m = new Map<string, number>();
  for (const c of s) {
    m.set(c, (m.get(c) || 0) + 1);
  }
  for (const [k, v] of m) {
    if (v === 1) {
      return k;
    }
  }
  return '_';
}
