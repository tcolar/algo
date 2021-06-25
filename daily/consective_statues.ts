function makeArrayConsecutive2(statues: number[]): number {
  const max = Math.max(...statues);
  const min = Math.min(...statues);
  return max - min - statues.length + 1;
}
