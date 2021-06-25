function isLucky(n: number): boolean {
  const a = [];
  while (n > 0) {
    a.push(n % 10);
    n = Math.floor(n / 10);
  }
  console.log(a);
  if (a.length % 2 === 1) return false;
  let sum = 0;
  for (let i = 0; i < a.length / 2; i++) {
    sum += a[i] - a[a.length - 1 - i];
  }
  return sum === 0;
}
