function shapeArea(n: number): number {
  if (n <= 0) return 0;
  return n ** 2 + (n - 1) ** 2;
}
